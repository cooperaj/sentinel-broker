package redis

import (
	"errors"
	"fmt"
	"net"

	"github.com/cooperaj/sentinel-broker/logging"
	"github.com/go-redis/redis"
)

// Cluster information that
type Cluster struct {
	Config Config

	Sentinels []Sentinel
	Redii     []Redis
}

// Sentinel class
type Sentinel struct {
	IP string `json:"ip"`
}

// Redis class
type Redis struct {
	IP string `json:"ip"`
}

// NewCluster Creates a new Redis cluster object configured with config
func NewCluster(config Config) *Cluster {
	c := new(Cluster)
	c.Config = config

	redis.SetLogger(logging.NullLogger)
	return c
}

// AddSentinel Adds a sentinel registration via IP address
func (c *Cluster) AddSentinel(ip string) {
	s := Sentinel{
		IP: ip,
	}
	c.Sentinels = append(c.Sentinels, s)

	SetupCluster(c)
}

// AddRedis Adds a redis instance via IP address
func (c *Cluster) AddRedis(ip string) {
	r := Redis{
		IP: ip,
	}
	c.Redii = append(c.Redii, r)

	SetupCluster(c)
}

// IsFunctional Test the cluster to see if it has a defined and functioning master
func (c *Cluster) IsFunctional() (bool, error) {
	options := &redis.FailoverOptions{
		MasterName: c.Config.Master,
		SentinelAddrs: []string{
			net.JoinHostPort(
				c.Config.Sentinel.Hostname,
				fmt.Sprintf("%d", c.Config.Sentinel.Port),
			),
		},
	}

	if c.Config.Redis.Password != "" {
		options.Password = c.Config.Redis.Password
	}

	failoverClient := redis.NewFailoverClient(options)

	info, err := failoverClient.Ping().Result()

	if info != "PONG" {
		err = errors.New("Sentinel/Redis cluster not functional")
	}

	return err == nil, err
}
