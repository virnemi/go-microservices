# Go Microservices

This is a response for a Go challenge.

There are two Go projects within this repo:

1. client: the microservice that has a RESTful API to serve PortDomains. It communicates with the PortDomainService via gRPC.
1. server: The PortDomainService that persists and serves Ports through gRPC.

## Build Prerequisites

1. Protobuffer v3

## Runtime Prerequisites

1. Docker and Docker compose

## Linux environment

**Install Protobuffer v3**
```bash
apt install -y protobuf-compiler
protoc --version # to ensure it´s version 3
```

**Install Go Protobuf plugin** 
```bash
export GO111MODULE=off  # Disable module mode

go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get -u github.com/golang/protobuf/proto
go install github.com/golang/protobuf/proto
go get -u "golang.org/x/net/http2"
go install "golang.org/x/net/http2"
go get -u "golang.org/x/net/http2/hpack"
go install "golang.org/x/net/http2/hpack"
go get -u "golang.org/x/net/trace"
go install "golang.org/x/net/trace"
go get -u "google.golang.org/genproto/googleapis/rpc/status"
go install "google.golang.org/genproto/googleapis/rpc/status"
```

## Docker runtime environment

To build and run the images just run the following commands:
```bash
docker-compose build
docker-compose up
```

# Using and testing the solution

## Populate the services with the ports data:
```
POST http://localhost:8080/ports
Content-Type: multipart/form-data
multipart: 
	name: file
	vaue: [SELECT THE ports.json FILE]
```

## Get a slice (page) of ports:
```
GET	http://localhost:8080/ports/list?start=15&size=20
```

## Get a specific port:
```
GET http://localhost:8080/port?id=USSTL
```
