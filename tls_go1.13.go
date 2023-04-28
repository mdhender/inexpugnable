// inexpugnable - an esmtp server
// Copyright (c) 2021, 2023 Michael D Henderson
// Copyright (c) 2016-2019 GuerrillaMail.com.

//go:build go1.13
// +build go1.13

package inexpugnable

import "crypto/tls"

// TLS 1.3 was introduced in go 1.12 as an option and enabled for production in go 1.13
// release notes: https://golang.org/doc/go1.12#tls_1_3
func init() {
	TLSProtocols["tls1.3"] = tls.VersionTLS13

	TLSCiphers["TLS_AES_128_GCM_SHA256"] = tls.TLS_AES_128_GCM_SHA256
	TLSCiphers["TLS_AES_256_GCM_SHA384"] = tls.TLS_AES_256_GCM_SHA384
	TLSCiphers["TLS_CHACHA20_POLY1305_SHA256"] = tls.TLS_CHACHA20_POLY1305_SHA256
}
