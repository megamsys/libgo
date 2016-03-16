package alerts

import (
     ldb "github.com/megamsys/libgo/db"
     log "github.com/Sirupsen/logrus"
	"time"
	"strings"
	)

const EVENTSBUCKET = "events"

type Scylla struct {
	scylla_host         []string 
	scylla_keyspace     string   
}

type Events struct {
	EventType   string   	`json:"event_type" cql:"event_type"`
	AccountId   string		`json:"account_id" cql:"account_id"`
	AssemblyId  string		`json:"assembly_id" cql:"assembly_id"`
	Data         []string	`json:"data" cql:"data"`
	CreatedAt   string		`json:"created_at" cql:"created_at"`
}

func NewScylla(m map[string]string) Notifier {
	return &Scylla{
		scylla_host:  strings.Split(m[SCYLLAHOST], ","),
		scylla_keyspace: m[SCYLLAKEYSPACE],
		}
}

func (s *Scylla) satisfied() bool {
	return true
}

func (s *Scylla) Notify(eva EventAction, edata EventData) error {
	if !s.satisfied() {
		return nil
	}
	s_data := parseMapToOutputFormat(edata)
	ops := ldb.Options{
			TableName:   EVENTSBUCKET,
			Pks:         []string{ASSEMBLY_ID, EVENT_TYPE, CREATED_AT},
			Ccms:        []string{ACCOUNT_ID},
			Hosts:       s.scylla_host,
			Keyspace:    s.scylla_keyspace,
			PksClauses:  map[string]interface{}{ASSEMBLY_ID: edata.M[ASSEMBLY_ID], EVENT_TYPE: edata.M[EVENT_TYPE], CREATED_AT: s_data.CreatedAt},
			CcmsClauses: map[string]interface{}{ACCOUNT_ID: edata.M[ACCOUNT_ID]},
		}	
		if err := ldb.Storedb(ops, s_data); err != nil {
			log.Debugf(err.Error())
			return err
		}
	
	return nil
}

func parseMapToOutputFormat(edata EventData) Events {
   	return Events{
   		EventType:  edata.M[EVENT_TYPE],   
		AccountId:  edata.M[ACCOUNT_ID],
		AssemblyId: edata.M[ASSEMBLY_ID],	
		Data: edata.D,
		CreatedAt: time.Now().String(),
   	}
}




