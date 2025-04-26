package mail

import (
	"crypto/tls"

	"golang.org/x/net/proxy"
)

type Option func(*MailConfig)

func SetUser(user string) Option {
	return func(mc *MailConfig) {
		mc.User = user
	}
}

func SetPassword(password string) Option {
	return func(mc *MailConfig) {
		mc.Password = password
	}
}

func SetHost(host string) Option {
	return func(mc *MailConfig) {
		mc.Host = host
	}
}

func SetPort(port int) Option {
	return func(mc *MailConfig) {
		mc.Port = port
	}
}

func SetFrom(from string) Option {
	return func(mc *MailConfig) {
		mc.From = from
	}
}

func SetSubject(subject string) Option {
	return func(mc *MailConfig) {
		mc.Subject = subject
	}
}

func SetContent(content string) Option {
	return func(mc *MailConfig) {
		mc.Content = content
	}
}

func SetDialer(dialer proxy.Dialer) Option {
	return func(mc *MailConfig) {
		mc.Dialer = dialer
	}
}

func SetSkipTLS(skip bool) Option {
	return func(mc *MailConfig) {
		if mc.TLSConfig == nil { // if not set, create a new one
			mc.TLSConfig = &tls.Config{}
		}
		mc.TLSConfig.InsecureSkipVerify = skip
	}
}

func SetTLSConfig(config *tls.Config) Option {
	return func(mc *MailConfig) {
		mc.TLSConfig = config
	}
}
