package api

import (
	// "encoding/json"
	// "fmt"
	// "io"
	// "os"
	"testing"
	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type Api interface {
	GetKeys() map[string]string
}

type S struct {
  Email string  `json:"email"`
	Api_key string `json:"api_key"`
	States  States `json:"states"`
}

type States struct {
	Authority string  `json:"authority"`
}

func (s *S) Set(email, key string) {
	s.Email = email
	s.Api_key = key
	s.States = States{Authority: "user"}
}

func (s *S) Get() Api {
	return s
}

func (s *S) GetKeys() map[string]string {
  m := make(map[string]string)
	m["email"] = s.Email
	m["api_key"] = s.Api_key
	return m
}
var _ = check.Suite(&S{})

//we need make sure the stub deploy methods are supported.
func (s *S) SetUpSuite(c *check.C) {
  //server, err := testing.NewServer("127.0.0.1:5555")
	s.Set("info@megam.io","fakeapikey")
	// s.Email = "info@megam.io"
	// s.Api_key = "fakeapikey"
  //c.Assert(err, check.IsNil)
}


// func (s *S) TearDownSuite(c *check.C) {
//   //just stop the server here.
// }
