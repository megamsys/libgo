package alerts

import (
	"os"
    "testing"
	"gopkg.in/check.v1"
)

var st = os.Getenv("NIL_SLACK_TOKEN")
var ch = "ahoy"

func Test(t *testing.T) { check.TestingT(t) }

type S struct{}

var _ = check.Suite(&S{})

func (s *S) TestSlack(c *check.C) {
	if st == "" {
		c.Skip("-Slack (token) not provided")
	}
	c.Assert(len(st) > 0, check.Equals, true)
	ms := NewSlack(map[string]string{TOKEN: st, CHANNEL: ch})
	c.Assert(ms, check.NotNil)
	err := ms.Notify(LAUNCHED, EventData{ M: map[string]string{"message": "Awesome vertice... :)"}})
	c.Assert(err, check.IsNil)
}
