// inexpugnable - an esmtp server
// Copyright (c) 2021, 2023 Michael D Henderson
// Copyright (c) 2016-2019 GuerrillaMail.com.

//go:build amd64 && linux
// +build amd64,linux

package oscompat

func DefaultPidFile() string {
	return "/var/run/go-guerrilla.pid"
}
