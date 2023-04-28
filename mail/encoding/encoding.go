// inexpugnable - an esmtp server
// Copyright (c) 2021, 2023 Michael D Henderson
// Copyright (c) 2016-2019 GuerrillaMail.com.

// encoding enables using golang.org/x/net/html/charset for converting 7bit to UTF-8.
// golang.org/x/net/html/charset supports a larger range of encodings.
// when importing, place an underscore _ in front to import for side-effects

package encoding

import (
	"io"

	"github.com/mdhender/inexpugnable/mail"
	cs "golang.org/x/net/html/charset"
)

func init() {

	mail.Dec.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		return cs.NewReaderLabel(charset, input)
	}

}
