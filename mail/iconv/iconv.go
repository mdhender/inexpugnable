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

// iconv enables using GNU iconv for converting 7bit to UTF-8.
// iconv supports a larger range of encodings.
// It's a cgo package, the build system needs have Gnu library headers available.
// when importing, place an underscore _ in front to import for side-effects
package iconv

import (
	"fmt"
	"io"

	"github.com/mdhender/inexpugnable/mail"
	ico "gopkg.in/iconv.v1"
)

func init() {
	mail.Dec.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if cd, err := ico.Open("UTF-8", charset); err == nil {
			r := ico.NewReader(cd, input, 32)
			return r, nil
		}
		return nil, fmt.Errorf("unhandled charset %q", charset)
	}

}
