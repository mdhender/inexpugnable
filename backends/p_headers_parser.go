// inexpugnable - an esmtp server
// Copyright (c) 2021, 2023 Michael D Henderson
// Copyright (c) 2016-2019 GuerrillaMail.com.

package backends

import (
	"github.com/mdhender/inexpugnable/mail"
)

// ----------------------------------------------------------------------------------
// Processor Name: headersparser
// ----------------------------------------------------------------------------------
// Description   : Parses the header using e.ParseHeaders()
// ----------------------------------------------------------------------------------
// Config Options: none
// --------------:-------------------------------------------------------------------
// Input         : envelope
// ----------------------------------------------------------------------------------
// Output        : Headers will be populated in e.Header
// ----------------------------------------------------------------------------------
func init() {
	processors["headersparser"] = func() Decorator {
		return HeadersParser()
	}
}

func HeadersParser() Decorator {
	return func(p Processor) Processor {
		return ProcessWith(func(e *mail.Envelope, task SelectTask) (Result, error) {
			if task == TaskSaveMail {
				if err := e.ParseHeaders(); err != nil {
					Log().WithError(err).Error("parse headers error")
				}
				// next processor
				return p.Process(e, task)
			} else {
				// next processor
				return p.Process(e, task)
			}
		})
	}
}
