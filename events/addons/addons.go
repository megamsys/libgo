package addons

import (
	log "github.com/Sirupsen/logrus"
	"github.com/megamsys/libgo/api"
	constants "github.com/megamsys/libgo/utils"
	"encoding/json"
	"io/ioutil"
	"github.com/megamsys/libgo/events/alerts"
	"time"
	"fmt"
)

const (
	ADDONS_NEW = "/addons/content"
	GETADDONS = "/addons/"
	PROVIDER_NAME = "provider_name"
	PROVIDER_ID = "provider_id"
	)

type Addons struct {
	Id           string   `json:"id" cql:"id"`
	ProviderName string   `json:"provider_name" cql:"provider_name"`
	ProviderId   string   `json:"provider_id" cql:"provider_id"`
	AccountId    string   `json:"account_id" cql:"account_id"`
	Options      []string `json:"options" cql:"options"`
	CreatedAt    string   `json:"created_at" cql:"created_at"`
}

type ApiAddons struct {
	JsonClaz  string `json:"json_claz"`
	Results  []Addons `json:"results"`
}

func NewAddons(edata alerts.EventData) *Addons {
	return &Addons{
		Id: "",
		ProviderName: edata.M[PROVIDER_NAME],
		ProviderId: edata.M[PROVIDER_ID],
		AccountId: edata.M[constants.ACCOUNT_ID],
		Options: edata.D,
		CreatedAt: time.Now().String(),
	}
}

func (s *Addons) Onboard(m map[string]string) error {
	args := api.NewArgs(m)
	cl := api.NewClient(args, ADDONS_NEW)
	_, err := cl.Post(s)
	if err != nil {
		log.Debugf(err.Error())
		return err
	}
	return nil
}

func (s *Addons) Get(m map[string]string) error {
	// Here skips balances fetching for the VMs which is launched on opennebula,
	// that does not have records on vertice database
	if s.AccountId == "" {
	 return fmt.Errorf("account_id should not be empty")
	}
	args := api.NewArgs(m)
	cl := api.NewClient(args, GETADDONS + s.Id)
	response, err := cl.Get()
	if err != nil {
		return err
	}
	htmlData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	o := &ApiAddons{}
	err = json.Unmarshal(htmlData, o)
	if err != nil {
		return err
	}
	s = &o.Results[0]
	return nil
}
