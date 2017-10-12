package main

import (
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Sentinel class
type Sentinel struct {
	IP string `json:"ip"`
}

// Sentinels Our list of active Sentinel instances
var Sentinels = []Sentinel{}

// Redis class
type Redis struct {
	IP string `json:"ip"`
}

// Redii Our list of active Redis instances
var Redii = []Redis{}

func main() {
	gin.DisableConsoleColor()

	r := gin.Default()

	r.POST("/sentinel", func(c *gin.Context) {
		addSentinel(c.ClientIP())

		go ShouldSetupSentinels(Sentinels, Redii)

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.GET("/sentinel", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"sentinels": Sentinels,
		})
	})

	r.POST("/redis", func(c *gin.Context) {
		addRedis(c.ClientIP())

		go ShouldSetupSentinels(Sentinels, Redii)

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.GET("/redis", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"redii": Redii,
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

func addSentinel(ip string) {
	s := Sentinel{
		IP: ip,
	}
	Sentinels = append(Sentinels, s)
}

func addRedis(ip string) {
	r := Redis{
		IP: ip,
	}
	Redii = append(Redii, r)
}
