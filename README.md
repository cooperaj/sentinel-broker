Sentinel Broker
--

[![Travis](https://img.shields.io/travis/cooperaj/sentinel-broker.svg?style=flat-square)](https://travis-ci.org/cooperaj/sentinel-broker) 
[![license](https://img.shields.io/github/license/cooperaj/sentinel-broker.svg?style=flat-square)]() 
[![Docker Pulls](https://img.shields.io/docker/pulls/cooperaj/sentinel-broker.svg?style=flat-square)](https://hub.docker.com/r/cooperaj/sentinel-broker/)

### Whats this?
*sentinel-broker* is a small Go application that runs as a webservice. In effect it acts as a service discovery layer for your Sentinel based Redis system. You could probably script this all out with etcd and such but I didn't want to run all that for this one thing.

Once *sentinel-broker* has done its job (i.e. setup your sentinel system) it exits. This works quite well in a 'run-once' situation as, if you have met the assumptions, your Redis system should be self healing thereafter.

*sentinel-broker* exposes 2 endpoints
 * _/sentinel_ **POST** to this to register as a Sentinel
 * _/redis_ **POST** to this to register as a Redis instance

You can also query these endpoints with GET requests to find out whats currently registered.

### Assumptions
 * You have a functioning Sentinel based Redis system (3 Sentinels, 2 Redis instances).
 * You can script or otherwise have your services call a web endpoint as a part of their startup procedure.
 * Your Redis servers check the active Sentinels for a master to slave to before starting in a standalone state.
 * Your Sentinel servers check for active Sentinels and their configured master before starting in a standalone state.

### General operation

 1. Startup *sentinel-broker*.
 2. Startup the Sentinel Redis system components.
   * Redis
     1. Before start, check for a working Sentinel
     1. Query it for master status
     1. Attach as a slave if it exists
     1. **If not**, send **POST** request to *sentinel-broker* and start in standalone/master mode
   * Sentinel
     1. Before start, check for a working Sentinel
     1. Query it for master status
     1. Configure self to monitor that master
     1. **If not**, send **POST** request to *sentinel-broker* and start in standalone mode
 3. Once *sentinel-broker* receives **POST** requests from 3 Sentinels and 2 Redis instances it configures them as a Sentinel controlled Redis HA system.
   * The first Redis instance to have called the endpoint gets to be master
   * The other gets slaved to it
   * The Sentinels all get attached to the new master
 4. *sentinel-broker* exits

 Assuming you meet the assumptions and rules workflow above the system will be self healing should a component restart.

### Building
```shell
$ # Have a working $GOPATH
$ git clone git@github.com:cooperaj/sentinel-broker.git \
    $GOPATH/github.com/cooperaj/sentinel-broker
$ cd $GOPATH/github.com/cooperaj/sentinel-broker
$ go get
$ go get github.com/mitchellh/gox
$ gox -os="linux darwin" -arch="amd64 arm64" -osarch="\!darwin/arm64" -ldflags="-s -w"
```

### Docker
```shell
$ # Have docker installed and working
$ docker build -t cooperaj/sentinel-broker .
```
