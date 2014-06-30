package amqp

import (
	"github.com/megamsys/libgo/safe"
	"gopkg.in/check.v1"
	"time"
)

type ExecutorSuite struct{}

var _ = check.Suite(&ExecutorSuite{})

func dumb() {
	time.Sleep(1e3)
}

func (s *ExecutorSuite) TestStart(c *check.C) {
	var ct safe.Counter
	h1 := executor{inner: func() { ct.Increment() }}
	h1.Start()
	c.Assert(h1.state, check.Equals, running)
	h1.Stop()
	h1.Wait()
	c.Assert(ct.Val(), check.Not(check.Equals), 0)
}

func (s *ExecutorSuite) TestPreempt(c *check.C) {
	h1 := executor{inner: dumb}
	h2 := executor{inner: dumb}
	h3 := executor{inner: dumb}
	h1.Start()
	h2.Start()
	h3.Start()
	Preempt()
	c.Assert(h1.state, check.Equals, stopped)
	c.Assert(h2.state, check.Equals, stopped)
	c.Assert(h3.state, check.Equals, stopped)
}

func (s *ExecutorSuite) TestStopNotRunningExecutor(c *check.C) {
	h := executor{inner: dumb}
	err := h.Stop()
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "Not running.")
}

func (s *ExecutorSuite) TestExecutorImplementsHandler(c *check.C) {
	var _ Handler = &executor{}
}
