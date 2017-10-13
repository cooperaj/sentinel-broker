package redis

import (
	"log"
	"time"

	redis "github.com/go-redis/redis"
	try "github.com/matryer/try"
)

// ConnectToClient Connects to a Redis server
func ConnectToClient(ip string, port string, config Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: ip + ":" + port,
	})

	return client
}

// IsWorkingInstance Checks that a registered IP is up and running. Blocking
func IsWorkingInstance(client *redis.Client) (bool, error) {
	err := try.Do(func(attempt int) (bool, error) {
		pong, err := client.Ping().Result()
		if err != nil && pong != "PONG" {
			log.Printf("Client (%s) not ready, will retry", client.Options().Addr)
		} else {
			log.Printf("Client (%s) ready", client.Options().Addr)
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
