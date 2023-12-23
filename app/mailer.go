package app

import (
	"github.com/mailgun/mailgun-go/v4"
)

type Mailer struct {
	Mg *mailgun.MailgunImpl
}

func (mailer *Mailer) Init(domain string, pkey string) {
	mailer.Mg = mailgun.NewMailgun(domain, pkey)
}
