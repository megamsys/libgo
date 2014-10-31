package cmd

import (
	"bytes"
	ttesting "github.com/megamsys/libgo/cmd/testing"
	"gopkg.in/check.v1"
	"net/http"
)

func (s *S) TestShouldSetCloseToTrue(c *check.C) {
	request, err := http.NewRequest("GET", "/", nil)
	c.Assert(err, check.IsNil)
	transport := ttesting.Transport{
		Status:  http.StatusOK,
		Message: "OK",
	}
	client := NewClient(&http.Client{Transport: &transport}, nil, manager)
	client.Do(request)
	c.Assert(request.Close, check.Equals, true)
}

func (s *S) TestShouldReturnBodyMessageOnError(c *check.C) {
	request, err := http.NewRequest("GET", "/", nil)
	c.Assert(err, check.IsNil)
	client := NewClient(&http.Client{Transport: &ttesting.Transport{Message: "You must be authenticated to execute this command.", Status: http.StatusUnauthorized}}, nil, manager)
	response, err := client.Do(request)
	c.Assert(response, check.NotNil)
	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "You must be authenticated to execute this command.")
}



