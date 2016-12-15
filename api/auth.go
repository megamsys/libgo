package api

import (
  //"encoding/base64"
  "fmt"
  "time"
)


type Authly struct {
	UrlSuffix string
	Date      string
	JSONBody  []byte
  Keys      map[string]string
	AuthMap   map[string]string
}

func NewAuthly() *Authly {
	return &Authly{
		Date:      time.Now().Format(time.RFC850),
		AuthMap:   map[string]string{},
	}
}

func GetPort() string {
	//port, _ := config.GetString("beego:http_port")
	return "port"
}

func (authly *Authly) AuthHeader() error {
  headMap := make(map[string]string)
  key := ""
	timeStampedPath := authly.Date + "\n" +"/v2" +  authly.UrlSuffix
	md5Body := ""
	if len(authly.JSONBody) > 0 {
		md5Body = GetMD5Hash(authly.JSONBody)
	}
  switch true {
  case (authly.Keys[API_KEY] != ""):
    key = authly.Keys[API_KEY]
  //  headMap[X_Megam_APIKEY] = key
  case (authly.Keys[PASSWORD] != ""):
    key = authly.Keys[EMAIL]
    headMap[X_Megam_PUTTUSAVI] = key
  case (authly.Keys[MASTER_KEY] != ""):
    key = authly.Keys[MASTER_KEY]
    headMap[X_Megam_MASTERKEY] = key
  }
 headMap[X_Megam_ORG] = "ORG846"
//  headMap[X_Megam_ORG] = "ORG8466190968478287925"
fmt.Println("body:  ",CalcBase64(md5Body))
	headMap[X_Megam_DATE] = authly.Date
	headMap[X_Megam_EMAIL] = authly.Keys[EMAIL]
	headMap[Accept] = application_vnd_megam_json
	headMap[X_Megam_HMAC] = authly.Keys[EMAIL] + ":" + CalcHMAC(key, (timeStampedPath+"\n"+ md5Body))
	headMap["Content-Type"] = "application/json"
	authly.AuthMap = headMap
	return nil
}
