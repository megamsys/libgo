package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Context struct {
	username   string
	api_key    string
	master_key string
	password   string
}

type Client struct {
	HTTPClient     *http.Client
	context        *Context
	Authly         *Authly
	progname       string
	currentVersion string
	versionHeader  string
}

func NewClient(client *http.Client, data map[string]string) *Client {
	//
	ctx := context(map[string]string{"email": "info@megam.io", "api_key": "fakeapikey"}) // context(data.GetKeys())
	return &Client{
		HTTPClient:     client,
		context:        ctx,
		Authly:         NewAuthly(),
		progname:       "sample",
		currentVersion: "2",
		versionHeader:  "Supported-Gulp",
	}
}

func context(m map[string]string) *Context {
	return &Context{
		username:   m["email"],
		api_key:    m["api_key"],
		password:   m["password"],
		master_key: m["master_key"],
	}
}

func (c *Client) detectClientError(err error) error {
	return fmt.Errorf("Failed to connect to api server.")
}

func (c *Client) Do(request *http.Request) (*http.Response, error) {
	fmt.Println("Header :")
	for headerKey, headerVal := range c.Authly.AuthMap {
		fmt.Println("Key :", headerKey, "  Value :", headerVal)
		request.Header.Add(headerKey, headerVal)
		//request.Header[headerKey] = []string{headerVal}
	}

	request.Close = true
	response, err := c.HTTPClient.Do(request)

	if err != nil {
		return nil, err
	}
	supported := response.Header.Get(c.versionHeader)
	format := `################################################################

WARNING: You're using an unsupported version of %s.

You must have at least version %s, your current
version is %s.

################################################################

`
	if !validateVersion(supported, c.currentVersion) {
		fmt.Println(format)
		fmt.Println(supported)
	}
	if response.StatusCode > 399 {
		defer response.Body.Close()
		result, _ := ioutil.ReadAll(response.Body)
		return response, errors.New(string(result))
	}
	return response, nil

}

// validateVersion checks whether current version is greater or equal to
// supported version.
func validateVersion(supported, current string) bool {
	var (
		bigger bool
		limit  int
	)
	if supported == "" {
		return true
	}
	partsSupported := strings.Split(supported, ".")
	partsCurrent := strings.Split(current, ".")
	if len(partsSupported) > len(partsCurrent) {
		limit = len(partsCurrent)
		bigger = true
	} else {
		limit = len(partsSupported)
	}
	for i := 0; i < limit; i++ {
		current, err := strconv.Atoi(partsCurrent[i])
		if err != nil {
			return false
		}
		supported, err := strconv.Atoi(partsSupported[i])
		if err != nil {
			return false
		}
		if current < supported {
			return false
		}
		if current > supported {
			return true
		}
	}
	if bigger {
		return false
	}
	return true
}
