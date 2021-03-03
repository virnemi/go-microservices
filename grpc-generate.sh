#!/bin/bash

protoc -I "gorpc/" \
	-I "${GOPATH}/src" \
	--go_out="gorpc" \
	gorpc/ports.proto