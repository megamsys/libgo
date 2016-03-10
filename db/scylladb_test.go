package db

/*import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/megamsys/gocql"
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

type Customer2 struct {
	Id   string
	Name string
	Age  int
}

var _ = check.Suite(&S{})

var noips = []string{"103.56.92.24"}

func (s *S) SetUpSuite(c *check.C) {
	s.sy, _ = NewScyllaDB(ScyllaDBOpts{
		KeySpaceName: "testing",
		NodeIps:      noips,
		Username:     "",
		Password:     "",
		Debug:        true,
	})
	c.Assert(s.sy, check.NotNil)

	if s.sy == nil {
		fmt.Println("------------- scylladb is not running")
		c.Skip("- ScyllaDB isn't running. Did you start it ? ")
	}

	cluster := gocql.NewCluster("103.56.92.24")
	cluster.Keyspace = "testing"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	var id gocql.UUID
	iter := session.Query(`SELECT * FROM CUSTOMER`).Iter()
	for iter.Scan(&id) {
		fmt.Println("Tweet:", id)
	}
}

func (s *S) TestReadWhereRowNotFound(c *check.C) {
	rand.Seed(time.Now().Unix())
	t := s.sy.Table("customer2", []string{"Id","Name"}, []string{}, &Customer2{})
	err := t.T.(gocassa.TableChanger).CreateIfNotExist()
	c.Assert(err, check.IsNil)
	err = t.Upsert(&Customer2{
		Id:   "1001",
		Name: "Hari",
		Age: 26,
	})
	c.Assert(err, check.IsNil)
	res := &Customer2{}
	err = t.ReadWhere(ScyllaWhere{clauses: map[string]string{"Id": "1001", "Name": ""}}, res)
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
	err = t.ReadWhere(ScyllaWhere{Clauses: map[string]string{"Id": "1001"}}, &res)
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

	res := &Customer{}
	err = t.ReadWhere(ScyllaWhere{clauses: map[string]string{"Id": "1001", "Name": "Joe"}}, res)
	c.Assert(err, check.NotNil)
}
	res := Customer{}
	err = t.ReadWhere(ScyllaWhere{Clauses: map[string]string{"Id": "1001", "Name": "Joe"}}, &res)
	c.Assert(err, check.IsNil)
}*/
