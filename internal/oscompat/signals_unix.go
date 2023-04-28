//go:build amd64 && linux
// +build amd64,linux

/*******************************************************************************
inexpugnable - an esmtp server

Copyright (c) 2021 Michael D Henderson
Copyright (c) 2016-2019 GuerrillaMail.com

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
