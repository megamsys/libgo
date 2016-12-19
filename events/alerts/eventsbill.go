package alerts

import (
	log "github.com/Sirupsen/logrus"
	"github.com/megamsys/libgo/api"
	constants "github.com/megamsys/libgo/utils"
	"time"
)

const EVENTBILL_NEW = "/eventsbilling/content"

type EventsBill struct {
	EventType  string   `json:"event_type" cql:"event_type"`
	AccountId  string   `json:"account_id" cql:"account_id"`
	AssemblyId string   `json:"assembly_id" cql:"assembly_id"`
	Data       []string `json:"data" cql:"data"`
	CreatedAt  time.Time   `json:"created_at" cql:"created_at"`
}

func (v *VerticeApi) NotifyBill(eva EventAction, edata EventData) error {
	if !v.satisfied(eva) {
		return nil
	}
	sdata := parseMapToOutputFormat(edata)
	v.Args.Path = EVENTBILL_NEW
	cl := api.NewClient(v.Args)
	_, err := cl.Post(sdata)
	if err != nil {
		log.Debugf(err.Error())
		return err
	}
	return nil
}

func parseMapToOutputBill(edata EventData) EventsBill {
	return EventsBill{
		EventType:  edata.M[constants.EVENT_TYPE],
		AccountId:  edata.M[constants.ACCOUNT_ID],
		AssemblyId: edata.M[constants.ASSEMBLY_ID],
		Data:       edata.D,
		CreatedAt:  time.Now(),
	}
}
