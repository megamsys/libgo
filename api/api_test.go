package api

import (
  "io/ioutil"
   "fmt"
  	"gopkg.in/check.v1"
)

func (s *S) TestGetUser(c *check.C) {
  response, err := s.testGet("/accounts/" + s.ApiArgs.Email)
  c.Assert(err, check.IsNil)
  htmlData, err := ioutil.ReadAll(response.Body) //<--- here!
  c.Assert(err, check.IsNil)
  fmt.Println("Success  :",string(htmlData)) //<-- here !
  fmt.Println(err)
  c.Assert(err, check.IsNil)
}

type Assembly struct {
	AccountId string `json:"accounts_id"  cql:"accounts_id"`
	OrgId   string `json:"org_id" cql:"org_id"`
	Id      string `json:"id"  cql:"id"`
  Status  string `json:"status" cql:"status"`
}

type Components struct {
	Id      string `json:"id"  cql:"id"`
  Status  string `json:"status" cql:"status"`
}

func (s *S) TestGetAssembly(c *check.C) {
  response, err := s.testGet("/assembly/ASM5285833184590940525")
  c.Assert(err, check.IsNil)
  htmlData, err := ioutil.ReadAll(response.Body) //<--- here!
  c.Assert(err, check.IsNil)
 	fmt.Println("Success  :",string(htmlData)) //<-- here !
  fmt.Println(err)
  c.Assert(err, check.IsNil)
}

func (s *S) TestAssemblyPost(c *check.C) {
  response, err := s.testGet("/assembly/ASM5285833184590940525")
  c.Assert(err, check.IsNil)
  response, err := s.testPost("/assembly/update", Assembly{AccountId: 	s.ApiArgs.Email, OrgId: s.ApiArgs.Org_Id , Id: "ASM5285833184590940525",Status:"testing"})
  htmlData, err := ioutil.ReadAll(response.Body) //<--- here!
  c.Assert(err, check.IsNil)
 	fmt.Println("Success  :",string(htmlData)) //<-- here !
  fmt.Println(err)
  c.Assert(err, check.IsNil)
}

func (s *S) TestComponentPost(c *check.C) {
  response, err := s.testGet("/components/CMP5285833184590940525")
  c.Assert(err, check.IsNil)
  response, err := s.testPost("/components/update", Components{Id: "CMP5285833184590940525",Status:"testing"})
  c.Assert(err, check.IsNil)
  htmlData, err := ioutil.ReadAll(response.Body) //<--- here!
  c.Assert(err, check.IsNil)
  fmt.Println("Success  :",string(htmlData)) //<-- here !
  fmt.Println(err)
//  c.Assert(nil, check.NotNil)
}

func (s *S) testGet(path string) (*http.Response, error) {
  s.ApiArgs.Path = path
  cl := NewClient(s.ApiArgs)
  c.Assert(cl, check.NotNil)
  return cl.Get()
}

func (s *S) testPost(path string, item interface{})  {
  s.ApiArgs.Path = path
  cl := NewClient(s.ApiArgs)
  c.Assert(cl, check.NotNil)
  return cl.Post(item)
}
