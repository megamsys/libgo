package db

import (
	"math/rand"
	"testing"
	"time"

	"github.com/megamsys/gocassa"
	"gopkg.in/check.v1"
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

	s.sy, _ = NewScyllaDB(ScyllaDBOpts{
		KeySpaceName: "testing",
		NodeIps:      noips,
		Username:     "",
		Password:     "",
		Debug:        true,
	})

	if s.sy == nil {
		c.Skip("- ScyllaDB isn't running. Did you start it ? ")
	}
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
	err = t.ReadWhere(ScyllaWhere{clauses: map[string]string{"Id": "1001", "Name": ""}}, &res)
	c.Assert(err, check.NotNil)
}

func (s *S) TestTablWithMultiplePKButReadUsingOnePK(c *check.C) {
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
	err = t.ReadWhere(ScyllaWhere{clauses: map[string]string{"Id": "1001", "Name": "Joe"}}, &res)
	c.Assert(err, check.IsNil)
}
