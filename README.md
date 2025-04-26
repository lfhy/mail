# Gomail
[![Documentation](https://godoc.org/github.com/lfhy/mail)](https://godoc.org/github.com/lfhy/mail)

## Introduction

Gomail is a simple and efficient package to send emails. It is well tested and
documented.

Gomail can only send emails using an SMTP server. But the API is flexible and it
is easy to implement other methods for sending emails using a local Postfix, an
API, etc.

It is versioned using [gopkg.in](https://gopkg.in) so I promise
there will never be backward incompatible changes within each version.

It requires Go 1.2 or newer. With Go 1.5, no external dependencies are used.


## Features

Gomail supports:
- Attachments
- Embedded images
- HTML and text templates
- Automatic encoding of special characters
- SSL and TLS
- Sending multiple emails with the same SMTP connection
- Support use proxy dialer


## Documentation

https://godoc.org/github.com/lfhy/mail


## Download

    go install github.com/lfhy/mail@latest


## Examples

See the [examples in the documentation](https://godoc.org/github.com/lfhy/mail).


## FAQ

### x509: certificate signed by unknown authority

If you get this error it means the certificate used by the SMTP server is not
considered valid by the client running Gomail. As a quick workaround you can
bypass the verification of the server's certificate chain and host name by using
`SetTLSConfig`:
```go
    package main

    import (
    	"crypto/tls"

    	"github.com/lfhy/mail"
    )

    func main() {
    	d := mail.New(mail.SetTLSConfig(&tls.Config{InsecureSkipVerify: true}))
        // or use mail.SetSkipTLS(true) to skip TLS
        // d := mail.New(mail.SetSkipTLS(true))
        // Send emails using d.
    }
```
Note, however, that this is insecure and should not be used in production.


## Contribute

Contributions are more than welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for
more info.


## Change log

See [CHANGELOG.md](CHANGELOG.md).


## License

[MIT](LICENSE)


## Contact

You can ask questions on the [Gomail
thread](https://groups.google.com/d/topic/golang-nuts/jMxZHzvvEVg/discussion)
in the Go mailing-list.
