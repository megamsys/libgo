package db

import (
	log "github.com/Sirupsen/logrus"
	"github.com/megamsys/gocassa"
	"github.com/megamsys/libgo/cmd"
)

type RelationsFunc func() gocassa.Relation

type ScyllaDB struct {
	NodeIps []string
	KS      gocassa.KeySpace
}

type ScyllaTable struct {
	T gocassa.Table
}

type ScyllaWhere struct {
	clauses map[string]string
}

type ScyllaDBOpts struct {
	KeySpaceName string
	NodeIps      []string
	Username     string
	Password     string
	Debug        bool
}

func NewScyllaDB(opts ScyllaDBOpts) (*ScyllaDB, error) {
	ks, err := connectToKeySpace(opts.KeySpaceName, opts.NodeIps, opts.Username, opts.Password)
	if err != nil {
		return nil, err
	}
	ks.DebugMode(opts.Debug)

	return &ScyllaDB{
		NodeIps: opts.NodeIps,
		KS:      ks,
	}, nil
}

// Connect to a certain keyspace directly. Same as using Connect().KeySpace(keySpaceName)
func connectToKeySpace(keySpace string, nodeIps []string, username, password string) (gocassa.KeySpace, error) {
	c, err := gocassa.Connect(nodeIps, username, password)
	if err != nil {
		return nil, err
	}
	log.Debugf(cmd.Colorfy("  > [scylla] keyspace "+keySpace, "blue", "", "bold"))
	return c.KeySpace(keySpace), nil
}

func (sy *ScyllaDB) Table(name string, pks []string, ccms []string, out interface{}) *ScyllaTable {
	log.Debugf(cmd.Colorfy("  > [scylla] table "+name, "blue", "", "bold"))
	return &ScyllaTable{T: sy.KS.Table(name, out, gocassa.Keys{
		PartitionKeys:     pks,
		ClusteringColumns: ccms,
	})}
}

func (st *ScyllaTable) Read(fn RelationsFunc, out interface{}) error {
	log.Debugf(cmd.Colorfy("  > [scylla] read", "blue", "", "bold"))
	return st.T.Where(fn()).ReadOne(&out).Run()
}

func (st *ScyllaTable) ReadWhere(where ScyllaWhere, out interface{}) error {
	log.Debugf(cmd.Colorfy("  > [scylla] readwhere", "blue", "", "bold"))
	return st.T.Where(where.toEqs()...).ReadOne(&out).Run()
}

func (st *ScyllaTable) Upsert(data interface{}) error {
	log.Debugf(cmd.Colorfy("  > [scylla] upsert", "blue", "", "bold"))
	return st.T.Set(data).Run()
}

func (wh ScyllaWhere) toEqs() []gocassa.Relation {
	r := make([]gocassa.Relation,0)
	for k, v := range wh.clauses {
		r = append(r, gocassa.Eq(k, v))
	}
	return r
}
