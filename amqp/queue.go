package amqp

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

// Package queue implements all the queue handling with megam. It abstracts
// which queue server is being used, how the message gets marshaled in to the
// wire and how it's read.
//
// It provides a basic type: Message. You can Put, Get, Delete and Release
// messages, using methods and functions with respective names.
//
// It also provides a generic, thread safe, handler for messages, with start
// and stop capability.
import (
	"fmt"
    "github.com/streadway/amqp"
	"github.com/tsuru/config"
)

// PubSubQ represents an implementation that allows Publishing and
// Subscribing messages.
type PubSubQ interface {
	// Publishes a message using the underlaying queue server.
	Pub(msg []byte) error

	// Returns a channel that will yield every message published to this
	// queue.
	Sub() (chan []byte, error)

	// Unsubscribe the queue, this should make sure the channel returned
	// by Sub() is closed.
	UnSub() error
}

// QFactory manages queues. It's able to create new queue and handler
// instances.
type QFactory interface {
	// Get returns a queue instance, identified by the given name.
	Get(name string) (PubSubQ, error)
	
	Dial() (*amqp.Connection, error)
}

var factories = map[string]QFactory{
	"rabbitmq": &rabbitmqQFactory{},
}

// Register registers a new queue factory. This is how one would add a new
// queue to tsuru.
func Register(name string, factory QFactory) {
	factories[name] = factory
}

// Factory returns an instance of the QFactory used in tsuru. It reads tsuru
// configuration to find the currently used queue system and returns an
// instance of the configured system, if it's registered. Otherwise it
// will return an error.
func Factory() (QFactory, error) {
	name, err := config.GetString("queue")
	if err != nil {
		name = "rabbitmq"
	}
	if f, ok := factories[name]; ok {
		return f, nil
	}
	return nil, fmt.Errorf("Queue %q is not known.", name)
}


// Message represents the message stored in the queue.
//
// A message is specified by an action and a slice of strings, representing
// arguments to the action.
//
// For example, the action "NSTART" could receive one argument: the
// name of the app for which the app will be stopped.
type Message struct {
	Action string //action NSTART, NSTOP, NRESTART etc.
	Args   string //any arguments as deemed fit.
	Id     string //the id in Riak which starts like RIP..
	mid    uint64 //a counter incremented each time the msg is received.
	delete bool
}