/*******************************************************************************
inexpugnable - an esmtp server

Copyright (c) 2016 GuerrillaMail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
******************************************************************************/

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
