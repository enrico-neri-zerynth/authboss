/*
Package authboss is a modular authentication system for the web. It tries to
remove as much boilerplate and "hard things" as possible so that each time you
start a new web project in Go, you can plug it in, configure and be off to the
races without having to think about the hard questions like how to store
Remember Me tokens, or passwords.
*/
package authboss // import "gopkg.in/authboss.v0"

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

var (
	logger  io.Writer = ioutil.Discard
	emailer mailer
)

// Init authboss and it's loaded modules with a configuration.
func Init(config *Config) error {
	if config.Storer == nil {
		return errors.New("Configuration must provide a storer.")
	}

	logger = config.LogWriter

	switch config.Mailer {
	case MailerSMTP:
		// dance
	default:
		emailer = newLogMailer(logger)
	}

	for name, mod := range modules {
		fmt.Fprintf(logger, "[%-10s] Initializing\n", name)
		if err := mod.Initialize(config); err != nil {
			return fmt.Errorf("[%s] Error Initializing: %v", name, err)
		}
	}

	return nil
}
