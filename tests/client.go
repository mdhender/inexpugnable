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

package test

import (
	"bufio"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/mdhender/inexpugnable"
	"net"
	"time"
)

func Connect(serverConfig inexpugnable.ServerConfig, deadline time.Duration) (net.Conn, *bufio.Reader, error) {
	var bufin *bufio.Reader
	var conn net.Conn
	var err error
	if serverConfig.TLS.AlwaysOn {
		// start tls automatically
		conn, err = tls.Dial("tcp", serverConfig.ListenInterface, &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         "127.0.0.1",
		})
	} else {
		conn, err = net.Dial("tcp", serverConfig.ListenInterface)
	}

	if err != nil {
		// handle error
		//t.Error("Cannot dial server", config.Servers[0].ListenInterface)
		return conn, bufin, errors.New("Cannot dial server: " + serverConfig.ListenInterface + "," + err.Error())
	}
	bufin = bufio.NewReader(conn)

	// should be ample time to complete the test
	if err = conn.SetDeadline(time.Now().Add(time.Second * deadline)); err != nil {
		return conn, bufin, err
	}
	// read greeting, ignore it
	_, err = bufin.ReadString('\n')
	return conn, bufin, err
}

func Command(conn net.Conn, bufin *bufio.Reader, command string) (reply string, err error) {
	_, err = fmt.Fprintln(conn, command+"\r")
	if err == nil {
		return bufin.ReadString('\n')
	}
	return "", err
}
