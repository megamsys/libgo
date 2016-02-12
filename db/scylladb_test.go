package db

import (
	"math/rand"
	"testing"
	"time"

	"gopkg.in/check.v1"
	"github.com/hailocab/gocassa"
)

func Test(t *testing.T) { check.TestingT(t) }

type S struct {
	sy *ScyllaDB
}

type Customer struct {
	Id   string
	Name string
}

var _ = check.Suite(&S{})

var noips = []string{"127.0.0.1"}

func (s *S) SetUpSuite(c *check.C) {
	ticker.Stop()

	s.sy, _ = NewScyllaDB(ScyllaDBOpts{
		KeySpaceName: "testing",
		NodeIps:      noips,
		Username:     "",
		Password:     "",
		Debug:        true,
	})
	c.Assert(s.sy, check.NotNil)
}

func (s *S) TestReadWhereRowNotFound(c *check.C) {
	rand.Seed(time.Now().Unix())
	t := s.sy.Table("customer", []string{"Id", "Name"}, []string{}, &Customer{})
	err := t.T.(gocassa.TableChanger).CreateIfNotExist()
	c.Assert(err, check.IsNil)
	err = t.Upsert(&Customer{
		Id:   "1001",
		Name: "Joe",
	})
	c.Assert(err, check.IsNil)
	res := Customer{}
	err = t.ReadWhere(ScyllaWhere{clauses: map[string]string{"Id": "1001", "Name":""}}, &res)
	c.Assert(err, check.NotNil)
}

func (s *S) TestMultiplePKReadUsingPK1(c *check.C) {
	rand.Seed(time.Now().Unix())
	t := s.sy.Table("customer", []string{"Id", "Name"}, []string{}, &Customer{})
	err := t.T.(gocassa.TableChanger).CreateIfNotExist()
	c.Assert(err, check.IsNil)
	err = t.Upsert(&Customer{
		Id:   "1001",
		Name: "Joe",
	})
	c.Assert(err, check.IsNil)
	res := Customer{}
	err = t.ReadWhere(ScyllaWhere{clauses: map[string]string{"Id": "1001"}}, &res)
	c.Assert(err, check.NotNil)
}

func (s *S) TestReadWhereRowFound(c *check.C) {
	rand.Seed(time.Now().Unix())
	t := s.sy.Table("customer", []string{"Id", "Name"}, []string{}, &Customer{})
	err := t.T.(gocassa.TableChanger).CreateIfNotExist()
	c.Assert(err, check.IsNil)
	err = t.Upsert(&Customer{
		Id:   "1001",
		Name: "Joe",
	})
	c.Assert(err, check.IsNil)
	res := Customer{}
	err = t.ReadWhere(ScyllaWhere{clauses: map[string]string{"Id": "1001", "Name":"Joe"}}, &res)
	c.Assert(err, check.IsNil)
}
