// inexpugnable - an esmtp server
// Copyright (c) 2021, 2023 Michael D Henderson
// Copyright (c) 2016-2019 GuerrillaMail.com.

package inexpugnable

import "time"

var (
	Version   string
	Commit    string
	BuildTime string

	StartTime      time.Time
	ConfigLoadTime time.Time
)

func init() {
	// If version, commit, or build time are not set, make that clear.
	const unknown = "unknown"
	if Version == "" {
		Version = unknown
	}
	if Commit == "" {
		Commit = unknown
	}
	if BuildTime == "" {
		BuildTime = unknown
	}

	StartTime = time.Now()
}
