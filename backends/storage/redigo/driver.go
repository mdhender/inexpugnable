// inexpugnable - an esmtp server
// Copyright (c) 2021, 2023 Michael D Henderson
// Copyright (c) 2016-2019 GuerrillaMail.com.

package redigo_driver

import "github.com/mdhender/inexpugnable/backends"
import redigo "github.com/gomodule/redigo/redis"

func init() {
	backends.RedisDialer = func(network, address string, options ...backends.RedisDialOption) (backends.RedisConn, error) {
		return redigo.Dial(network, address)
	}
}
