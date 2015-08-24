package amqp

import (
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type S struct{}

var _ = check.Suite(&S{})

func (s *S) TestFactory(c *check.C) {
	f, err := Factory()
	c.Assert(err, check.IsNil)
	_, ok := f.(*rabbitmqQFactory)
	c.Assert(ok, check.Equals, true)
}

func (s *S) TestRegister(c *check.C) {
	Register("unregistered", &rabbitmqQFactory{})
	_, err := Factory()
	c.Assert(err, check.IsNil)
}
