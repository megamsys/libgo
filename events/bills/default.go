package bills

//import (
	//"github.com/megamsys/vertice/carton"
//)

func init() {
	Register("scylladb", scylladbManager{})
}

type scylladbManager struct{}

func (m scylladbManager) IsEnabled() bool {
	return true
}

func (m scylladbManager) Onboard(o *BillOpts) error {
	return nil
}

func (m scylladbManager) Deduct(o *BillOpts) error {
	//b, err := carton.NewBalances(o.AccountsId)
	b, err := NewBalances(o.AccountsId)
	if err != nil {
		return err
	}
	//if err = b.Deduct(&carton.BalanceOpts{
	if err = b.Deduct(&BalanceOpts{
		Id:        o.AccountsId,
		Consumed:  o.Consumed,
		Timestamp: o.Timestamp,
	}); err != nil {
		return err
	}
	return nil
}

func (m scylladbManager) Transaction(o *BillOpts) error {
	//bt, err := carton.NewBillTransaction(o.AccountsId)
	bt, err := NewBillTransaction(o.AccountsId)
	if err != nil {
		return err
	}
	//if err = bt.Transact(&carton.BillTransactionOpts{
	if err = bt.Transact(&BillTransactionOpts{
		Id:        o.AccountsId,
		Timestamp: o.Timestamp,
	}); err != nil {
		return err
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
