package bills

import (
	"github.com/megamsys/libgo/utils"
	"gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

type S struct {
	m map[string]string
}

var _ = check.Suite(&S{})

func (s *S) SetUpSuite(c *check.C) {
	m := make(map[string]string, 0)
	m[utils.MASTER_KEY] = "3b8eb672aa7c8db82e5d34a0744740b20ed59e1f6814cfb63364040b0994ee3f"
	m[utils.API_URL] = "http://188.240.231.85:8999/v2"
	s.m = m
}
