package mail_test

import (
	"github.com/lfhy/mail"
	"golang.org/x/net/proxy"
)

func Example() {
	mailer := mail.New(
		mail.SetHost("smtp.example.com"),
		mail.SetPort(587),
		mail.SetUser("user@example.com"),
		mail.SetPassword("password"),
		mail.SetFrom("user"),
		mail.SetSkipTLS(true),
		mail.SetDialer(proxy.Direct),
	)
	mailer.SendMail("title", "test mail", "user2@example.com")
}
