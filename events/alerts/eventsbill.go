package alerts

import (
	log "github.com/Sirupsen/logrus"
	ldb "github.com/megamsys/libgo/db"
	constants "github.com/megamsys/libgo/utils"
	"time"
)

const EVENTBILL = "events_for_billings"

type EventsBill struct {
	EventType  string   `json:"event_type" cql:"event_type"`
	AccountId  string   `json:"account_id" cql:"account_id"`
	AssemblyId string   `json:"assembly_id" cql:"assembly_id"`
	Data       []string `json:"data" cql:"data"`
	CreatedAt  string   `json:"created_at" cql:"created_at"`
}

func (s *Scylla) NotifyBill(eva EventAction, edata EventData) error {
	if !s.satisfied(eva) {
		return nil
	}
	s_data := parseMapToOutputBill(edata)
	ops := ldb.Options{
		TableName:   EVENTBILL,
		Pks:         []string{constants.EVENT_TYPE, constants.CREATED_AT},
		Ccms:        []string{constants.ASSEMBLY_ID, constants.ACCOUNT_ID},
		Hosts:       s.scylla_host,
		Keyspace:    s.scylla_keyspace,
		Username:    s.scylla_username,
		Password:    s.scylla_password,
		PksClauses:  map[string]interface{}{constants.EVENT_TYPE: edata.M[constants.EVENT_TYPE], constants.CREATED_AT: s_data.CreatedAt},
		CcmsClauses: map[string]interface{}{constants.ASSEMBLY_ID: edata.M[constants.ASSEMBLY_ID], constants.ACCOUNT_ID: edata.M[constants.ACCOUNT_ID]},
	}
	if err := ldb.Storedb(ops, s_data); err != nil {
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
		CreatedAt:  time.Now().String(),
	}
}
