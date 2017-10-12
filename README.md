Sentinel Broker
--

[![Build Status](https://travis-ci.org/cooperaj/sentinel-broker.svg?branch=master)](https://travis-ci.org/cooperaj/sentinel-broker) [![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)

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
