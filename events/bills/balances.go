/*
** Copyright [2013-2016] [Megam Systems]
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
 */
package bills

import (
	"encoding/json"
	"fmt"
	"github.com/megamsys/libgo/api"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
	"time"
)

const (
	UPDATEBALANCES   = "/balances/update"
	GETBALANCE       = "/balances/"
	EVENTBALANCEJSON = "Megam::Balances"
)

type BalanceOpts struct {
	Id       string
	Consumed string
}

type ApiBalances struct {
	JsonClaz string     `json:"json_claz" cql:"json_claz"`
	Results  []Balances `json:"results" json:"results"`
}
type Balances struct {
	Id        string    `json:"id" cql:"id"`
	AccountId string    `json:"account_id" cql:"account_id"`
	Credit    string    `json:"credit" cql:"credit"`
	CreatedAt time.Time `json:"created_at" cql:"created_at"`
	UpdatedAt time.Time `json:"updated_at" cql:"updated_at"`
	JsonClaz  string    `json:"json_claz" cql:"json_claz"`
}

func (b *Balances) String() string {
	if d, err := yaml.Marshal(b); err != nil {
		return err.Error()
	} else {
		return string(d)
	}
}

//Temporary hack to create an assembly from its id.
//This is used by SetStatus.
//We need add a Notifier interface duck typed by Box and Carton ?
func NewBalances(id string, m map[string]string) (*Balances, error) {
	// Here skips balances fetching for the VMs which is launched on opennebula,
	// that does not have records on vertice database
	if id == "" {
		return nil, fmt.Errorf("account_id should not be empty")
	}

	args := api.NewArgs(m)
	cl := api.NewClient(args, GETBALANCE+id)
	response, err := cl.Get()
	if err != nil {
		return nil, err
	}
	htmlData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	ac := &ApiBalances{}
	err = json.Unmarshal(htmlData, ac)
	if err != nil {
		return nil, err
	}
	b := &ac.Results[0]
	return b, nil
}

func (b *Balances) Deduct(bopts *BalanceOpts, m map[string]string) error {
	avail, err := strconv.ParseFloat(b.Credit, 64)
	if err != nil {
		return err
	}

	consume, cerr := strconv.ParseFloat(bopts.Consumed, 64)
	if cerr != nil {
		return cerr
	}

	b.UpdatedAt = time.Now()
	b.Credit = strconv.FormatFloat(avail-consume, 'f', 2, 64)

	args := api.NewArgs(m)
	cl := api.NewClient(args, UPDATEBALANCES)
	_, err = cl.Post(b)
	if err != nil {
		return err
	}
	return nil
}
