package amqp

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
	"github.com/tsuru/config"
	"time"
)

// Q represents a queue. A queue is a type that supports the set of
// operations described by this interface.
type Q interface {
	// Get retrieves a message from the queue.
	Get(timeout time.Duration) (*Message, error)

	// Put sends a message to the queue after the given delay. When delay
	// is 0, the message is sent immediately to the queue.
	Put(m *Message, delay time.Duration) error

	// Acknowledge that the message was successfully received from the queue.
	Delete(m *Message, tag uint64, multiple bool) error

	// Release sends a Not Acknowledged message in the queue.When the requeue
	// flag is true, the messaged is released again.
	//
	// This method should be used when handling a message that you cannot
	// handle, maximizing throughput.
	Release(m *Message, tag uint64, multiple bool, requeue bool) error
}

// Handler represents a runnable routine. It can be started and stopped.
type Handler interface {
	// Start starts the handler. It must be safe to call this function
	// multiple times, even if the handler is already running.
	Start()

	// Stop sends a signal to stop the handler, it won't stop the handler
	// immediately. After calling Stop, one should call Wait for blocking
	// until the handler is stopped.
	//
	// This method will return an error if the handler is not running.
	Stop() error

	// Wait blocks until the handler actually stops.
	Wait()
}

// QFactory manages queues. It's able to create new queue and handler
// instances.
type QFactory interface {
	// Get returns a queue instance, identified by the given name.
	Get(name string) (Q, error)

	// Handler returns a handler for the given queue names. Once the
	// handler is started (after calling Start method), it will call f
	// whenever a new message arrives in one of the given queue names.
	Handler(f func(*Message), name ...string) (Handler, error)
}

var factories = map[string]QFactory{
	"rabbitmq": rabbitmqFactory{},
}

// Register registers a new queue factory. This is how one would add a new
// queue to megam.
func Register(name string, factory QFactory) {
	factories[name] = factory
}

// Factory returns an instance of the QFactory used in megam. It reads gulpd
// configuration to find the currently used queue system (for example,
// rabbitmq) and returns an instance of the configured system, if it's
// registered. Otherwise it will return an error.
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

// Delete deletes the message from the queue.
func (m *Message) Delete() {
	m.delete = true
}
