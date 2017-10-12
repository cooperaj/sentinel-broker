Sentinel Broker
--

[![Build Status](https://travis-ci.org/cooperaj/sentinel-broker.svg?branch=master)](https://travis-ci.org/cooperaj/sentinel-broker) [![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)

### Whats this?
Sentinel-Broker is a small Go application that runs as a webservice. In effect it acts as a service discovery layer for your Sentinel based Redis system. You could probably script this all out with etcd and such but I didn't want to run all that for this one thing.

Once Sentinel-Broker has done its job (i.e. setup your sentinel system) it exits. This works quite well in a 'run-once' situation as, if you have met the assumptions, your Redis system should be self healing thereafter.

### Assumptions
 * You have a functioning Sentinel based Redis system (3 Sentinels, 2 Redis instances).
 * You can script or otherwise have your services call a web endpoint as a part of their startup procedure.
 * Your Redis servers check the active Sentinels for a master to slave to before starting as a master themselves.
 * Your Sentinel servers check for active Sentinels and their configured master before starting in a standalone state.

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
