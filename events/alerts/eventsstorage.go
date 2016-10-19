package alerts

import (
	log "github.com/Sirupsen/logrus"
	ldb "github.com/megamsys/libgo/db"
	constants "github.com/megamsys/libgo/utils"
	"time"
)

const EVENTSTORAGE = "events_for_storages"

type EventsStorage struct {
	EventType string   `json:"event_type" cql:"event_type"`
	AccountId string   `json:"account_id" cql:"account_id"`
	Data      []string `json:"data" cql:"data"`
	CreatedAt string   `json:"created_at" cql:"created_at"`
}

func (s *Scylla) NotifyStorage(eva EventAction, edata EventData) error {

	if !s.satisfied(eva) {
		return nil
	}
	s_data := parseMapToOutputStorage(edata)
	ops := ldb.Options{
		TableName:   EVENTSTORAGE,
		Pks:         []string{constants.EVENT_TYPE, constants.CREATED_AT},
		Ccms:        []string{constants.ACCOUNT_ID},
		Hosts:       s.Scylla_host,
		Keyspace:    s.Scylla_keyspace,
		Username:    s.Scylla_username,
		Password:    s.Scylla_password,
		PksClauses:  map[string]interface{}{constants.EVENT_TYPE: edata.M[constants.EVENT_TYPE], constants.CREATED_AT: s_data.CreatedAt},
		CcmsClauses: map[string]interface{}{constants.ACCOUNT_ID: edata.M[constants.ACCOUNT_ID]},
	}
	if err := ldb.Storedb(ops, s_data); err != nil {
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
		CreatedAt: time.Now().String(),
	}
}
