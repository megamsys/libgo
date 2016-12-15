package api

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
)

const (
	EMAIL = "email"
	PASSWORD = "password"
	MASTER_KEY = "master_key"
	API_KEY    = "api_key"
	X_Megam_EMAIL              = "X-Megam-EMAIL"
	X_Megam_MASTERKEY          = "X-Megam-MASTERKEY"
	X_Megam_PUTTUSAVI          = "X-Megam-PUTTUSAVI"
	X_Megam_DATE               = "X-Megam-DATE"
	X_Megam_HMAC               = "X-Megam-HMAC"
	X_Megam_OTTAI              = "X-Megam-OTTAI"
	X_Megam_ORG                = "X-Megam-ORG"
	Content_Type               = "Content-Type"
	Accept                     = "Accept"
	application_vnd_megam_json = "application/vnd.megam+json"
)


func CalcHMAC(secret string, message string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	sumh := h.Sum(nil)

	sumi := make([]string, len(sumh))
	for i := 0; i < len(sumh); i++ {
		sumi[i] = ("00" + fmt.Sprintf("%x", (sumh[i]&0xff)))
		sumi[i] = sumi[i][len(sumi[i])-2 : len(sumi[i])]
	}
	outs := strings.Join(sumi, "")
	return outs
}

func GetMD5Hash(text []byte) string {
	hasher := md5.New()
	hasher.Write(text)
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}

// func GetMD5Hash(text []byte) string {
// 	hasher := md5.New()
// 	hasher.Write(text)
// 	return CalcBase64(string(hasher.Sum(nil)))
// }

func CalcBase64(data string) string {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	encoder.Write([]byte(data))
	encoder.Close()
	return buf.String()
}

func GetURL(path string) (string, error) {
	var prefix string
	//	target, err := config.GetString("api:server")
	//	if err != nil {
	//		return "", err
	//	}
	target :=  "localhost"  //config.GetString("api:host")
	//API_GATEWAY_VERSION, _ := config.GetString("api:version")
	if m, _ := regexp.MatchString("^https?://", target); !m {
		prefix = "http://"
	}
	return prefix + strings.TrimRight(target, "/") + strings.TrimRight("v2", "/") + path, nil
}

/*
type ServiceModel struct {
	Service   string
	Instances []string
}

func ShowServicesInstancesList(b []byte) ([]byte, error) {
	var services []ServiceModel
	err := json.Unmarshal(b, &services)
	if err != nil {
		return []byte{}, err
	}
	if len(services) == 0 {
		return []byte{}, nil
	}
	table := NewTable()
	table.Headers = Row([]string{"Services", "Instances"})
	for _, s := range services {
		insts := strings.Join(s.Instances, ", ")
		r := Row([]string{s.Service, insts})
		table.AddRow(r)
	}
	return table.Bytes(), nil
}*/
