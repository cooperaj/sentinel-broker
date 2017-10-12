Sentinel Broker
--

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
