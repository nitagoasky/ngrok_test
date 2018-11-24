.PHONY: all deps server
export GOPATH:=$(shell pwd)

BUILDTAGS=debug
default: all

all: server

deps: 
	go get -tags '$(BUILDTAGS)' -d -v ngrok/...

server: deps
	cd src && go run ngrok/main/ngrokd/ngrokd.go