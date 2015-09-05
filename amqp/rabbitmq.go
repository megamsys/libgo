/*
** Copyright [2012-2014] [Megam Systems]
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
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/streadway/amqp"
	"sync"
)

type rabbitmqQ struct {
	name    string
	prefix  string
	factory *rabbitmqQFactory
	psc     *amqp.Channel
}

func (r *rabbitmqQ) exchname() string {
	return r.prefix + "_" + r.name + "_exchange"
}

func (r *rabbitmqQ) qname() string {
	return r.prefix + "_" + r.name + "_queue"
}

func (r *rabbitmqQ) tag() string {
	return r.prefix + "_" + r.name + "_tag"
}

func (r *rabbitmqQ) key() string {
	return r.prefix + "_" + r.name + "_key"
}

func (r *rabbitmqQ) Pub(msg []byte) error {
	chnl, err := r.factory.dial(r.exchname()) // return amqp.Channel
	if err != nil {
		return err
	}

	log.Info(fmt.Printf("pubh msg (%s, %s) (%q).", r.exchname(), r.key(), msg))

	if err = chnl.Publish(
		r.exchname(), // publish to an exchange
		r.key(),      // routing to 0 or more queues
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            msg,
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		return fmt.Errorf("Failed to publish message in exchange: %s", err)
	}
	log.Info("pubh msg SUCCESS.")
	return err
}

func (r *rabbitmqQ) UnSub() error {
	if r.psc == nil {
		return nil
	}
	err := r.psc.Cancel(r.tag(), false)
	if err != nil {
		return err
	}
	return nil
}

func (r *rabbitmqQ) Sub() (chan []byte, error) {
	chnl, err := r.factory.getChonn(r.key(), r.exchname(), r.qname())
	if err != nil {
		return nil, err
	}

	r.psc = chnl

	msgChan := make(chan []byte)

	log.Infof("subs (%s,%s)", r.qname(), r.tag())

	deliveries, err := chnl.Consume(
		r.qname(), // name
		r.tag(),   // consumerTag,
		true,      // noAck
		false,     // exclusive
		false,     // noLocal
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return nil, err
	}
	log.Infof("subs %s (%s) SUCCESS.", r.qname(), r.tag())

	//This is asynchronous, the callee will have to wait.
	go func() {
		//defer close(msgChan)
		for d := range deliveries {
			log.Infof("subd %s: (%v,%q)", r.qname(), d.DeliveryTag, d.Body)
			msgChan <- d.Body
		}

	}()
	return msgChan, nil
}

type rabbitmqQFactory struct {
	BindAddress string
	sync.Mutex
}

func (f *rabbitmqQFactory) Get(name string) (PubSubQ, error) {
	return &rabbitmqQ{name: name, prefix: "megam", factory: f}, nil
}

func (f *rabbitmqQFactory) Dial() (*amqp.Connection, error) {
	log.Infof("dial  amqp (%s)\n", f.BindAddress)
	conn, err := amqp.Dial(f.BindAddress)

	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (f *rabbitmqQFactory) dial(exchname string) (*amqp.Channel, error) {
	conn, err := f.Dial()

	if err != nil {
		return nil, err
	}

	log.Infof("dial amqp (%s)\n", f.BindAddress)

	chnl, err := conn.Channel()

	if err != nil {
		return nil, err
	}

	//defer chnl.Close()

	if err = chnl.ExchangeDeclare(
		exchname, // name of the exchange
		"fanout", // exchange Type
		true,     // durable
		false,    // delete when complete
		false,    // internal
		false,    // noWait
		nil,      // arguments
	); err != nil {
		return nil, err
	}

	log.Infof("conn amqp (%s) SUCCESS.\n", f.BindAddress, exchname)
	return chnl, err
}

func (factory *rabbitmqQFactory) getChonn(key string, exchname string, qname string) (*amqp.Channel, error) {
	chnl, err := factory.dial(exchname)
	if err != nil {
		return nil, err
	}
	log.Infof("dial amqp  (%s)\n", exchname)

	qu, err := chnl.QueueDeclare(
		qname, // name of the queue
		true,  // durable
		false, // delete when usused
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}

	log.Infof("decl queue (%s)\n", qname)

	if err = chnl.QueueBind(
		qu.Name, // name of the queue
		key,
		exchname,
		false, // noWait
		nil,   // arguments
	); err != nil {
		return nil, err
	}

	log.Infof("bound queue (%s,%s,%s)\n", qname, exchname, key)
	return chnl, nil
}
