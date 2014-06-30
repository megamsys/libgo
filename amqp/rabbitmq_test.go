package amqp

import ()

import (
	"github.com/tsuru/config"
	"gopkg.in/check.v1"
)

type RabbitMQSuite struct{}

var _ = check.Suite(&RabbitMQSuite{})

func (s *RabbitMQSuite) SetUpSuite(c *check.C) {
	//	config.Set("queue-server", "127.0.0.1:11300")
}

func (s *RabbitMQSuite) TestConnection(c *check.C) {
	_, err := connection()
	c.Assert(err, check.IsNil)
}

func (s *RabbitMQSuite) TestConnectionQueueServerUndefined(c *check.C) {
	old, _ := config.Get("amqp:url")
	config.Unset("amqp:url")
	defer config.Set("amqp:url", old)
	conn, err := connection()
	c.Assert(err, check.IsNil)
	c.Assert(conn, check.NotNil)
}

func (s *RabbitMQSuite) TestConnectionResfused(c *check.C) {
	old, _ := config.Get("amqp:url")
	config.Set("amqp:url", "127.0.0.1:11301")
	defer config.Set("amqp:url", old)
	conn, err := connection()
	c.Assert(conn, check.IsNil)
	c.Assert(err, check.NotNil)
}

/*func (s *RabbitMQSuite) TestPut(c *check.C) {
	msg := Message{
		Action: "startapp",
		Args:   []string{"node1.megam.co"},
	}
	q := rabbitmqQ{name: "default"}
	err := q.Put(&msg, 0)
	c.Assert(err, check.IsNil)
	c.Assert(msg.id, check.Not(check.Equals), 0)
	defer conn.Delete(msg.id)
	id, body, err := conn.Reserve(1e6)
	c.Assert(err, check.IsNil)
	c.Assert(id, check.Equals, msg.id)
	var got Message
	buf := bytes.NewBuffer(body)
	err = gob.NewDecoder(buf).Decode(&got)
	c.Assert(err, check.IsNil)
	got.id = msg.id
	c.Assert(got, check.DeepEquals, msg)
}*/

func (s *RabbitMQSuite) TestRabbitMQFactoryIsInFactoriesMap(c *check.C) {
	f, ok := factories["rabbitmq"]
	c.Assert(ok, check.Equals, true)
	_, ok = f.(rabbitmqFactory)
	c.Assert(ok, check.Equals, true)
}
