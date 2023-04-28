// inexpugnable - an esmtp server
// Copyright (c) 2021, 2023 Michael D Henderson
// Copyright (c) 2016-2019 GuerrillaMail.com.

//go:build windows
// +build windows

package oscompat

import (
	"github.com/mdhender/inexpugnable"
	"github.com/mdhender/inexpugnable/log"
	"golang.org/x/sys/windows"
	"os"
	"os/signal"
	"time"
)

// SignalHandler captures and manages SIGHUP and friends.
// It expects signalChannel to be a buffered channel. For example:
//
//	signalChannel := make(chan os.Signal, 1)
//
// Warning: Windows does not have SIGUSR1, so there is no way to reset log files
func SignalHandler(signalChannel chan os.Signal, d inexpugnable.Daemon, mainlog log.Logger, readConfig func() (*inexpugnable.AppConfig, error)) {
	signal.Notify(signalChannel,
		windows.SIGHUP,
		windows.SIGTERM,
		windows.SIGQUIT,
		windows.SIGINT,
		windows.SIGKILL,
		os.Kill,
	)
	for sig := range signalChannel {
		if sig == windows.SIGHUP {
			if ac, err := readConfig(); err == nil {
				_ = d.ReloadConfig(*ac)
			} else {
				mainlog.WithError(err).Error("Could not reload config")
			}
		} else if sig == windows.SIGTERM || sig == windows.SIGQUIT || sig == windows.SIGINT || sig == os.Kill {
			mainlog.Infof("Shutdown signal caught")
			go func() {
				select {
				// exit if graceful shutdown not finished in 60 sec.
				case <-time.After(time.Second * 60):
					mainlog.Error("graceful shutdown timed out")
					os.Exit(1)
				}
			}()
			d.Shutdown()
			mainlog.Infof("Shutdown completed, exiting.")
			return
		} else {
			mainlog.Infof("Shutdown, unknown signal caught")
			return
		}
	}
}
