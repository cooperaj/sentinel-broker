# Build sentinel-broker binary
FROM golang:1.9 AS build-env

ADD . /src

RUN cd /src \
    && go get -d -v \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o app .

# Run image
FROM alpine

COPY --from=build-env /src/app /app/sentinel-broker

WORKDIR /app
EXPOSE 8080

ENTRYPOINT ./sentinel-broker
