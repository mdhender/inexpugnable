// inexpugnable - an esmtp server
// Copyright (c) 2021, 2023 Michael D Henderson
// Copyright (c) 2016-2019 GuerrillaMail.com.

//go:build amd64 && linux
// +build amd64,linux

package oscompat

import (
	"github.com/mdhender/inexpugnable"
	"github.com/mdhender/inexpugnable/log"
	"golang.org/x/sys/unix"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// SignalHandler captures and manages SIGHUP and friends.
// It expects signalChannel to be a buffered channel. For example:
//
//	signalChannel := make(chan os.Signal, 1)
func SignalHandler(signalChannel chan os.Signal, d inexpugnable.Daemon, mainlog log.Logger, readConfig func() (*inexpugnable.AppConfig, error)) {
	signal.Notify(signalChannel,
		unix.SIGHUP,
		unix.SIGTERM,
		unix.SIGQUIT,
		unix.SIGINT,
		unix.SIGKILL,
		unix.SIGUSR1,
		os.Kill,
	)
	for sig := range signalChannel {
		if sig == unix.SIGHUP {
			if ac, err := readConfig(configPath, pidFile); err == nil {
				_ = d.ReloadConfig(*ac)
			} else {
				mainlog.WithError(err).Error("Could not reload config")
			}
		} else if sig == syscall.SIGUSR1 {
			if err := d.ReopenLogs(); err != nil {
				mainlog.WithError(err).Error("reopening logs failed")
			}
		} else if sig == unix.SIGTERM || sig == unix.SIGQUIT || sig == unix.SIGINT || sig == os.Kill {
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
