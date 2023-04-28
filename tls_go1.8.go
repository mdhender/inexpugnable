// inexpugnable - an esmtp server
// Copyright (c) 2021, 2023 Michael D Henderson
// Copyright (c) 2016-2019 GuerrillaMail.com.

//go:build go1.8
// +build go1.8

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
