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

package mocks

import (
	"io"
	"net"
	"time"
)

// Mocks a net.Conn - server and client sides.
// See server_test.go for usage examples
// Taken from https://github.com/jordwest/mock-conn
// This great answer http://stackoverflow.com/questions/1976950/simulate-a-tcp-connection-in-go

// Addr is a fake network interface which implements the net.Addr interface
type Addr struct {
	NetworkString string
	AddrString    string
}

func (a Addr) Network() string {
	return a.NetworkString
}

func (a Addr) String() string {
	return a.AddrString
}

// End is one 'end' of a simulated connection.
type End struct {
	Reader *io.PipeReader
	Writer *io.PipeWriter
}

func (e End) Close() error {
	if err := e.Writer.Close(); err != nil {
		return err
	}
	if err := e.Reader.Close(); err != nil {
		return err
	}
	return nil
}

func (e End) Read(data []byte) (n int, err error)  { return e.Reader.Read(data) }
func (e End) Write(data []byte) (n int, err error) { return e.Writer.Write(data) }

func (e End) LocalAddr() net.Addr {
	return Addr{
		NetworkString: "tcp",
		AddrString:    "127.0.0.1",
	}
}

func (e End) RemoteAddr() net.Addr {
	return Addr{
		NetworkString: "tcp",
		AddrString:    "127.0.0.1",
	}
}

func (e End) SetDeadline(t time.Time) error      { return nil }
func (e End) SetReadDeadline(t time.Time) error  { return nil }
func (e End) SetWriteDeadline(t time.Time) error { return nil }

// MockConn facilitates testing by providing two connected ReadWriteClosers
// each of which can be used in place of a net.Conn
type Conn struct {
	Server *End
	Client *End
}

func NewConn() *Conn {
	// A connection consists of two pipes:
	// Client      |      Server
	//   writes   ===>  reads
	//    reads  <===   writes

	serverRead, clientWrite := io.Pipe()
	clientRead, serverWrite := io.Pipe()

	return &Conn{
		Server: &End{
			Reader: serverRead,
			Writer: serverWrite,
		},
		Client: &End{
			Reader: clientRead,
			Writer: clientWrite,
		},
	}
}

func (c *Conn) Close() error {
	if err := c.Server.Close(); err != nil {
		return err
	}
	if err := c.Client.Close(); err != nil {
		return err
	}
	return nil
}
