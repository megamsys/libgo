package amqp

import ()

import (
	"fmt"
	"time"

	"gopkg.in/check.v1"
)

type RabbitMQSuite struct{ psq PubSubQ }

var _ = check.Suite(&RabbitMQSuite{})

func (s *RabbitMQSuite) SetUpSuite(c *check.C) {
	// init BeeLogger
	p, err := NewRabbitMQ("amqp://guest:guest@localhost:5672/", "testq")
	s.psq = p
	c.Assert(err, check.IsNil)
}

func (s *RabbitMQSuite) TestRabbitMQDial(c *check.C) {
	rf, err := Factory()
	rf.(*rabbitmqQFactory).BindAddress = "amqp://guest:guest@localhost:5672/"
	_, err = rf.Dial()
	c.Assert(err, check.IsNil)
}

func (s *RabbitMQSuite) TestRabbitMQPubSub(c *check.C) {
	pubSubQ, ok := s.psq.(PubSubQ)
	c.Assert(ok, check.Equals, true)
	msgChan, err := pubSubQ.Sub()
	c.Assert(err, check.IsNil)
	err = pubSubQ.Pub([]byte("howdy pubsub"))
	c.Assert(err, check.IsNil)
	c.Assert(<-msgChan, check.DeepEquals, []byte("howdy pubsub"))
}

func (s *RabbitMQSuite) TestRabbitMQPubSubUnsub(c *check.C) {
	pubSubQ, ok := s.psq.(PubSubQ)
	c.Assert(ok, check.Equals, true)
	err := pubSubQ.Pub([]byte("howdy pubsubunsub"))
	c.Assert(err, check.IsNil)

	msgChan, err := pubSubQ.Sub()
	c.Assert(err, check.IsNil)

	done := make(chan bool)

	err = pubSubQ.Pub([]byte("howdy pubsubunsub"))
	c.Assert(err, check.IsNil)

	go func() {
		time.Sleep(1e9)
		pubSubQ.UnSub()
	}()
	go func() {
		msgs := make([][]byte, 0)

		for msg := range msgChan {
			fmt.Println(" [x] %q", msg)
			msgs = append(msgs, msg)
		}
		c.Assert(msgs, check.DeepEquals, [][]byte{[]byte("howdy pubsubunsub")})
		done <- true
	}()
	select {
	case <-done:
	case <-time.After(5e9):
		c.Error("Timeout waiting for message.")
	}
}
