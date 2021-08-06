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

package guerrilla

import (
	"bufio"
	"errors"
	"io"
)

var (
	LineLimitExceeded   = errors.New("maximum line length exceeded")
	MessageSizeExceeded = errors.New("maximum message size exceeded")
)

// we need to adjust the limit, so we embed io.LimitedReader
type adjustableLimitedReader struct {
	R *io.LimitedReader
}

// bolt this on so we can adjust the limit
func (alr *adjustableLimitedReader) setLimit(n int64) {
	alr.R.N = n
}

// Returns a specific error when a limit is reached, that can be differentiated
// from an EOF error from the standard io.Reader.
func (alr *adjustableLimitedReader) Read(p []byte) (n int, err error) {
	n, err = alr.R.Read(p)
	if err == io.EOF && alr.R.N <= 0 {
		// return our custom error since io.Reader returns EOF
		err = LineLimitExceeded
	}
	return
}

// allocate a new adjustableLimitedReader
func newAdjustableLimitedReader(r io.Reader, n int64) *adjustableLimitedReader {
	lr := &io.LimitedReader{R: r, N: n}
	return &adjustableLimitedReader{lr}
}

// This is a bufio.Reader what will use our adjustable limit reader
// We 'extend' buffio to have the limited reader feature
type smtpBufferedReader struct {
	*bufio.Reader
	alr *adjustableLimitedReader
}

// Delegate to the adjustable limited reader
func (sbr *smtpBufferedReader) setLimit(n int64) {
	sbr.alr.setLimit(n)
}

// Set a new reader & use it to reset the underlying reader
func (sbr *smtpBufferedReader) Reset(r io.Reader) {
	sbr.alr = newAdjustableLimitedReader(r, CommandLineMaxLength)
	sbr.Reader.Reset(sbr.alr)
}

// Allocate a new SMTPBufferedReader
func newSMTPBufferedReader(rd io.Reader) *smtpBufferedReader {
	alr := newAdjustableLimitedReader(rd, CommandLineMaxLength)
	s := &smtpBufferedReader{bufio.NewReader(alr), alr}
	return s
}
