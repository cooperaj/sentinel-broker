package main

import (
	"encoding/json"
	"fmt"

	"os"

	logging "github.com/cooperaj/sentinel-broker/logging"
	cluster "github.com/cooperaj/sentinel-broker/redis"
	ws "github.com/cooperaj/sentinel-broker/webservice"
)

func loadConfiguration(file string) cluster.Config {
	var config cluster.Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	authPass := os.Getenv("REDIS_PASSWORD")
	if authPass != "" {
		config.Redis.Password = authPass
	}

	return config
}

func main() {
	logging.Create("sentinel-broker: ")

	configuration := loadConfiguration("sentinel-config.json")
	redisCluster := cluster.NewCluster(configuration)

	if working, err := redisCluster.IsFunctional(); working {
		logging.Logf("%s", "Sentinel cluster operational, exiting...")
		os.Exit(0)
	} else {
		logging.Logf("%s, continuing...", err)
	}

	ws.Run(redisCluster)
}
