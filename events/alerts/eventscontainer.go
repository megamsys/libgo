package alerts

import (
	log "github.com/Sirupsen/logrus"
	ldb "github.com/megamsys/libgo/db"
	constants "github.com/megamsys/libgo/utils"
	"github.com/pborman/uuid"
	"time"
	"fmt"
)

const (
	EVENTCONTAINER          = "events_for_containers"
	EVENTCONTAINER_JSONCLAZ = "Megam::EventsContainer"
)

type EventsContainer struct {
	Id         string    `josn:"id" cql:"id"`
	EventType  string    `json:"event_type" cql:"event_type"`
	AccountId  string    `json:"account_id" cql:"account_id"`
	AssemblyId string    `json:"assembly_id" cql:"assembly_id"`
	Data       []string  `json:"data" cql:"data"`
	CreatedAt  time.Time `json:"created_at" cql:"created_at"`
	JsonClaz   string    `json:"json_claz" cql:"json_claz"`
}

func (s *Scylla) NotifyContainer(eva EventAction, edata EventData) error {
	if !s.satisfied(eva) {
		return nil
	}
	s_data := parseMapToOutputContainer(edata)
	ops := ldb.Options{
		TableName:   EVENTCONTAINER,
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
	fmt.Println("--------events container----------")
	fmt.Println(s_data)
	fmt.Println("------------------")
	return nil
}

func parseMapToOutputContainer(edata EventData) EventsContainer {
	return EventsContainer{
		Id:         uuid.New(),
		EventType:  edata.M[constants.EVENT_TYPE],
		AccountId:  edata.M[constants.ACCOUNT_ID],
		AssemblyId: edata.M[constants.ASSEMBLY_ID],
		Data:       edata.D,
		CreatedAt:  time.Now(),
		JsonClaz:   EVENTCONTAINER_JSONCLAZ,
	}
}
