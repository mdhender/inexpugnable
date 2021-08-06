package backends

import (
	"fmt"

	"github.com/mdhender/inexpugnable/mail"
	"github.com/mdhender/inexpugnable/response"
)

// ----------------------------------------------------------------------------------
// Processor Name: redis
// ----------------------------------------------------------------------------------
// Description   : Saves the e.Data (email data) and e.DeliveryHeader together in redis
//               : using the hash generated by the "hash" processor and stored in
//               : e.Hashes
// ----------------------------------------------------------------------------------
// Config Options: redis_expire_seconds int - how many seconds to expiry
//               : redis_interface string - <host>:<port> eg, 127.0.0.1:6379
// --------------:-------------------------------------------------------------------
// Input         : e.Data
//               : e.DeliveryHeader generated by Header() processor
//               :
// ----------------------------------------------------------------------------------
// Output        : Sets e.QueuedId with the first item fromHashes[0]
// ----------------------------------------------------------------------------------
func init() {

	processors["redis"] = func() Decorator {
		return Redis()
	}
}

type RedisProcessorConfig struct {
	RedisExpireSeconds int    `json:"redis_expire_seconds"`
	RedisInterface     string `json:"redis_interface"`
}

type RedisProcessor struct {
	isConnected bool
	conn        RedisConn
}

func (r *RedisProcessor) redisConnection(redisInterface string) (err error) {
	if r.isConnected == false {
		r.conn, err = RedisDialer("tcp", redisInterface)
		if err != nil {
			// handle error
			return err
		}
		r.isConnected = true
	}
	return nil
}

// The redis decorator stores the email data in redis

func Redis() Decorator {

	var config *RedisProcessorConfig
	redisClient := &RedisProcessor{}
	// read the config into RedisProcessorConfig
	Svc.AddInitializer(InitializeWith(func(backendConfig BackendConfig) error {
		configType := BaseConfig(&RedisProcessorConfig{})
		bcfg, err := Svc.ExtractConfig(backendConfig, configType)
		if err != nil {
			return err
		}
		config = bcfg.(*RedisProcessorConfig)
		if redisErr := redisClient.redisConnection(config.RedisInterface); redisErr != nil {
			err := fmt.Errorf("redis cannot connect, check your settings: %s", redisErr)
			return err
		}
		return nil
	}))
	// When shutting down
	Svc.AddShutdowner(ShutdownWith(func() error {
		if redisClient.isConnected {
			return redisClient.conn.Close()
		}
		return nil
	}))

	var redisErr error

	return func(p Processor) Processor {
		return ProcessWith(func(e *mail.Envelope, task SelectTask) (Result, error) {

			if task == TaskSaveMail {
				hash := ""
				if len(e.Hashes) > 0 {
					e.QueuedId = e.Hashes[0]
					hash = e.Hashes[0]
					var stringer fmt.Stringer
					// a compressor was set
					if c, ok := e.Values["zlib-compressor"]; ok {
						stringer = c.(*DataCompressor)
					} else {
						stringer = e
					}
					redisErr = redisClient.redisConnection(config.RedisInterface)
					if redisErr != nil {
						Log().WithError(redisErr).Warn("Error while connecting to redis")
						result := NewResult(response.Canned.FailBackendTransaction)
						return result, redisErr
					}
					_, doErr := redisClient.conn.Do("SETEX", hash, config.RedisExpireSeconds, stringer)
					if doErr != nil {
						Log().WithError(doErr).Warn("Error while SETEX to redis")
						result := NewResult(response.Canned.FailBackendTransaction)
						return result, doErr
					}
					e.Values["redis"] = "redis" // the next processor will know to look in redis for the message data
				} else {
					Log().Error("Redis needs a Hasher() process before it")
					result := NewResult(response.Canned.FailBackendTransaction)
					return result, StorageError
				}

				return p.Process(e, task)
			} else {
				// nothing to do for this task
				return p.Process(e, task)
			}

		})
	}
}