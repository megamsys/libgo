package amqp

import ()

import (
	"log"
	"time"

	"github.com/tsuru/config"
	"gopkg.in/check.v1"
)

type RabbitMQSuite struct{ factory *rabbitmqQFactory }

var _ = check.Suite(&RabbitMQSuite{})

func (s *RabbitMQSuite) SetUpSuite(c *check.C) {
	s.factory = &rabbitmqQFactory{}
	config.Set("queue", "rabbitmq")
	_, err := s.factory.dial("unittest_exchange")
	c.Assert(err, check.IsNil)
}

func (s *RabbitMQSuite) TestFactoryGet(c *check.C) {
	var factory rabbitmqQFactory
	q, err := factory.Get("ancient")
	c.Assert(err, check.IsNil)
	rq, ok := q.(*rabbitmqQ)
	c.Assert(ok, check.Equals, true)
	c.Assert(rq.name, check.Equals, "ancient")
}

func (s *RabbitMQSuite) TestRabbitMQFactoryIsInFactoriesMap(c *check.C) {
	f, ok := factories["rabbitmq"]
	c.Assert(ok, check.Equals, true)
	_, ok = f.(*rabbitmqQFactory)
	c.Assert(ok, check.Equals, true)
}

func (s *RabbitMQSuite) TestRabbitMQPubSub(c *check.C) {
	var factory rabbitmqQFactory
	q, err := factory.Get("mypubsub")
	c.Assert(err, check.IsNil)
	pubSubQ, ok := q.(PubSubQ)
	c.Assert(ok, check.Equals, true)
	msgChan, err := pubSubQ.Sub()
	c.Assert(err, check.IsNil)
	err = pubSubQ.Pub([]byte("howdy pubsub"))
	c.Assert(err, check.IsNil)
	c.Assert(<-msgChan, check.DeepEquals, []byte("howdy pubsub"))
}

func (s *RabbitMQSuite) TestRabbitMQPubSubUnsub(c *check.C) {
	var factory rabbitmqQFactory
	q, err := factory.Get("mypubsub")
	c.Assert(err, check.IsNil)
	pubSubQ, ok := q.(PubSubQ)
	c.Assert(ok, check.Equals, true)
	msgChan, err := pubSubQ.Sub()
	c.Assert(err, check.IsNil)

	err = pubSubQ.Pub([]byte("howdy pubsubunsub"))
	c.Assert(err, check.IsNil)
	
	done := make(chan bool)
	
	go func() {
		time.Sleep(1e9)
		pubSubQ.UnSub()
	}()
	go func() {
		msgs := make([][]byte, 0)
		for msg := range msgChan {
			log.Printf(" [x] %q", msg)
			msgs = append(msgs, msg)
		}
		c.Assert(msgs, check.DeepEquals, [][]byte{[]byte("howdy pubsubunsub")})
		done <- true
	}()
	select {
	case <-done:
	case <-time.After(10e9):
		c.Error("Timeout waiting for message.")
	}
}
