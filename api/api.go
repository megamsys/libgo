package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type VerticeApi interface {
	ToMap() map[string]string
}

type ApiArgs struct {
	Email      string
	Api_Key    string
	Master_Key string
	Password   string
	Org_Id     string
	Url        string
	Path       string
}

func (c ApiArgs) ToMap() map[string]string {
	keys := make(map[string]string)
	s := reflect.ValueOf(&c).Elem()
	typ := s.Type()
	if s.Kind() == reflect.Struct {
		for i := 0; i < s.NumField(); i++ {
			key := s.Field(i)
			value := s.FieldByName(typ.Field(i).Name)
			switch key.Interface().(type) {
			case string:
				if value.String() != "" {
					keys[strings.ToLower(typ.Field(i).Name)] = value.String()
				}
			}
		}
	}
	return keys
}

func (c *Client) Get() (*http.Response, error) {
		fmt.Println("Request [URL] ==> " + c.Url)
	return c.run("GET")
}

func (c *Client) Post(data interface{}) (*http.Response, error) {
	jsonbody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	c.Authly.JSONBody = jsonbody
	fmt.Println("Request [URL] ==> " + c.Url)
 return c.run("POST")
}

func (c *Client) Delete() (*http.Response, error) {
	fmt.Println("Request [URL] ==> " + c.Url)
 return c.run("DELETE")
}

func (c *Client) run(method string) (*http.Response, error) {
		err := c.Authly.AuthHeader()
		if err != nil {
			return nil, err
		}
		request, err := http.NewRequest(method, c.Url, bytes.NewReader(c.Authly.JSONBody))
		if err != nil {
			return nil, err
		}

		return c.Do(request)
}
