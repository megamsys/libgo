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
	log "github.com/Sirupsen/logrus"
	"github.com/megamsys/libgo/api"
	"github.com/megamsys/libgo/utils"
	"gopkg.in/yaml.v2"
	//"strconv"
	"time"
)

const (
	NEWTRANSACTION = "/billedhistories/content"
	BILLJSONCLAZ   = "Megam::Billedhistories"
)

type BillTransactionOpts struct {
	AccountId    string
	AssemblyId   string
	AssemblyName string
	Consumed     string
}

type BillTransaction struct {
	AccountId     string    `json:"-" cql:"account_id"`
	AssemblyId    string    `json:"assembly_id" cql:"assembly_id"`
	BillType      string    `json:"bill_type" cql:"bill_type"`
	BillingAmount string    `json:"billing_amount" cql:"billing_amount"`
	StateDate     time.Time `json:"start_date" cql:"start_date"`
	EndDate       time.Time `json:"end_date" cql:"end_date"`
	CurrencyType  string    `json:"currency_type" cql:"currency_type"`
}

func (bt *BillTransactionOpts) String() string {
	if d, err := yaml.Marshal(bt); err != nil {
		return err.Error()
	} else {
		return string(d)
	}
}

func NewBillTransaction(topts *BillOpts) (*BillTransaction, error) {
	//start, _ := strconv.ParseInt(topts.StartTime, 10, 64)
	//end, _ := strconv.ParseInt(topts.EndTime, 10, 64)
	return &BillTransaction{
		AccountId:     topts.AccountId,
		AssemblyId:    topts.AssemblyId,
		BillType:      "VM",
		BillingAmount: topts.Consumed,
		StateDate:     time.Now(), //time.Unix(start, 0),
		EndDate:       time.Now(), //time.Unix(end, 0),
		CurrencyType:  "",
	}, nil
}

func (bt *BillTransaction) Transact(m map[string]string) error {
	m[utils.USERMAIL] = bt.AccountId
	args := api.NewArgs(m)
	cl := api.NewClient(args, NEWTRANSACTION)
	_, err := cl.Post(bt)
	if err != nil {
		log.Debugf(err.Error())
		return err
	}
	return nil
}
