PWD := $(shell pwd)

depinit: 
	go get -u github.com/rancher/trash

dep: depupdate
	PATH=$(PATH):$(GOPATH)/bin trash -u

depupdate: 
	PATH=$(PATH):$(GOPATH)/bin trash
