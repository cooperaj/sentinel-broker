package redis

import (
	"fmt"
	"os"
)

var (
	inProgress bool
)

// SetupCluster Works with the registered entities and starts the cluster when ready
func SetupCluster(cluster *Cluster) {
	if len(cluster.Sentinels) > 2 { // all the sentinels in place
		if len(cluster.Redii) > 1 && !inProgress { // master and slave in place
			inProgress = true

			masterIP := cluster.Redii[0].IP
			slaveIP := cluster.Redii[1].IP

			// Waits till connection is up
			verifyMaster(masterIP, cluster)

			setupSlave(slaveIP, masterIP, cluster)

			setupSentinels(masterIP, cluster)

			// Success, we have a working sentinel stack
			os.Exit(0)
		}
	}
}

func setupSentinels(masterIP string, cluster *Cluster) {
	for _, sentinel := range cluster.Sentinels {
		client := ConnectToClient(
			sentinel.IP,
			fmt.Sprintf("%d", cluster.Config.Sentinel.Port),
			cluster.Config,
		)

		defer func() {
			client.Close()
		}()

		err := AttachSentinelToMaster(client, masterIP, cluster)
		if err != nil {
			panic(err.Error())
		}
	}
}

func setupSlave(slaveIP string, masterIP string, cluster *Cluster) {
	client := ConnectToClient(
		slaveIP,
		fmt.Sprintf("%d", cluster.Config.Redis.Port),
		cluster.Config,
	)

	defer func() {
		client.Close()
	}()

	working, err := IsWorkingInstance(client)
	if err != nil || !working {
		panic(fmt.Sprintf("Slave (%s) not reachable", client.String()))
	}

	err = client.SlaveOf(masterIP, fmt.Sprintf("%v", cluster.Config.Redis.Port)).Err()
	if err != nil {
		panic(err.Error())
	}
}

func verifyMaster(masterIP string, cluster *Cluster) {
	client := ConnectToClient(
		masterIP,
		fmt.Sprintf("%d", cluster.Config.Redis.Port),
		cluster.Config,
	)

	defer func() {
		client.Close()
	}()

	working, err := IsWorkingInstance(client)
	if err != nil || !working {
		panic(fmt.Sprintf("Master (%s) not reachable", client.String()))
	}
}
