package alerts

import (
	constants "github.com/megamsys/libgo/utils"
	"strings"
)

type Scylla struct {
	scylla_host     []string
	scylla_keyspace string
}

func NewScylla(m map[string]string) Notifier {
	return &Scylla{
		scylla_host:  strings.Split(m[constants.SCYLLAHOST], ","),
		scylla_keyspace: m[constants.SCYLLAKEYSPACE],
	}
}

func (s *Scylla) satisfied(eva EventAction) bool {
	if eva == STATUS {
		return true
	}
	return false
}

func (s *Scylla) Notify(eva EventAction, edata EventData) error {
	value := edata.M[constants.EVENT_TYPE]
	et := strings.Split(value, ".")
	if et[0] == "compute" {
		return s.NotifyVm(eva, edata)
	} else if et[0] == "bill" {
		return s.NotifyBill(eva, edata)
	} else if et[0] == "storage" {
		return s.NotifyStorage(eva, edata)
	} else if et[0] == "obc" {
		return s.NotifyOBC(eva, edata)
	} else {
		return s.NotifyContainer(eva, edata)
	}

	return nil
}
