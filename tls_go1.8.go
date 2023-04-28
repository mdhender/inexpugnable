//go:build go1.8
// +build go1.8

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

package inexpugnable

import "crypto/tls"

// add ciphers introduced since Go 1.8
func init() {
	TLSCiphers["TLS_RSA_WITH_AES_128_CBC_SHA256"] = tls.TLS_RSA_WITH_AES_128_CBC_SHA256
	TLSCiphers["TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256"] = tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256
	TLSCiphers["TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256"] = tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256
	TLSCiphers["TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256"] = tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256
	TLSCiphers["TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305"] = tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305
	TLSCiphers["TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305"] = tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305

	TLSCurves["X25519"] = tls.X25519
}
