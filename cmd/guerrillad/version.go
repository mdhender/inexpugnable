// inexpugnable - an esmtp server
// Copyright (c) 2021, 2023 Michael D Henderson
// Copyright (c) 2016-2019 GuerrillaMail.com.

package main

import (
	"github.com/spf13/cobra"

	"github.com/mdhender/inexpugnable"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version info",
	Long:  `Every software has a version. This is Guerrilla's`,
	Run: func(cmd *cobra.Command, args []string) {
		logVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func logVersion() {
	mainlog.Infof("guerrillad %s", inexpugnable.Version)
	mainlog.Debugf("Build Time: %s", inexpugnable.BuildTime)
	mainlog.Debugf("Commit:     %s", inexpugnable.Commit)
}
