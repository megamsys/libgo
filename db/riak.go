// Package db encapsulates connection with Riak.
//
// The function Open dials to Riak and returns a connection (represented by
// the Storage type). It manages an internal pool of connections, and
// reconnects in case of failures. That means that you should not store
// references to the connection, but always call Open.
package db

import (
	"fmt"
	"github.com/mrb/riakpbc"
	"github.com/tsuru/config"
	"log"
	"strings"
	"sync"
	"time"
)

var (
	conn   = make(map[string]*session) // pool of connections
	mut    sync.RWMutex                // for pool thread safety
	ticker *time.Ticker                // for garbage collection
)

const (
	DefaultRiakURL    = "127.0.0.1:8087"
	DefaultBucketName = "appreqs"
)

const period time.Duration = 7 * 24 * time.Hour

type session struct {
	s    *riakpbc.Client
	used time.Time
}

// Storage holds the connection with the bucket name.
type Storage struct {
	coder_client *riakpbc.Client
	bktname      string
}

func open(addr []string, bucketname string) (*Storage, error) {
	log.Printf("--> Dialing to %v", addr)
	coder := riakpbc.NewCoder("json", riakpbc.JsonMarshaller, riakpbc.JsonUnmarshaller)
	riakCoder := riakpbc.NewClientWithCoder(addr, coder)
	if err := riakCoder.Dial(); err != nil {
		return nil, err
	}

	// Set Client ID
	/*if _, err := riakCoder.SetClientId("coolio"); err != nil {
		log.Fatalf("Setting client ID failed: %v", err)
	}
	*/
	storage := &Storage{coder_client: riakCoder, bktname: bucketname}

	mut.Lock()
	conn[strings.Join(addr, "::")] = &session{s: riakCoder, used: time.Now()}
	mut.Unlock()
	return storage, nil
}

// Open dials to the Riak database, and return a new connection (represented
// by the type Storage).
//
// addr is a Riak connection URI, and bktname is the name of the bucket.
//
// This function returns a pointer to a Storage, or a non-nil error in case of
// any failure.
func Open(addr []string, bktname string) (storage *Storage, err error) {
	log.Printf("--> Connecting to %v", addr)
	defer func() {
		if r := recover(); r != nil {
			log.Printf("--> Recovered from panic")
			storage, err = open(addr, bktname)
		}
	}()
	mut.RLock()
	if session, ok := conn[strings.Join(addr, "::")]; ok {
		mut.RUnlock()
		if _, err = session.s.Ping(); err == nil {
			mut.Lock()
			session.used = time.Now()
			conn[strings.Join(addr, "::")] = session
			mut.Unlock()
		}
		return open(addr, bktname)
	}
	mut.RUnlock()
	return open(addr, bktname)
}

// Conn reads the megam config and calls Open to get a database connection.
//
// Most megam packages should probably use this function. Open is intended for
// use when supporting more than one database.
func Conn(bktname string) (*Storage, error) {
	url, _ := config.GetString("riak:url")
	if url == "" {
		url = DefaultRiakURL
	}
	//bktname, _ := config.GetString("riak:bucket")
	if bktname == "" {
		bktname = DefaultBucketName
	}
	tadr := []string{url}
	log.Printf("%v %s", tadr, bktname)
	return Open(tadr, bktname)
}

// Close closes the storage, releasing the connection.
func (s *Storage) Close() {
	log.Printf("---] Closing storage %v", s)
	s.coder_client.Close()
}

// FetchStruct stores a struct  as JSON
//   eg: data := ExampleData{
//        Field1: "ExampleData1",
//        Field2: 1,
//   }
// So the send can pass in 	out := &ExampleData{}
// Apps returns the apps collection from MongoDB.
func (s *Storage) FetchStruct(key string, out interface{}) error {
	if _, err := s.coder_client.FetchStruct(s.bktname, key, out); err != nil {
		return fmt.Errorf("Convert fetched JSON to the Struct, and return it failed: %s", err)
	}
	fmt.Println(out)
	//TO-DO:
	//we need to return the fetched json -> to struct interface
	return nil
}

// StoreStruct returns the apps collection from MongoDB.
func (s *Storage) StoreStruct(key string, data interface{}) error {
	if _, err := s.coder_client.StoreStruct(s.bktname, key, data); err != nil {
		return fmt.Errorf("Convert fetched JSON to the Struct, and return it failed: %s", err)
	}
	return nil
}

type SshObject struct{
	  Data string
	}

// Fetch raw data (int, string, []byte)
 func (s *Storage) FetchObject(key string, out *SshObject) error {

	obj, err := s.coder_client.FetchObject(s.bktname, key)
	if err != nil {
		return fmt.Errorf("Convert fetched JSON to the Struct, and return it failed: %s", err)
	}
    out.Data = string(obj.GetContent()[0].GetValue())
	return nil
}

// Store raw data (int, string, []byte)
func (s *Storage) StoreObject(key string, data string) error {
	if _, err := s.coder_client.StoreObject(s.bktname, key, []byte(data)); err != nil {
		return fmt.Errorf("Convert fetched JSON to the Struct, and return it failed: %s", err)
	}
	return nil
}

func init() {
	ticker = time.NewTicker(time.Hour)
	go retire(ticker)
}

// retire retires old connections
func retire(t *time.Ticker) {
	for _ = range t.C {
		now := time.Now()
		var old []string
		mut.RLock()
		for k, v := range conn {

			if now.Sub(v.used) >= period {
				old = append(old, k)
			}
		}
		mut.RUnlock()
		mut.Lock()
		for _, c := range old {
			log.Printf("--> Stale conn[%s]", c)
			conn[c].s.Close()
			delete(conn, c)
		}
		mut.Unlock()
	}
}
