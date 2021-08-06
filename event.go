/*******************************************************************************
inexpugnable - an esmtp server

Copyright (c) 2016 GuerrillaMail.com

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

package inexpugnable

import (
	evbus "github.com/asaskevich/EventBus"
)

type Event int

const (
	// when a new config was loaded
	EventConfigNewConfig Event = iota
	// when allowed_hosts changed
	EventConfigAllowedHosts
	// when pid_file changed
	EventConfigPidFile
	// when log_file changed
	EventConfigLogFile
	// when it's time to reload the main log file
	EventConfigLogReopen
	// when log level changed
	EventConfigLogLevel
	// when the backend's config changed
	EventConfigBackendConfig
	// when a new server was added
	EventConfigServerNew
	// when an existing server was removed
	EventConfigServerRemove
	// when a new server config was detected (general event)
	EventConfigServerConfig
	// when a server was enabled
	EventConfigServerStart
	// when a server was disabled
	EventConfigServerStop
	// when a server's log file changed
	EventConfigServerLogFile
	// when it's time to reload the server's log
	EventConfigServerLogReopen
	// when a server's timeout changed
	EventConfigServerTimeout
	// when a server's max clients changed
	EventConfigServerMaxClients
	// when a server's TLS config changed
	EventConfigServerTLSConfig
)

var eventList = [...]string{
	"config_change:new_config",
	"config_change:allowed_hosts",
	"config_change:pid_file",
	"config_change:log_file",
	"config_change:reopen_log_file",
	"config_change:log_level",
	"config_change:backend_config",
	"server_change:new_server",
	"server_change:remove_server",
	"server_change:update_config",
	"server_change:start_server",
	"server_change:stop_server",
	"server_change:new_log_file",
	"server_change:reopen_log_file",
	"server_change:timeout",
	"server_change:max_clients",
	"server_change:tls_config",
}

func (e Event) String() string {
	return eventList[e]
}

type EventHandler struct {
	evbus.Bus
}

func (h *EventHandler) Subscribe(topic Event, fn interface{}) error {
	if h.Bus == nil {
		h.Bus = evbus.New()
	}
	return h.Bus.Subscribe(topic.String(), fn)
}

func (h *EventHandler) Publish(topic Event, args ...interface{}) {
	h.Bus.Publish(topic.String(), args...)
}

func (h *EventHandler) Unsubscribe(topic Event, handler interface{}) error {
	return h.Bus.Unsubscribe(topic.String(), handler)
}
