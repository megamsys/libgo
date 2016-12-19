package alerts

import (
	log "github.com/Sirupsen/logrus"
  "github.com/megamsys/libgo/api"
	constants "github.com/megamsys/libgo/utils"
	"github.com/pborman/uuid"
	"time"
)

const (
	EVENTSVM_NEW     = "/eventsvm/content"
	EVENTVM_JSONCLAZ = "Megam::EventsVm"
)

type EventsVm struct {
	Id         string   `json:"di" cql:"id"`
	EventType  string   `json:"event_type" cql:"event_type"`
	AccountId  string   `json:"account_id" cql:"account_id"`
	AssemblyId string   `json:"assembly_id" cql:"assembly_id"`
	Data       []string `json:"data" cql:"data"`
	CreatedAt  time.Time `json:"created_at" cql:"created_at"`
	JsonClaz   string   `json:"json_claz" cql:"json_claz"`
}

func (v *VerticeApi) NotifyVm(eva EventAction, edata EventData) error {
	if !v.satisfied(eva) {
		return nil
	}
	sdata := parseMapToOutputFormat(edata)
	v.Args.Path = EVENTSVM_NEW
	cl := api.NewClient(v.Args)
	_, err := cl.Post(sdata)
	if err != nil {
		log.Debugf(err.Error())
		return err
	}
	return nil
}

func parseMapToOutputFormat(edata EventData) EventsVm {
	return EventsVm{
		Id:         uuid.New(),
		EventType:  edata.M[constants.EVENT_TYPE],
		AccountId:  edata.M[constants.ACCOUNT_ID],
		AssemblyId: edata.M[constants.ASSEMBLY_ID],
		Data:       edata.D,
		CreatedAt:  time.Now(),
		JsonClaz:   EVENTVM_JSONCLAZ,
	}
}
