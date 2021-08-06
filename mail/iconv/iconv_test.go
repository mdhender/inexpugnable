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

package iconv

import (
	"github.com/mdhender/inexpugnable/mail"
	"strings"
	"testing"
)

// This will use the iconv encoder
func TestIconvMimeHeaderDecode(t *testing.T) {
	str := mail.MimeHeaderDecode("=?ISO-2022-JP?B?GyRCIVo9dztSOWJAOCVBJWMbKEI=?=")
	if i := strings.Index(str, "【女子高生チャ"); i != 0 {
		t.Error("expecting 【女子高生チャ, got:", str)
	}
	str = mail.MimeHeaderDecode("=?ISO-8859-1?Q?Andr=E9?= Pirard <PIRARD@vm1.ulg.ac.be>")
	if strings.Index(str, "André Pirard") != 0 {
		t.Error("expecting André Pirard, got:", str)
	}
}
