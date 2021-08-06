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
	"fmt"
	"net/smtp"
)

const (
	URL = "127.0.0.1:2500"
)

func lastWords(message string, err error) {
	fmt.Println(message, err.Error())
}

func sendMail(i int) {
	fmt.Printf("Sending %d mail\n", i)
	c, err := smtp.Dial(URL)
	if err != nil {
		lastWords("Dial ", err)
	}
	defer func() {
		_ = c.Close()
	}()

	from := "somebody@gmail.com"
	to := "somebody.else@gmail.com"

	if err = c.Mail(from); err != nil {
		lastWords("Mail ", err)
	}

	if err = c.Rcpt(to); err != nil {
		lastWords("Rcpt ", err)
	}

	wr, err := c.Data()
	if err != nil {
		lastWords("Data ", err)
	}
	defer func() {
		_ = wr.Close()
	}()

	msg := fmt.Sprint("Subject: something\n")
	msg += "From: " + from + "\n"
	msg += "To: " + to + "\n"
	msg += "\n\n"
	msg += "hello\n"

	_, err = fmt.Fprint(wr, msg)
	if err != nil {
		lastWords("Send ", err)
	}

	fmt.Printf("About to quit %d\n", i)
	err = c.Quit()
	if err != nil {
		lastWords("Quit ", err)
	}
	fmt.Printf("Finished sending %d mail\n", i)
}
