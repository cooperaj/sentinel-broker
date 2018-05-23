package redis_test

import (
	"github.com/go-redis/redis"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cooperaj/sentinel-broker/redis"
)

var _ = Describe("Client", func() {
	Describe("connecting to a client", func() {
		Context("given some connection options", func() {
			It("will return a Client instance", func() {
				options := &redis.Options{}
				client := ConnectToClient(options)

				Expect(client).To(BeAssignableToTypeOf(&redis.Client{}))
			})
		})
	})

	Describe("connecting to a Redis instance", func() {
		Context("given some connection options", func() {
			It("will return a Client instance", func() {
				config := &Config{}

				client := ConnectToRedis("127.0.0.1", "2637", *config)

				Expect(client).To(BeAssignableToTypeOf(&redis.Client{}))
			})
		})

		Context("given some connection options with a password", func() {
			It("will return a Client instance", func() {
				redisConfig := &ConfigRedis{
					Password: "test",
				}
				config := &Config{
					Redis: *redisConfig,
				}

				client := ConnectToRedis("127.0.0.1", "2637", *config)

				Expect(client).To(BeAssignableToTypeOf(&redis.Client{}))
			})
		})
	})
})
