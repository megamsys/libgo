package api

import (
	"testing"
	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type Assembly struct {
	AccountId string `json:"accounts_id"  cql:"accounts_id"`
	OrgId   string `json:"org_id" cql:"org_id"`
	Id      string `json:"id"  cql:"id"`
}

type S struct {
  Email string  `json:"email"`
	Api_Key string `json:"api_key"`
	Master_Key string `json:"master_key"`
	Assembly Assembly
}

var _ = check.Suite(&S{})

//we need make sure the stub deploy methods are supported.
func (s *S) SetUpSuite(c *check.C) {
	s.Email = "info@megam.io"
	s.Api_Key = "fakeapikey"
	s.Assembly = Assembly{AccountId: 	s.Email, OrgId: "ORG123", Id: "asdf"}
}


// func (s *S) TearDownSuite(c *check.C) {
//   //just stop the server here.
// }
