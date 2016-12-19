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
	"time"
	"github.com/megamsys/libgo/api"
	"gopkg.in/yaml.v2"
	log "github.com/Sirupsen/logrus"
)

const (
	NEWTRANSACTION = "/billedhistories/content"
	BILLJSONCLAZ      = "Megam::Billedhistories"
)

type BillTransactionOpts struct {
	AccountId    string
	AssemblyId   string
	AssemblyName string
	Consumed     string
}

type BillTransaction struct {
	Id            string `json:"id" cql:"id"`
	AccountId     string `json:"account_id" cql:"account_id"`
	AssemblyId    string `json:"assembly_id" cql:"assembly_id"`
	BillType      string `json:"bill_type" cql:"bill_type"`
	BillingAmount string `json:"billing_amount" cql:"billing_amount"`
	CurrencyType  string `json:"currency_type" cql:"currency_type"`
	JsonClaz      string `json:"json_claz" cql:"json_claz"`
	CreatedAt     time.Time `json:"created_at" cql:"created_at"`
}

func (bt *BillTransactionOpts) String() string {
	if d, err := yaml.Marshal(bt); err != nil {
		return err.Error()
	} else {
		return string(d)
	}
}

func NewBillTransaction(topts *BillOpts) (*BillTransaction, error) {
	return &BillTransaction{
		AccountId:     topts.AccountId,
		AssemblyId:    topts.AssemblyId,
		BillType:      "VM",
		BillingAmount: topts.Consumed,
		JsonClaz: BILLJSONCLAZ,
		CurrencyType:  "",
		CreatedAt:     time.Now(),
	}, nil
}

func (bt *BillTransaction) Transact(m map[string]string) error {
	args := api.NewArgs(m)
	args.Path = NEWTRANSACTION
	cl := api.NewClient(args)
	_, err := cl.Post(bt)
	if err != nil {
		log.Debugf(err.Error())
		return err
	}
	return nil
}
