package etcd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"sync"
	"time"
)

// Errors introduced by handling requests
var (
	ErrRequestCancelled = errors.New("sending request is cancelled")
)

type RawRequest struct {
	Method       string
	RelativePath string
	Values       url.Values
	Cancel       <-chan bool
}

// NewRawRequest returns a new RawRequest
func NewRawRequest(method, relativePath string, values url.Values, cancel <-chan bool) *RawRequest {
	return &RawRequest{
		Method:       method,
		RelativePath: relativePath,
		Values:       values,
		Cancel:       cancel,
	}
}

// getCancelable issues a cancelable GET request
func (c *Client) getCancelable(key string, options Options,
	cancel <-chan bool) (*RawResponse, error) {
	p := keyToPath(key)

	str, err := options.toParameters(VALID_GET_OPTIONS)
	if err != nil {
		return nil, err
	}
	p += str

	req := NewRawRequest("GET", p, nil, cancel)
	resp, err := c.SendRequest(p, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// get issues a GET request
func (c *Client) get(key string, options Options) (*RawResponse, error) {
	return c.getCancelable(key, options, nil)
}

// put issues a PUT request
func (c *Client) put(key string, value string, options Options) (*RawResponse, error) {
        
	logger.Debugf("put %s, %s", key, value)
	p := keyToPath(key)
        fmt.Println(p)

	str, err := options.toParameters(VALID_PUT_OPTIONS)
	if err != nil {
		return nil, err
	}
	p += str
        fmt.Println(p)

	req := NewRawRequest("PUT", p, buildValues(value), nil)
         fmt.Println(req)
	resp, err := c.SendRequest(p, req)
        fmt.Println(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// post issues a POST request
func (c *Client) post(key string, value string) (*RawResponse, error) {
	logger.Debugf("post %s, %s, ttl: %d", key, value)
	p := keyToPath(key)

	req := NewRawRequest("POST", p, buildValues(value), nil)
	resp, err := c.SendRequest(p, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// delete issues a DELETE request
func (c *Client) delete(key string, options Options) (*RawResponse, error) {
	logger.Debugf("delete %s", key)
	p := keyToPath(key)

	str, err := options.toParameters(VALID_DELETE_OPTIONS)
	if err != nil {
		return nil, err
	}
	p += str

	req := NewRawRequest("DELETE", p, nil, nil)
	resp, err := c.SendRequest(p, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SendRequest sends a HTTP request and returns a Response as defined by etcd
func (c *Client) SendRequest(path string, rr *RawRequest) (*RawResponse, error) {

	var req *http.Request
	var resp *http.Response
	var httpPath string
	var err error
	var respBody []byte

	var numReqs = 1
        fmt.Println("------------entry------")
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
				logger.Debug("send.request is cancelled")
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

	// If we connect to a follower and consistency is required, retry until
	// we connect to a leader
	//sleep := 25 * time.Millisecond
	//maxSleep := time.Second
	//for attempt := 0; ; attempt++ {
		/*if attempt > 0 {
			select {
			case <-cancelled:
				return nil, ErrRequestCancelled
			case <-time.After(sleep):
				sleep = sleep * 2
				if sleep > maxSleep {
					sleep = maxSleep
				}
			}
		}*/
                fmt.Println("Connecting to etcd: attempt")
		//logger.Debug("Connecting to etcd: attempt", attempt+1, "for", rr.RelativePath)

		httpPath = c.url + version + "/" + path
		

		// Return a cURL command if curlChan is set
                c.OpenCURL()
		if c.cURLch != nil {
			command := fmt.Sprintf("curl -X %s %s", rr.Method, httpPath)
			for key, value := range rr.Values {
				command += fmt.Sprintf(" -d %s=%s", key, value[0])
			}
                        fmt.Println(command)
			c.sendCURL(command)
		}

		logger.Debug("send.request.to ", httpPath, " | method ", rr.Method)
                 
		reqLock.Lock()
 
		if rr.Values == nil {
			if req, err = http.NewRequest(rr.Method, httpPath, nil); err != nil {
				return nil, err
			}
		} else {
			body := strings.NewReader(rr.Values.Encode())
			if req, err = http.NewRequest(rr.Method, httpPath, body); err != nil {
				return nil, err
			}

			req.Header.Set("Content-Type",
				"application/x-www-form-urlencoded; param=value")
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
			logger.Debug("network error:", err.Error())
			lastResp := http.Response{}
			if checkErr := checkRetry(numReqs, lastResp, err); checkErr != nil {
				return nil, checkErr
			}

			//c.cluster.switchLeader(attempt % len(c.cluster.Machines))
			//continue
		}

		// if there is no error, it should receive response
		logger.Debug("recv.response.from", httpPath)
                fmt.Println(resp)
		if validHttpStatusCode[resp.StatusCode] {
			// try to read byte code and break the loop
                        fmt.Println("--------if entry----------")
                         fmt.Println(resp.StatusCode)
			respBody, err = ioutil.ReadAll(resp.Body)
			if err == nil {
				logger.Debug("recv.success.", httpPath)
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
				logger.Warning(err)
			} else {
				// Update cluster leader based on redirect location
				// because it should point to the leader address
				//c.cluster.updateLeaderFromURL(u)
				logger.Debug("recv.response.relocate", u.String())
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

	logger.Warning("bad response status code", code)
	return nil
}

// buildValues builds a url.Values map according to the given value and ttl
func buildValues(value string) url.Values {
	v := url.Values{}

	if value != "" {
		v.Set("value", value)
	}

	return v
}

// convert key string to http path exclude version
// for example: key[foo] -> path[keys/foo]
// key[/] -> path[keys/]
func keyToPath(key string) string {
	p := path.Join("keys", key)

	// corner case: if key is "/" or "//" ect
	// path join will clear the tailing "/"
	// we need to add it back
	if p == "keys" {
		p = "keys/"
	}

	return p
}
