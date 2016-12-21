package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"github.com/megamsys/libgo/utils"
	log "github.com/Sirupsen/logrus"
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

func NewArgs(args map[string]string) ApiArgs {
	return ApiArgs{
		Email:      args[utils.USERMAIL],
		Api_Key:    args[utils.API_KEY],
		Master_Key: args[utils.MASTER_KEY],
		Password:   args[utils.PASSWORD],
		Org_Id:     args[utils.ORG_ID],
		Url:        args[utils.API_URL],
	}
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
		fmt.Println("Request [GET] ==> " + c.Url)
	return c.run("GET")
}

func (c *Client) Post(data interface{}) (*http.Response, error) {
	jsonbody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	c.Authly.JSONBody = jsonbody
	fmt.Println("Request [POST] ==> " + c.Url)
	log.Debugf("[Body]  (%s)",string(jsonbody))
 return c.run("POST")
}

func (c *Client) Delete() (*http.Response, error) {
	fmt.Println("Request [DELETE] ==> " + c.Url)
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

		res, err := c.Do(request)
		if err != nil {
			fmt.Println("  api error :",err)
			return nil, err
		}

		return res, nil
}
