package bills

import (
	"github.com/megamsys/libgo/utils"
	"strconv"
	"fmt"
)

func init() {
	Register(utils.SCYLLAMGR, scylladbManager{})
}

type scylladbManager struct{}

func (m scylladbManager) IsEnabled() bool {
	return true
}

func (m scylladbManager) Onboard(o *BillOpts, mi map[string]string) error {
	return nil
}

func (m scylladbManager) Deduct(o *BillOpts, mi map[string]string) error {
	b, err := NewBalances(o.AccountId, mi)
	if err != nil {
		return err
	}

	if err = b.Deduct(&BalanceOpts{
		Id:        o.AccountId,
		Consumed:  o.Consumed,
	}, mi); err != nil {
		return err
	}
	return nil
}

func (m scylladbManager) Transaction(o *BillOpts, mi map[string]string) error {
	//bt, err := carton.NewBillTransaction(o.AccountsId)
	bt, err := NewBilledHistories(o)
	if err != nil {
		return err
	}
	//if err = bt.Transact(&carton.BillTransactionOpts{
	if err = bt.BilledHistories(mi); err != nil {
		return err
	}
	return nil
}

func (m scylladbManager) RecurringSkews(o *BillOpts, mi map[string]string) error {
 	sk :=  &SkewsEvents{
		AccountId: o.AccountId,
		CatId: o.AssemblyId,
		EventType: "VM",
	}

	b, err := NewBalances(o.AccountId, mi)
	if err != nil {
		return err
	}
  cb, _ := strconv.ParseFloat(b.Credit, 64)
  if  cb <= 0 {
    return sk.ActionSkewsEvents(o, b.Credit,mi)
	}

  return nil
}

func (m scylladbManager) Invoice(o *BillOpts) error {
	return nil
}

func (m scylladbManager) Nuke(o *BillOpts) error {
	return nil
}

func (m scylladbManager) Suspend(o *BillOpts) error {
	return nil
}

func (m scylladbManager) Notify(o *BillOpts) error {
	return nil
}
