package alerts

import (
	log "github.com/Sirupsen/logrus"
	mailgun "github.com/mailgun/mailgun-go"
	constants "github.com/megamsys/libgo/utils"
	"strings"
)

const (
	LAUNCHED EventAction = iota
	DESTROYED
	STATUS
	DEDUCT
	ONBOARD
	RESET
	INVITE
	BALANCE
	INVOICE
	BILLEDHISTORY
	TRANSACTION
	DESCRIPTION
	SNAPSHOTTING
	SNAPSHOTTED
	RUNNING
	FAILURE
	INSUFFICIENT_FUND
	QUOTA_UNPAID
	SKEWS_ACTIONS
	SKEWS_WARNING
)

var Mailer map[string]string

type Notifier interface {
	Notify(eva EventAction, edata EventData) error
	satisfied(eva EventAction) bool
}

// Extra information about an event.
type EventData struct {
	M map[string]string
	D []string
}

type EventAction int

func (v *EventAction) String() string {
	switch *v {
	case LAUNCHED:
		return "launched"
	case DESTROYED:
		return "destroyed"
	case STATUS:
		return "status"
	case DEDUCT:
		return "deduct"
	case ONBOARD:
		return "onboard"
	case RESET:
		return "reset"
	case INVITE:
		return "invite"
	case BALANCE:
		return "balance"
	case INSUFFICIENT_FUND:
		return "insufficientfunds"
	case QUOTA_UNPAID:
		return "quotaoverdue"
	case DESCRIPTION:
		return "description"
	case SNAPSHOTTING:
		return "snapshotting"
	case SNAPSHOTTED:
		return "snapshotted"
	case RUNNING:
		return "running"
	case FAILURE:
		return "failure"
	case SKEWS_WARNING:
		return constants.SKEWS_WARNING
	default:
		return "arrgh"
	}
}

type mailgunner struct {
	api_key string
	domain  string
	sender  string
	nilavu  string
	logo    string
	home    string
	dir     string
}

func NewMailgun(m map[string]string, n map[string]string) Notifier {
	mg := &mailgunner{
		api_key: m[constants.API_KEY],
		sender:  m[constants.SENDER],
		domain:  m[constants.DOMAIN],
		nilavu:  m[constants.NILAVU],
		logo:    m[constants.LOGO],
		home:    n[constants.HOME],
		dir:     n[constants.DIR],
	}
	mg.makeGlobal()
	return mg
}

func (m *mailgunner) makeGlobal() {
	mm := make(map[string]string, 0)
	mm[constants.API_KEY] = m.api_key
	mm[constants.SENDER] = m.sender
	mm[constants.DOMAIN] = m.domain
	mm[constants.NILAVU] = m.nilavu
	mm[constants.LOGO] = m.logo
	mm[constants.HOME] = m.home
	mm[constants.DIR] = m.dir
	Mailer = mm
}

func (m *mailgunner) satisfied(eva EventAction) bool {
	if eva == STATUS {
		return false
	}
	return true
}

/*{
		"email":     "nkishore@megam.io",
		"logo":      "vertice.png",
		"nilavu":    "console.megam.io",
		"appname": "vertice.megambox.com"
		"type": "torpedo"
		"token": "9090909090",
		"days":      "20",
		"cost":      "$12",
}*/

func (m *mailgunner) Notify(eva EventAction, edata EventData) error {
	if !m.satisfied(eva) {
		return nil
	}
	edata.M[constants.NILAVU] = m.nilavu
	edata.M[constants.LOGO] = m.logo

	bdy, err := body(eva.String(), edata.M, m.dir)
	if err != nil {
		return err
	}
	m.Send(bdy, "", subject(eva), edata.M[constants.EMAIL])
	return nil
}

func (m *mailgunner) Send(msg string, sender string, subject string, to string) error {
	if len(strings.TrimSpace(sender)) <= 0 {
		sender = m.sender
	}
	mg := mailgun.NewMailgun(m.domain, m.api_key, "")
	g := mailgun.NewMessage(
		sender,
		subject,
		"You are in !",
		to,
	)
	g.SetHtml(msg)
	g.SetTracking(false)
	//g.SetTrackingClicks(false)
	//g.SetTrackingOpens(false)
	_, id, err := mg.Send(g)
	if err != nil {
		return err
	}
	log.Infof("Mailgun sent %s", id)
	return nil
}

func subject(eva EventAction) string {
	var sub string
	switch eva {
	case ONBOARD:
		sub = "Ahoy. Welcome aboard!"
	case RESET:
		sub = "You have fat finger.!"
	case INVITE:
		sub = "Lets party!"
	case BALANCE:
		sub = "Piggy bank!"
	case INSUFFICIENT_FUND:
		sub = "Insufficient funds!"
	case QUOTA_UNPAID:
		sub = "Payment pending!"
	case LAUNCHED:
		sub = "Up!"
	case RUNNING:
		sub = "Ahoy! Your application is running "
	case DESTROYED:
		sub = "Nuked"
	case SNAPSHOTTING:
		sub = "Snapshot creating!"
	case SNAPSHOTTED:
		sub = "Ahoy! Snapshot created"
	case FAILURE:
		sub = "Your application failure"
	case SKEWS_WARNING:
		return "Payment Reminder !"
	default:
		break
	}
	return sub
}
