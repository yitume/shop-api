ROOT:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
OUT:=$(ROOT)/bin/shop-api
PKG:=$(shell go list -m)

.PHONY: all clean build test
all: all clean build test

build:
	./tool/build.sh $(OUT) $(PKG)/pkg/version $(PKG)

clean:
	@rm -rf bin

test:
