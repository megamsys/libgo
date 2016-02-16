package db
/*
import (
//	"gopkg.in/check.v1"
//	"strings"
//	"sync"
//	"time"
)

type ExampleData struct {
	Field1 string `riak:"index" json:"field1"`
	Field2 int    `json:"field2"`
}

type SampleObject struct {
	Data string `json:"data"`
}

var addr = []string{"127.0.0.1:8087"}

const bkt = "sample1"

func (s *S) TestOpenReconnects(c *check.C) {
	storage, err := Open(addr, bkt)
	c.Assert(err, check.IsNil)
	storage.Close()
	storage, err = Open(addr, bkt)
	defer storage.Close()
	c.Assert(err, check.IsNil)
	_, err = storage.coder_client.Ping()
	c.Assert(err, check.IsNil)
}

func (s *S) TestOpenConnectionRefused(c *check.C) {
	storage, err := Open([]string{"127.0.0.1:68098"}, bkt)
	c.Assert(storage, check.IsNil)
	c.Assert(err, check.NotNil)
}

func (s *S) TestClose(c *check.C) {

	defer func() {
		r := recover()
		c.Check(r, check.IsNil)
	}()

	storage, err := Open(addr, bkt)
	defer storage.Close()
	c.Assert(err, check.IsNil)
	c.Assert(storage, check.NotNil)
	_, err = storage.coder_client.Ping()
	c.Check(err, check.IsNil)
}

func (s *S) TestConn(c *check.C) {
	r, err := NewRiakDB(addr, bkt)
	storage, err := r.Conn()
	defer storage.Close()
	c.Assert(storage, check.NotNil)
	c.Assert(err, check.IsNil)
	_, err = storage.coder_client.Ping()
	c.Check(err, check.IsNil)
}

func (s *S) TestStore(c *check.C) {
	r, err := NewRiakDB(addr, bkt)
	storage, err := r.Conn()
	defer storage.Close()
	c.Assert(storage, check.NotNil)
	c.Assert(err, check.IsNil)
	data := ExampleData{
		Field1: "ExampleData1",
		Field2: 1,
	}
	err = storage.StoreStruct("sampledata", &data)
	c.Assert(err, check.IsNil)
}

func (s *S) TestFetch(c *check.C) {
	r, err := NewRiakDB(addr, bkt)
	storage, err := r.Conn()
	defer storage.Close()
	c.Assert(storage, check.NotNil)
	c.Assert(err, check.IsNil)
	out := &ExampleData{}
	err = storage.FetchStruct("sampledata", out)
	c.Assert(err, check.IsNil)
}

func (s *S) TestStoreObject(c *check.C) {
	r, err := NewRiakDB(addr, bkt)
	storage, err := r.Conn()
	defer storage.Close()
	c.Assert(storage, check.NotNil)
	c.Assert(err, check.IsNil)
	data := "sampledata"
	err = storage.StoreObject("sampleobject", data)
	c.Assert(err, check.IsNil)
}

func (s *S) TestFetchObject(c *check.C) {
	r, err := NewRiakDB(addr, bkt)
	storage, err := r.Conn()
	defer storage.Close()
	c.Assert(storage, check.NotNil)
	c.Assert(err, check.IsNil)
	out := &SomeObject{}
	err = storage.FetchObject("sampleobject", out)
	c.Assert(err, check.IsNil)
}

func (s *S) TestDeleteObject(c *check.C) {
	r, err := NewRiakDB(addr, bkt)
	storage, err := r.Conn()
	defer storage.Close()
	c.Assert(storage, check.NotNil)
	c.Assert(err, check.IsNil)
	data := "sampledata"
	err = storage.StoreObject("sampleobject1", data)
	c.Assert(err, check.IsNil)
	err = storage.DeleteObject("sampleobject1")
	c.Assert(err, check.IsNil)
	out := &SomeObject{}
	err = storage.FetchObject("sampleobject1", out)
	c.Assert(err, check.NotNil)
}

func (s *S) TestRetire(c *check.C) {
	defer func() {
		if r := recover(); !c.Failed() && r == nil {
			c.Errorf("Should panic in ping, but did not!")
		}
	}()
	Open(addr, bkt)
	ky := strings.Join(addr, "::")
	sess := conn[ky]
	sess.used = sess.used.Add(-1 * 2 * period)
	conn[ky] = sess
	var ticker time.Ticker
	ch := make(chan time.Time, 1)
	ticker.C = ch
	ch <- time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		retire(&ticker)
		wg.Done()
	}()
	close(ch)
	wg.Wait()
	_, ok := conn[ky]
	c.Check(ok, check.Equals, false)
	sess1 := conn[ky]
	sess1.s.Ping()
}
*/
