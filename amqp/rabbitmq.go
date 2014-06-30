/*
** Copyright [2012-2013] [Megam Systems]
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
 */
package amqp

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/tsuru/config"
	"log"
	"regexp"
	"sync"
	"time"
)

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   *amqp.Queue
	tag     string
	done    chan error
}

type rabbitmqQ struct {
	name string
}

const (
	DefaultAMQPURL      = "amqp://localhost:5672/"
	DefaultQueue        = "megam_node_stale"
	DefaultExchange     = "megam_node_exchange_stale"
	DefaultExchangeType = "fanout"
	DefaultRoutingKey   = "megam_key"
	DefaultConsumerTag  = "megam_node_consumer_stale"
)

var (
	mut            sync.Mutex // for conn access
	timeoutRegexp  = regexp.MustCompile(`(TIMED_OUT|timeout)$`)
	notFoundRegexp = regexp.MustCompile(`not found$`)
)

func (b *rabbitmqQ) Get(timeout time.Duration) (*Message, error) {
	return nil, errors.New("Get: Not supported, Handler.start(), subscribe for RabbitMQ.")

}

func (b *rabbitmqQ) Put(m *Message, delay time.Duration) error {
	cons, err := connection()
	if err != nil {
		return err
	}

	//convert Message to "body" bytes
	var body = m.Args
	log.Printf("Publishing %dB message (%q).", len(body), body)

	exchange_conf, _ := config.GetString("amqp:exchange")
	if exchange_conf == "" {
		exchange_conf = DefaultExchange
	}
	routingkey_conf, _ := config.GetString("amqp:routingkey")
	if routingkey_conf == "" {
		routingkey_conf = DefaultRoutingKey
	}

	if err = cons.channel.Publish(
		exchange_conf,   // publish to an exchange
		routingkey_conf, // routing to 0 or more queues
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(body),
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		return fmt.Errorf("Exchange Publish: %s", err)
	}
	return err
}

func (b *rabbitmqQ) Delete(m *Message, tag uint64, multiple bool) error {
	log.Printf("%-6s:%s [%v][%d]", "WARN", "Acknowledge not Implemented yet.", m, tag)
	return nil

}

func (b *rabbitmqQ) Release(m *Message, tag uint64, multiple bool, requeue bool) error {
	log.Printf("%-6s:%s [%v][%d]", "WARN", "NAcknowledge not Implemented yet.", m, tag)
	return nil
}

type rabbitmqFactory struct{}

func (b rabbitmqFactory) Get(name string) (Q, error) {
	return &rabbitmqQ{name: name}, nil
}

func (b rabbitmqFactory) Handler(f func(*Message), name ...string) (Handler, error) {
	log.Printf("Attaching handler for RabbitMQ.")
	return &executor{
		inner: func() {
			log.Printf("Waiting for deliveries from consumers.")

			if deliveries, err := consume(5e9); err == nil {
				for d := range deliveries {
					log.Printf("%dB : [%v] %q", len(d.Body), d.DeliveryTag, d.Body)
					var message Message
					err := json.Unmarshal(d.Body, &message)
					if err != nil {
						fmt.Println("error:", err)
					}
					log.Printf("%+v", message)

					go func(m *Message, tag uint64) {
						f(m)
						q := rabbitmqQ{}
						if m.delete {
							q.Delete(m, tag, false)
						} else {
							q.Release(m, tag, false, false) //don't requeue it.
						}
					}(&message, d.DeliveryTag)
				}
				log.Printf("TO-DO: Deliveries channel closed")
				//done <- nil
			} else {
				log.Println(fmt.Errorf("Dial: %s", err))
				time.Sleep(5e9)
			}
		},
	}, nil
}

func connection() (*Consumer, error) {
	var (
		addr string
		err  error
	)

	mut.Lock()
	c := &Consumer{
		conn:    nil,
		channel: nil,
		tag:     DefaultConsumerTag,
		done:    make(chan error),
	}

	if c.conn == nil {
		mut.Unlock()
		addr, err = config.GetString("amqp:url")
		if err != nil {
			addr = DefaultAMQPURL
		}
		mut.Lock()
		if c.conn, err = amqp.Dial(addr); err != nil {
			mut.Unlock()
			return nil, err
		}
	}
	log.Printf("Connected to (%s)", addr)

	if c.channel, err = c.conn.Channel(); err != nil {
		mut.Unlock()
		return nil, err
	}

	exchange_conf, _ := config.GetString("amqp:exchange")
	if exchange_conf == "" {
		exchange_conf = DefaultExchange
	}
	log.Printf("Connected to (%s)", exchange_conf)

	if err = c.channel.ExchangeDeclare(
		exchange_conf,       // name of the exchange
		DefaultExchangeType, // exchange Type
		true,                // durable
		false,               // delete when complete
		false,               // internal
		false,               // noWait
		nil,                 // arguments
	); err != nil {
		mut.Unlock()
		return nil, err
	}

	log.Printf("Connection successful to  (%s,%s)", addr, exchange_conf)
	mut.Unlock()
	return c, err
}

func rconnection() (*Consumer, error) {
	cons, err := connection()
	if err != nil {
		return nil, err
	}

	mut.Lock()
	queue_conf, _ := config.GetString("amqp:queue")
	if queue_conf == "" {
		queue_conf = DefaultQueue
	}

	decl_q, err := cons.channel.QueueDeclare(
		queue_conf, // name of the queue
		true,       // durable
		false,      // delete when usused
		false,      // exclusive
		false,      // noWait
		nil,        // arguments
	)

	if err != nil {
		mut.Unlock()
		return nil, err
	}
	log.Printf("Connected to (%s)", queue_conf)

	cons.queue = &decl_q

	exchange_conf, _ := config.GetString("amqp:exchange")
	if exchange_conf == "" {
		exchange_conf = DefaultExchange
	}
	routingkey_conf, _ := config.GetString("amqp:routingkey")
	if routingkey_conf == "" {
		routingkey_conf = DefaultRoutingKey
	}

	if err = cons.channel.QueueBind(
		cons.queue.Name, // name of the queue
		routingkey_conf,
		exchange_conf,
		false, // noWait
		nil,   // arguments
	); err != nil {
		mut.Unlock()
		return nil, err
	}
	mut.Unlock()

	log.Printf("Connection successful to (%s,%s,%s)", queue_conf, exchange_conf, routingkey_conf)
	return cons, nil
}

//returns AMQP Consumer (ASynchronous, blocked - dies on shutdown)
func consume(timeout time.Duration) (<-chan amqp.Delivery, error) {

	cons, err := rconnection()
	if err != nil {
		return nil, err
	}

	log.Printf("Starting consumer (%s,%s)", cons.queue.Name, cons.tag)

	deliveries, err := cons.channel.Consume(
		cons.queue.Name, // name
		cons.tag,        // consumerTag,
		false,           // noAck
		false,           // exclusive
		false,           // noLocal
		false,           // noWait
		nil,             // arguments
	)

	if err != nil {
		return nil, err
	}
	log.Printf("Started consumer (%s,%s)", cons.queue.Name, cons.tag)

	return deliveries, nil
}

/*
//shut it down, the handler actually shuts it down.
func (c *Consumer) Shutdown() error {
	// will close() the deliveries channel
	if err := c.channel.Cancel(c.tag, true); err != nil {
		return fmt.Errorf("Consumer cancel failed: %s", err)
	}

	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}

	defer log.Printf("AMQP shutdown OK")

	// wait for handle() to exit
	return <-c.done
}

*/
