package mail

import (
	"github.com/lfhy/mail/gomail"
	"golang.org/x/net/proxy"
)

type MailConfig struct {
	User     string // Send mail user: user@test.com
	Password string // Send mail password: password
	Host     string // Send mail host: stmp.test.com
	Port     int    // Send mail port: 465
	From     string // Send mail show from: testuser

	Subject string       // Send mail subject: test mail
	Content string       // Send mail content: test mail content
	Dialer  proxy.Dialer // Send mail dialer

	mailDialer *gomail.Dialer
}

func New() *MailConfig {
	return &MailConfig{}
}

func (m *MailConfig) SetUser(user string) *MailConfig {
	m.User = user
	return m
}

func (m *MailConfig) SetPassword(password string) *MailConfig {
	m.Password = password
	return m
}

func (m *MailConfig) SetHost(host string) *MailConfig {
	m.Host = host
	return m
}

func (m *MailConfig) SetPort(port int) *MailConfig {
	m.Port = port
	return m
}

func (m *MailConfig) SetFrom(from string) *MailConfig {
	m.From = from
	return m
}

func (m *MailConfig) SetSubject(subject string) *MailConfig {
	m.Subject = subject
	return m
}

func (m *MailConfig) SetContent(content string) *MailConfig {
	m.Content = content
	return m
}

func (m *MailConfig) SetDialer(dialer proxy.Dialer) *MailConfig {
	m.Dialer = dialer
	return m
}

func (m *MailConfig) Send(mailTo ...string) error {
	if len(mailTo) == 0 {
		return nil
	}
	msg := gomail.NewMessage(
		gomail.SetEncoding(gomail.Base64),
	)
	if m.Subject == "" {
		m.Subject = "new mail"
	}
	if m.From == "" {
		m.From = m.User
	}
	msg.SetHeader("From", msg.FormatAddress(m.User, m.From))
	msg.SetHeader("To", mailTo...)
	msg.SetHeader("Subject", m.Subject)
	msg.SetBody("text/html", m.Content)
	if m.mailDialer == nil {
		m.mailDialer = gomail.NewDialer(m.Host, m.Port, m.User, m.Password)
		if m.Dialer != nil {
			m.mailDialer.Dialer = m.Dialer
		}
	}
	return m.mailDialer.DialAndSend(msg)
}

func (m *MailConfig) SendMail(subject, content string, mailTo ...string) error {
	if len(mailTo) == 0 {
		return nil
	}
	msg := gomail.NewMessage(
		gomail.SetEncoding(gomail.Base64),
	)
	msg.SetHeader("From", msg.FormatAddress(m.User, m.From))
	msg.SetHeader("To", mailTo...)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", content)
	if m.mailDialer == nil {
		m.mailDialer = gomail.NewDialer(m.Host, m.Port, m.User, m.Password)
		if m.Dialer != nil {
			m.mailDialer.Dialer = m.Dialer
		}
	}
	return m.mailDialer.DialAndSend(msg)
}
