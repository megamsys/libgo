package dockerapi

import (
	"fmt"
	"strings"
	log "code.google.com/p/log4go"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
	)

type RawRequest struct {
	Method       string
	Url          string
	
}


func NewRawRequest(method, url) *RawRequest {
	return &RawRequest{
		Method:       method,
		Url:          url,
		
	}
}

func (c *Client) List(name, string) (*RawResponse, error) {
	url := "http://"+c.Host+":"+c.Port+"/images/"+name+"/json"
	req := NewRawRequest("GET", url)
	resp, err := c.SendRequest(req, "")
	return resp, err
	
	
	
func (c *Client) SendRequest(rr *RawRequest, json string) (*RawResponse, error) {	
	
	   httpPath = rr.Url

		// Return a cURL command if curlChan is set
        c.OpenCURL()
		if c.cURLch != nil {
			command := fmt.Sprintf("curl")
			if rr.Method != "" {
				command += fmt.Sprintf(" -X %s", rr.Method)
			 }
			if rr.Url != "" {
				command += fmt.Sprintf(" %s", rr.Url)
			 }
			
			
			 }
            log.Info(command)
			c.sendCURL(command)
		}
	resp, err = c.httpClient.Do(req)
		r := &RawResponse{
		StatusCode: resp.StatusCode,
		Body:       respBody,
		Header:     resp.Header,
	}

	return r, nil
}
	