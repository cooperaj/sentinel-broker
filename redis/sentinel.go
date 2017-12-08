package redis

import (
	"fmt"
	"net"

	"github.com/go-redis/redis"
)

// AttachSentinelToMaster Configures a sentinel to monitor a master
func AttachSentinelToMaster(sentinel *redis.Client, masterIP string, cluster *Cluster) error {
	working, err := IsWorkingInstance(sentinel)

	if err != nil || !working {
		panic(fmt.Sprintf("Sentinel (%s) not reachable", sentinel.String()))
	}

	err = SentinelMonitorCommand(
		sentinel,
		cluster.Config.Master,
		masterIP,
		fmt.Sprintf("%d", cluster.Config.Redis.Port),
	).Err()

	if err != nil {
		return err
	}

	return err
}

// ConfigureSentinel Configures the sentinel with settings
func ConfigureSentinel(client *redis.Client, cluster *Cluster) error {
	for command, rawValue := range cluster.Config.Sentinel.Config {
		value, err := cluster.Config.ConvertToString(rawValue)
		if err != nil {
			return err
		}

		err = SentinelSetCommand(client, cluster.Config.Master, command, value).Err()
		if err != nil {
			return err
		}
	}

	if cluster.Config.Redis.Password != "" {
		err := SentinelSetCommand(
			client,
			cluster.Config.Master,
			"auth-pass",
			cluster.Config.Redis.Password,
		).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

// ConnectToSentinel Connects to a Redis server
func ConnectToSentinel(ip string, port string, config Config) *redis.Client {
	options := &redis.Options{
		Addr: net.JoinHostPort(ip, port),
	}

	return ConnectToClient(options)
}

// SentinelMonitorCommand Configures sentinal intance to monitor a new master
func SentinelMonitorCommand(
	sentinel *redis.Client,
	masterName string,
	masterIP string,
	redisPort string,
) *redis.StringCmd {
	cmd := redis.NewStringCmd(
		"sentinel",
		"monitor",
		masterName,
		masterIP,
		redisPort,
		"2",
	)
	sentinel.Process(cmd)
	return cmd
}

// SentinelSetCommand Configures sentinel instance with key value
func SentinelSetCommand(
	client *redis.Client,
	masterName string,
	key string,
	value string,
) *redis.StringCmd {
	cmd := redis.NewStringCmd("sentinel", "set", masterName, key, value)
	client.Process(cmd)
	return cmd
}
