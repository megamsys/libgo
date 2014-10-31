package geard

import (
	"errors"
	"fmt"
	"strings"
	log "code.google.com/p/log4go"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// Errors introduced by handling requests
var (
	ErrRequestCancelled = errors.New("sending request is cancelled")
)

type RawRequest struct {
	Method       string
	Url          string
	ContentType         string
	Cancel       <-chan bool
}

// NewRawRequest returns a new RawRequest
func NewRawRequest(method, url, contenttype string, cancel <-chan bool) *RawRequest {
	return &RawRequest{
		Method:       method,
		Url:          url,
		ContentType:  contenttype,
		Cancel:       cancel,
	}
}

func (c *Client) Install(name, json string) (*RawResponse, error) {
	url := "http://"+c.Host+":"+c.Port+"/container/"+name
	req := NewRawRequest("PUT", url, "Content-Type: application/json", nil)
	resp, err := c.SendRequest(req, json)
	return resp, err
}

func (c *Client) Start(name string) (*RawResponse, error) {
	url := "http://"+c.Host+":"+c.Port+"/container/"+name+"/started"
	req := NewRawRequest("PUT", url, "Content-Type: application/json", nil)
	resp, err := c.SendRequest(req, "")
	return resp, err
}

func (c *Client) Stop(name string) (*RawResponse, error) {
	url := "http://"+c.Host+":"+c.Port+"/container/"+name+"/stopped"
	req := NewRawRequest("PUT", url, "Content-Type: application/json", nil)
	resp, err := c.SendRequest(req, "")
	return resp, err
}

func (c *Client) Restart(name string) (*RawResponse, error) {
	url := "http://"+c.Host+":"+c.Port+"/container/"+name+"/restart"
	req := NewRawRequest("POST", url, "Content-Type: application/json", nil)
	resp, err := c.SendRequest(req, "")
	return resp, err
}


// SendRequest sends a HTTP request and returns a Response as defined by etcd
func (c *Client) SendRequest(rr *RawRequest, json string) (*RawResponse, error) {

	var req *http.Request
	var resp *http.Response
	var httpPath string
	var err error
	var respBody []byte

	var numReqs = 1
        log.Info("------------entry------")
	checkRetry := c.CheckRetry
	if checkRetry == nil {
		checkRetry = DefaultCheckRetry
	}

	cancelled := make(chan bool, 1)
	reqLock := new(sync.Mutex)

	if rr.Cancel != nil {
		cancelRoutine := make(chan bool)
		defer close(cancelRoutine)

		go func() {
			select {
			case <-rr.Cancel:
				cancelled <- true
				fmt.Println("send.request is cancelled")
			case <-cancelRoutine:
				return
			}

			// Repeat canceling request until this thread is stopped
			// because we have no idea about whether it succeeds.
			for {
				reqLock.Lock()
				c.httpClient.Transport.(*http.Transport).CancelRequest(req)
				reqLock.Unlock()

				select {
				case <-time.After(100 * time.Millisecond):
				case <-cancelRoutine:
					return
				}
			}
		}()
	}

	    log.Info("Connecting to geard deamon: attempt")

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
			
			if json != "" {
				command += fmt.Sprintf(" -d %s", json)
			 }
            log.Info(command)
			c.sendCURL(command)
		}

		reqLock.Lock()
 
		if json == "" {
			if req, err = http.NewRequest(rr.Method, httpPath, nil); err != nil {
				return nil, err
			}
		} else {
			body := strings.NewReader(json)
			if req, err = http.NewRequest(rr.Method, httpPath, body); err != nil {
				return nil, err
			}

			req.Header.Set("Content-Type", rr.ContentType)
		}
		reqLock.Unlock()
		resp, err = c.httpClient.Do(req)
		defer func() {
			if resp != nil {
				resp.Body.Close()
			}
		}()

		// If the request was cancelled, return ErrRequestCancelled directly
		select {
		case <-cancelled:
			return nil, ErrRequestCancelled
		default:
		}

		numReqs++
               
		// network error, change a machine!
		if err != nil {
			log.Error("network error:", err.Error())
			lastResp := http.Response{}
			if checkErr := checkRetry(numReqs, lastResp, err); checkErr != nil {
				return nil, checkErr
			}

			//c.cluster.switchLeader(attempt % len(c.cluster.Machines))
			//continue
		}

		// if there is no error, it should receive response
		log.Error("recv.response.from", httpPath)
                log.Info(resp)
		if validHttpStatusCode[resp.StatusCode] {
			// try to read byte code and break the loop
                        log.Info("--------if entry----------")
                         log.Info(resp.StatusCode)
			respBody, err = ioutil.ReadAll(resp.Body)
			if err == nil {
				log.Error("recv.success.", httpPath)
				//break
			}
			// ReadAll error may be caused due to cancel request
			select {
			case <-cancelled:
				return nil, ErrRequestCancelled
			default:
			}
		}

		// if resp is TemporaryRedirect, set the new leader and retry
		if resp.StatusCode == http.StatusTemporaryRedirect {
			u, err := resp.Location()

			if err != nil {
				log.Error(err)
			} else {
				// Update cluster leader based on redirect location
				// because it should point to the leader address
				//c.cluster.updateLeaderFromURL(u)
				log.Error("recv.response.relocate", u.String())
			}
			resp.Body.Close()
			//continue
		}

		if checkErr := checkRetry(numReqs, *resp,
			errors.New("Unexpected HTTP status code")); checkErr != nil {
			return nil, checkErr
		}
		resp.Body.Close()
	//}

	r := &RawResponse{
		StatusCode: resp.StatusCode,
		Body:       respBody,
		Header:     resp.Header,
	}

	return r, nil
}

// DefaultCheckRetry defines the retrying behaviour for bad HTTP requests
// If we have retried 2 * machine number, stop retrying.
// If status code is InternalServerError, sleep for 200ms.
func DefaultCheckRetry(numReqs int, lastResp http.Response,
	err error) error {

	code := lastResp.StatusCode
	if code == http.StatusInternalServerError {
		time.Sleep(time.Millisecond * 200)

	}

	log.Error("bad response status code", code)
	return nil
}

