package redis_test

import (
	"github.com/go-redis/redis"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cooperaj/sentinel-broker/redis"
)

var _ = Describe("Client", func() {
	Describe("Managing a Redis client", func() {
		Context("Given some connection options", func() {
			It("will return a Client instance", func() {
				options := &redis.Options{}
				client := ConnectToClient(options)

				Expect(client).To(BeAssignableToTypeOf(&redis.Client{}))
			})
		})
	})
})
