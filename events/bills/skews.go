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
	"time"
	"strconv"
	"github.com/megamsys/libgo/api"
  "github.com/megamsys/libgo/pairs"
	log "github.com/Sirupsen/logrus"
  constants "github.com/megamsys/libgo/utils"
)

const (
	EVENTSKEWS       = "/eventsskews"
	EVENTSKEWS_NEW = "/eventsskews/content"
	EVENTEVENTSKEWSJSON = "Megam::Skews"
	HARDSKEWS = "terminate"
	SOFTSKEWS = "suspend"
	WARNING = "warning"
)


type ApiSkewsEvents struct {
	JsonClaz string     `json:"json_claz"`
	Results  []SkewsEvents `json:"results"`
}
type SkewsEvents struct {
	Id        string    `json:"id"`
	AccountId string    `json:"account_id"`
	CatId     string    `json:"cat_id"`
	Inputs    pairs.JsonPairs    `json:"inputs"`
  Outputs   pairs.JsonPairs    `json:"outputs"`
  Actions   pairs.JsonPairs    `json:"actions"`
	JsonClaz  string    `json:"json_claz"`
  Status    string    `json:"status"`
  EventType string    `json:"event_type"`
}


func NewEventsSkews(email, cat_id string, mi map[string]string) ([]SkewsEvents, error) {

	if email == "" {
		return nil, fmt.Errorf("account_id should not be empty")
	}

	args := api.NewArgs(mi)
	args.Email = email
	cl := api.NewClient(args, EVENTSKEWS + "/" + cat_id)
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

func (s *SkewsEvents) CreateEvent(o *BillOpts, TYPE string, mi map[string]string) error {
	var exp_at, gen_at time.Time
	var action, next string
	mm := make(map[string][]string, 0)
	if s.Inputs != nil {
		/// have to change GENERATED_AT action_initiated_at next_action_due_at current_action => action
		gen_at, _ = time.Parse(time.RFC3339, s.Inputs.Match(constants.GENERATED_AT))
	} else {
		gen_at = time.Now()
	}

	softDue, err :=  time.ParseDuration(o.SoftGracePeriod)
	hardDue, err :=  time.ParseDuration(o.HardGracePeriod)
	if err != nil {
		return err
	}
	switch TYPE {
	case HARDSKEWS:
		exp_at = gen_at.Add(hardDue)
		action = HARDSKEWS
		next = "unrecoverable"
	case SOFTSKEWS:
    exp_at = gen_at.Add(hardDue)
		action = SOFTSKEWS
		next = HARDSKEWS
	case WARNING:
		mm[constants.GENERATED_AT] = []string{gen_at.Format(time.RFC3339)}
		exp_at = gen_at.Add(softDue)
    action = WARNING
		next = SOFTSKEWS
	}
  mm[constants.NEXT_DUE_AT] = []string{exp_at.Format(time.RFC3339)}
	mm[constants.CURRENT_ACTION] = []string{action}
	mm[constants.NEXT_ACTION] = []string{next}

	s.Inputs.NukeAndSet(mm)
	s.Status = "active"
  return s.Create(mi)
}


func (s *SkewsEvents) Create(mi map[string]string) error {
	args := api.NewArgs(mi)
	args.Email = s.AccountId
	cl := api.NewClient(args, EVENTSKEWS_NEW)
	_, err := cl.Post(s)
	if err != nil {
		return err
	}
	return nil
}

func (s *SkewsEvents) ActionSkewsEvents(o *BillOpts, currentBal string, mi map[string]string) error {
			log.Debugf("checks actions for skews")
	sk := make(map[string]*SkewsEvents, 0)
	// to get skews events for that particular cat_id/ asm_id
	evts, err := NewEventsSkews(o.AccountId, o.AssemblyId, mi)
	for _, v := range evts {
    if v.Status == "active" {
      sk[v.Actions.Match(constants.CURRENT_ACTION)] = &v
		}
	}
	TYPE := s.Type(o, currentBal)
	if len(sk) > 0 {
       if sk[TYPE] != nil {
			   switch true {
			   case TYPE == HARDSKEWS && sk[HARDSKEWS].IsExpired():
            // we have to do force action to Hardaction
						return nil
				 case TYPE == SOFTSKEWS && sk[SOFTSKEWS].IsExpired():
					 return sk[SOFTSKEWS].CreateEvent(o,HARDSKEWS, mi)
				 case TYPE == WARNING && sk[WARNING].IsExpired():
					 return sk[SOFTSKEWS].CreateEvent(o,SOFTSKEWS, mi)
			   }
				 return nil
			 } else {
				 return s.CreateEvent(o, TYPE, mi)
			 }
	} else {
      return s.CreateEvent(o, TYPE, mi)
	}
	return nil
}



func (s *SkewsEvents) Type(o *BillOpts,currentBal string) string {
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

func (s *SkewsEvents) IsExpired() bool {
  t1, _ := time.Parse(time.RFC3339, s.Inputs.Match("generated_at"))
  t2, _ := time.Parse(time.RFC3339, s.Inputs.Match("next_due_at"))
  duration := t2.Sub(t1)
  return t1.Add(duration).Sub(time.Now()) < time.Minute
}
