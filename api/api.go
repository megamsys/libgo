package api

import (
	"net/http"
	"reflect"
	"strings"
	//  "encoding/json"
	//  "bytes"
	"fmt"
)

type VerticeApi interface {
	GetKeys() map[string]string
}

type Credentials struct {
	Email      string
	Api_Key    string
	Master_Key string
	Password   string
	HostUrl    string
}

func (c *Credentials) ToMap() map[string]string {
	keys := make(map[string]string)
	s := reflect.ValueOf(c).Elem()
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
	tmpinp := map[string]string{"email": "info@megam.io", "api_key": "152efc6f782f1be4a2565293fb7a066a7dfbe5c5"} //data.GetKeys()

  urlsuffix := "/accounts/info@megam.io"
	// jsonbody, err := json.Marshal(data)
	// if err != nil {
	// 	return nil, err
	// }

  client.Authly.UrlSuffix = urlsuffix
  client.Authly.Keys = tmpinp
  //client.Authly.JSONBody =  jsonbody
	url := "http://40.74.121.55:9000/v2/accounts/info@megam.io" //GetURL("/auth")
	// if err != nil {
	// 	return nil, err
	// }

	fmt.Println("==> " + url)
	//authly.JSONBody = jsonMsg

	err := client.Authly.AuthHeader()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("GET", url, nil) //bytes.NewReader(jsonMsg)
	if err != nil {
		return nil, err
	}

	return client.Do(request)
}

// func (this *LoginRouter) Login() {
// 	data := &utils.User{this.GetString("username"), this.GetString("password")}
// 	client := utils.NewClient(&http.Client{}, data)
// 	response, _ := this.Auth(client, data)
//     fmt.Println(response)
//     if response != nil {
//     	if response.StatusCode > 399 && response.StatusCode < 498 {
// 		   this.FlashWrite("LoginError", "true")
// 		   this.Redirect("/", 302)
// 	   } else if response.StatusCode > 499 {
// 		   this.FlashWrite("ServerError", "true")
// 		   this.Redirect("/", 302)
// 	   } else {
// 		this.LoginUser(data, true)
// 		this.Redirect("/index", 302)
// 	  }
//     } else {
//     	this.FlashWrite("ServerError", "true")
// 		this.Redirect("/", 302)
//     }
//
// }
