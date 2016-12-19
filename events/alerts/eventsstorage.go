package alerts

import (
	log "github.com/Sirupsen/logrus"
	"github.com/megamsys/libgo/api"
	constants "github.com/megamsys/libgo/utils"
	"time"
)

const EVENTSTORAGE_NEW = "/eventsstorage/content"

type EventsStorage struct {
	EventType string   `json:"event_type" cql:"event_type"`
	AccountId string   `json:"account_id" cql:"account_id"`
	Data      []string `json:"data" cql:"data"`
	CreatedAt time.Time   `json:"created_at" cql:"created_at"`
}

func (v *VerticeApi) NotifyStorage(eva EventAction, edata EventData) error {

	if !v.satisfied(eva) {
		return nil
	}
	sdata := parseMapToOutputFormat(edata)
	v.Args.Path = EVENTSTORAGE_NEW
	cl := api.NewClient(v.Args)
	_, err := cl.Post(sdata)
	if err != nil {
		log.Debugf(err.Error())
		return err
	}
	return nil
}

func parseMapToOutputStorage(edata EventData) EventsStorage {
	return EventsStorage{
		EventType: edata.M[constants.EVENT_TYPE],
		AccountId: edata.M[constants.ACCOUNT_ID],
		Data:      edata.D,
		CreatedAt: time.Now(),
	}
}
