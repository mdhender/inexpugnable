// +build go1.13

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

import "crypto/tls"

// TLS 1.3 was introduced in go 1.12 as an option and enabled for production in go 1.13
// release notes: https://golang.org/doc/go1.12#tls_1_3
func init() {
	TLSProtocols["tls1.3"] = tls.VersionTLS13

	TLSCiphers["TLS_AES_128_GCM_SHA256"] = tls.TLS_AES_128_GCM_SHA256
	TLSCiphers["TLS_AES_256_GCM_SHA384"] = tls.TLS_AES_256_GCM_SHA384
	TLSCiphers["TLS_CHACHA20_POLY1305_SHA256"] = tls.TLS_CHACHA20_POLY1305_SHA256
}
