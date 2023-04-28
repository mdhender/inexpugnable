// inexpugnable - an esmtp server
// Copyright (c) 2021, 2023 Michael D Henderson
// Copyright (c) 2016-2019 GuerrillaMail.com.

//go:build windows
// +build windows

package oscompat

func DefaultPidFile() string {
	return "C:\\Go-Guerilla.pid"
}
