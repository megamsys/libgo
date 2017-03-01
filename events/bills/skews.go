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
	log "github.com/Sirupsen/logrus"
	"github.com/megamsys/libgo/api"
	"github.com/megamsys/libgo/pairs"
	constants "github.com/megamsys/libgo/utils"
	"strconv"
	"time"
)

const (
	EVENTSKEWS          = "/eventsskews"
	EVENTSKEWS_NEW      = "/eventsskews/content"
	EVENTEVENTSKEWSJSON = "Megam::Skews"
	HARDSKEWS           = "terminate"
	SOFTSKEWS           = "suspend"
	WARNING             = "warning"
	ACTIVE              = "active"
)

type ApiSkewsEvents struct {
	JsonClaz string        `json:"json_claz"`
	Results  []EventsSkews `json:"results"`
}
type EventsSkews struct {
	Id        string          `json:"id"`
	AccountId string          `json:"account_id"`
	CatId     string          `json:"cat_id"`
	Inputs    pairs.JsonPairs `json:"inputs"`
	Outputs   pairs.JsonPairs `json:"outputs"`
	Actions   pairs.JsonPairs `json:"actions"`
	JsonClaz  string          `json:"json_claz"`
	Status    string          `json:"status"`
	EventType string          `json:"event_type"`
}

func NewEventsSkews(email, cat_id string, mi map[string]string) ([]EventsSkews, error) {

	if email == "" {
		return nil, fmt.Errorf("account_id should not be empty")
	}

	args := api.NewArgs(mi)
	args.Email = email
	cl := api.NewClient(args, EVENTSKEWS+"/"+cat_id)
	response, err := cl.Get()
	if err != nil {
		return nil, err
	}

	ac := &ApiSkewsEvents{}
	err = json.Unmarshal(response, ac)
	if err != nil {
		return nil, err
	}
	return ac.Results, nil
}

func (s *EventsSkews) CreateEvent(o *BillOpts, ACTION string, mi map[string]string) error {
	var exp_at, gen_at time.Time
	var action, next string
	mm := make(map[string][]string, 0)
	if s.Inputs != nil {
		gen_at, _ = time.Parse(time.RFC3339, s.Inputs.Match(constants.ACTION_TRIGGERED_AT))
	} else {
		gen_at = time.Now()
	}

	softDue, err := time.ParseDuration(o.SoftGracePeriod)
	hardDue, err := time.ParseDuration(o.HardGracePeriod)
	if err != nil {
		return err
	}
	switch ACTION {
	case HARDSKEWS:
		exp_at = gen_at.Add(hardDue)
		action = HARDSKEWS
		next = "unrecoverable"
	case SOFTSKEWS:
		exp_at = gen_at.Add(hardDue)
		action = SOFTSKEWS
		next = HARDSKEWS
	case WARNING:
		mm[constants.ACTION_TRIGGERED_AT] = []string{gen_at.Format(time.RFC3339)}
		exp_at = gen_at.Add(softDue)
		action = WARNING
		next = SOFTSKEWS
	}
	mm[constants.NEXT_ACTION_DUE_AT] = []string{exp_at.Format(time.RFC3339)}
	mm[constants.ACTION] = []string{action}
	mm[constants.NEXT_ACTION] = []string{next}

	s.Inputs.NukeAndSet(mm)
	s.Status = ACTIVE
	return s.Create(mi)
}

func (s *EventsSkews) Create(mi map[string]string) error {
	args := api.NewArgs(mi)
	args.Email = s.AccountId
	cl := api.NewClient(args, EVENTSKEWS_NEW)
	_, err := cl.Post(s)
	if err != nil {
		return err
	}
	return nil
}

func (s *EventsSkews) ActionEvents(o *BillOpts, currentBal string, mi map[string]string) error {
	log.Debugf("checks skews actions for ondemand")
	sk := make(map[string]*EventsSkews, 0)
	// to get skews events for that particular cat_id/ asm_id
	evts, err := NewEventsSkews(o.AccountId, o.AssemblyId, mi)
	if err != nil {
		return err
	}
	for _, v := range evts {
		if v.Status == ACTIVE {
			sk[v.Actions.Match(constants.ACTION)] = &v
		}
	}
	ACTION := s.action(o, currentBal)
	if len(sk) > 0 {
		if sk[ACTION] != nil {
			switch true {
			case ACTION == HARDSKEWS && sk[HARDSKEWS].isExpired():
				return sk[HARDSKEWS].CreateEvent(o, HARDSKEWS, mi)
			case ACTION == SOFTSKEWS && sk[SOFTSKEWS].isExpired():
				return sk[SOFTSKEWS].CreateEvent(o, HARDSKEWS, mi)
			case ACTION == WARNING && sk[WARNING].isExpired():
				return sk[SOFTSKEWS].CreateEvent(o, SOFTSKEWS, mi)
			}
			return nil
		} else {
			return s.CreateEvent(o, ACTION, mi)
		}
	} else {
		return s.CreateEvent(o, ACTION, mi)
	}
	return nil
}

func (s *EventsSkews) SkewsQuotaUnpaid(o *BillOpts, mi map[string]string) error {
	log.Debugf("checks skews actions for ondemand")
	actions := make(map[string]string, 0)
	sk := make(map[string]*EventsSkews, 0)
	// to get skews events for that particular cat_id/ asm_id
	evts, err := NewEventsSkews(o.AccountId, o.AssemblyId, mi)
	if err != nil {
		return err
	}
	for _, v := range evts {
		if v.Status == ACTIVE {
			sk[v.Actions.Match(constants.ACTION)] = &v
			actions[v.Actions.Match(constants.ACTION)] = ACTIVE
		}
	}
	if len(sk) > 0 {
		switch true {
		case actions[HARDSKEWS] == ACTIVE && sk[HARDSKEWS].isExpired():
			return sk[HARDSKEWS].CreateEvent(o, HARDSKEWS, mi)
		case actions[SOFTSKEWS] == ACTIVE && sk[SOFTSKEWS].isExpired():
			return sk[SOFTSKEWS].CreateEvent(o, HARDSKEWS, mi)
		case actions[WARNING] == ACTIVE && sk[WARNING].isExpired():
			return sk[SOFTSKEWS].CreateEvent(o, SOFTSKEWS, mi)
		}
	}

	return s.CreateEvent(o, WARNING, mi)
}

func (s *EventsSkews) action(o *BillOpts, currentBal string) string {
	cb, _ := strconv.ParseFloat(currentBal, 64)
	slimit, _ := strconv.ParseFloat(o.SoftLimit, 64)
	hlimit, _ := strconv.ParseFloat(o.HardLimit, 64)
	if cb <= hlimit {
		return HARDSKEWS
	} else if cb <= slimit {
		return SOFTSKEWS
	}
	return WARNING
}

func (s *EventsSkews) isExpired() bool {
	t1, _ := time.Parse(time.RFC3339, s.Inputs.Match("generated_at"))
	t2, _ := time.Parse(time.RFC3339, s.Inputs.Match("next_due_at"))
	duration := t2.Sub(t1)
	return t1.Add(duration).Sub(time.Now()) < time.Minute
}
