package webservice

import (
	"log"
	"net"
	"net/http"

	logging "github.com/cooperaj/sentinel-broker/logging"
	"github.com/cooperaj/sentinel-broker/redis"
	"github.com/gin-gonic/gin"
)

// Run Creates a webservice that listens for redis and sentinel registrations
func Run(redis *redis.Cluster) {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.GET("/config", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"config": redis.Config,
		})
	})

	r.POST("/sentinel", func(c *gin.Context) {
		go redis.AddSentinel(c.ClientIP())
		logging.Logf("Sentinel registered at %s", c.ClientIP())

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.GET("/sentinel", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"sentinels": redis.Sentinels,
		})
	})

	r.POST("/redis", func(c *gin.Context) {
		go redis.AddRedis(c.ClientIP())
		logging.Logf("Redis server registered at %s", c.ClientIP())

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.GET("/redis", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"redii": redis.Redii,
		})
	})

	server := &http.Server{Handler: r}
	l, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	err = server.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
