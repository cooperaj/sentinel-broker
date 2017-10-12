package main

import (
	"log"
	"os"
	"time"

	redis "github.com/go-redis/redis"
	try "github.com/matryer/try"
)

var (
	inProgress bool
)

// ShouldSetupSentinels Attempts to figure out if the sentinels should be setup
func ShouldSetupSentinels(s []Sentinel, r []Redis) {
	if len(s) > 2 { // all the sentinels in place
		if len(r) > 1 && !inProgress { // master and slave in place
			inProgress = true

			err := try.Do(func(attempt int) (bool, error) {
				var err error
				master, err := isWorkingInstance(r[0].IP)
				if master != "pong" {
					log.Println("Unable to connect to master: " + r[0].IP)
				}

				slave, err := isWorkingInstance(r[1].IP)
				if slave != "pong" {
					log.Println("Unable to connect to slave: " + r[1].IP)
				}

				if err != nil {
					time.Sleep(2 * time.Second)
				}

				return true, err // infinite retry
			})

			inProgress = false

			if err == nil {
				os.Exit(0)
			}
		}
	}
}

func isWorkingInstance(ip string) (string, error) {
	client := redis.NewClient(&redis.Options{
		Addr: ip + ":6379",
	})

	pong, err := client.Ping().Result()

	return pong, err
}
