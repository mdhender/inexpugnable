//go:build windows
// +build windows

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
