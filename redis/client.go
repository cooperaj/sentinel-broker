package redis

import (
	"net"
	"time"

	"github.com/cooperaj/sentinel-broker/logging"
	"github.com/go-redis/redis"
	"github.com/matryer/try"
)

// Client Interface to our Redis library
type Client interface {
	String() string
	Ping() *redis.StatusCmd
	Options() *redis.Options
	SlaveOf(host, port string) *redis.StatusCmd
	Process(redis.Cmder) error
	Close() error
}

// ConnectToClient Connects to a Redis server
func ConnectToClient(options *redis.Options) Client {
	return redis.NewClient(options)
}

// ConnectToRedis Connects to a Redis server
func ConnectToRedis(ip string, port string, config Config) Client {
	options := &redis.Options{
		Addr: net.JoinHostPort(ip, port),
	}

	if config.Redis.Password != "" {
		options.Password = config.Redis.Password
	}

	return ConnectToClient(options)
}

// IsWorkingInstance Checks that a registered IP is up and running. Blocking
func IsWorkingInstance(client Client) (bool, error) {
	err := try.Do(func(attempt int) (bool, error) {
		pong, err := client.Ping().Result()
		if err != nil && pong != "PONG" {
			logging.Logf("Client (%s) not ready, will retry", client.Options().Addr)
		} else {
			logging.Logf("Client (%s) ready", client.Options().Addr)
		}

		if err != nil {
			time.Sleep(2 * time.Second)
		}

		return attempt < 10, err
	})

	if err == nil {
		return true, err
	}

	return false, err
}
