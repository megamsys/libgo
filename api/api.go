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

type ApiOrgs struct {
	Email      string
	Api_Key    string
	Master_Key string
	Password   string
	Org_Id     string
	Url        string
	Path       string
}

func (c ApiOrgs) ToMap() map[string]string {
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

func (client *Client) Get() (*http.Response, error) {
	url := client.GetURL()
	fmt.Println("==> " + url)
	err := client.Authly.AuthHeader()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return client.Do(request)
}

func (c *Client) Post(data interface{}) (*http.Response, error) {
	jsonbody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	c.Authly.JSONBody = jsonbody
	url := c.GetURL()
	fmt.Println("==> " + url)

	err = c.Authly.AuthHeader()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", url, bytes.NewReader(c.Authly.JSONBody))
	if err != nil {
		return nil, err
	}

	return c.Do(request)
}
