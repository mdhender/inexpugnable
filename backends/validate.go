// inexpugnable - an esmtp server
// Copyright (c) 2021, 2023 Michael D Henderson
// Copyright (c) 2016-2019 GuerrillaMail.com.

package backends

import (
	"errors"
)

type RcptError error

var (
	NoSuchUser          = RcptError(errors.New("no such user"))
	StorageNotAvailable = RcptError(errors.New("storage not available"))
	StorageTooBusy      = RcptError(errors.New("storage too busy"))
	StorageTimeout      = RcptError(errors.New("storage timeout"))
	QuotaExceeded       = RcptError(errors.New("quota exceeded"))
	UserSuspended       = RcptError(errors.New("user suspended"))
	StorageError        = RcptError(errors.New("storage error"))
)
