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

package response

import (
	"testing"
)

func TestGetBasicStatusCode(t *testing.T) {
	// Known status code
	a := getBasicStatusCode(EnhancedStatusCode{2, OtherOrUndefinedProtocolStatus})
	if a != 250 {
		t.Errorf("getBasicStatusCode. Int \"%d\" not expected.", a)
	}

	// Unknown status code
	b := getBasicStatusCode(EnhancedStatusCode{2, OtherStatus})
	if b != 200 {
		t.Errorf("getBasicStatusCode. Int \"%d\" not expected.", b)
	}
}

// TestString for the String function
func TestCustomString(t *testing.T) {
	// Basic testing
	resp := &Response{
		EnhancedCode: OtherStatus,
		BasicCode:    200,
		Class:        ClassSuccess,
		Comment:      "Test",
	}

	if resp.String() != "200 2.0.0 Test" {
		t.Errorf("CustomString failed. String \"%s\" not expected.", resp)
	}

	// Default String
	resp2 := &Response{
		EnhancedCode: OtherStatus,
		Class:        ClassSuccess,
	}
	if resp2.String() != "200 2.0.0 OK" {
		t.Errorf("String failed. String \"%s\" not expected.", resp2)
	}
}

func TestBuildEnhancedResponseFromDefaultStatus(t *testing.T) {
	//a := buildEnhancedResponseFromDefaultStatus(ClassPermanentFailure, InvalidCommand)
	a := EnhancedStatusCode{ClassPermanentFailure, InvalidCommand}.String()
	if a != "5.5.1" {
		t.Errorf("buildEnhancedResponseFromDefaultStatus failed. String \"%s\" not expected.", a)
	}
}
