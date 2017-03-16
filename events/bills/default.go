package bills

import (
	"github.com/megamsys/libgo/events/alerts"
	"github.com/megamsys/libgo/utils"
	constants "github.com/megamsys/libgo/utils"
	"strconv"
	"strings"
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
		Id:       o.AccountId,
		Consumed: o.Consumed,
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

func (m scylladbManager) AuditUnpaid(o *BillOpts, mi map[string]string) error {
	sk := &EventsSkews{
		AccountId: o.AccountId,
		CatId:     o.AssemblyId,
		EventType: o.SkewsType,
	}

	if strings.Split(o.SkewsType, ".")[1] == "quota" {
		err := sk.SkewsQuotaUnpaid(o, mi)
		if err != nil {
			return err
		}
		return m.skewsWarning(o, sk)
	}

	b, err := NewBalances(o.AccountId, mi)
	if err != nil {
		return err
	}
	cb, _ := strconv.ParseFloat(b.Credit, 64)
	if cb <= 0 {
		err = sk.ActionEvents(o, b.Credit, mi)
		if err != nil {
			return err
		}
		return m.skewsWarning(o, sk)
	}

	return nil
}

func (m scylladbManager) skewsWarning(o *BillOpts, sk *EventsSkews) error {
	mm := make(map[string]string, 0)
	mm[constants.EMAIL] = sk.AccountId
	mm[constants.VERTNAME] = o.AssemblyName
	mm[constants.SOFT_ACTION] = SOFTSKEWS
	mm[constants.SOFT_GRACEPERIOD] = o.SoftGracePeriod
	mm[constants.SOFT_LIMIT] = o.SoftLimit
	mm[constants.HARD_GRACEPERIOD] = o.HardGracePeriod
	mm[constants.HARD_ACTION] = HARDSKEWS
	mm[constants.HARD_LIMIT] = o.HardLimit
	mm[constants.ACTION_TRIGGERED_AT] = sk.Inputs.Match(constants.ACTION_TRIGGERED_AT)
	mm[constants.NEXT_ACTION_DUE_AT] = sk.Inputs.Match(constants.NEXT_ACTION_DUE_AT)
	mm[constants.ACTION] = sk.Inputs.Match(constants.ACTION)
	mm[constants.NEXT_ACTION] = sk.Inputs.Match(constants.NEXT_ACTION)
	notifier := alerts.NewMailer(alerts.Mailer, alerts.Mailer)
	return notifier.Notify(alerts.SKEWS_WARNING, alerts.EventData{M: mm})
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
