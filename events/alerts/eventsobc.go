package alerts

import (
	log "github.com/Sirupsen/logrus"
	ldb "github.com/megamsys/libgo/db"
	constants "github.com/megamsys/libgo/utils"
	"time"
)

const EVENTSOBCBUCKET = "events_for_obc"

type EventsObc struct {
	EventType  string   `json:"event_type" cql:"event_type"`
	AccountId  string   `json:"account_id" cql:"account_id"`
	HostIp     string   `json:"host_ip" cql:"host_ip"`
	Data       []string `json:"data" cql:"data"`
	CreatedAt  string   `json:"created_at" cql:"created_at"`
}

func (s *Scylla) NotifyOBC(eva EventAction, edata EventData) error {
	if !s.satisfied(eva) {
		return nil
	}
	s_data := parseMapToOutputObc(edata)
	ops := ldb.Options{
		TableName:   EVENTSOBCBUCKET,
		Pks:         []string{constants.EVENT_TYPE, constants.CREATED_AT},
		Ccms:        []string{constants.HOST_IP, constants.ACCOUNT_ID},
		Hosts:       s.scylla_host,
		Keyspace:    s.scylla_keyspace,
		Username:    s.scylla_username,
		Password:    s.scylla_password,
		PksClauses:  map[string]interface{}{constants.EVENT_TYPE: edata.M[constants.EVENT_TYPE], constants.CREATED_AT: s_data.CreatedAt},
		CcmsClauses: map[string]interface{}{constants.HOST_IP: edata.M[constants.HOST_IP], constants.ACCOUNT_ID: edata.M[constants.ACCOUNT_ID]},
	}
	if err := ldb.Storedb(ops, s_data); err != nil {
		log.Debugf(err.Error())
		return err
	}
	return nil
}

func parseMapToOutputObc(edata EventData) EventsObc {
	return EventsObc{
		EventType:  edata.M[constants.EVENT_TYPE],
		AccountId:  edata.M[constants.ACCOUNT_ID],
		HostIp:     edata.M[constants.HOST_IP],
		Data:       edata.D,
		CreatedAt:  time.Now().String(),
	}
}
