// inexpugnable - an esmtp server
// Copyright (c) 2021, 2023 Michael D Henderson
// Copyright (c) 2016-2019 GuerrillaMail.com.

//go:build !darwin && !dragonfly && !freebsd && !linux && !netbsd && !openbsd
// +build !darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd

package inexpugnable

import "errors"

// getFileLimit checks how many files we can open
// Don't know how to get that info (yet?), so returns false information & error
func getFileLimit() (uint64, error) {
	return 1000000, errors.New("syscall.RLIMIT_NOFILE not supported on your OS/platform")
}
