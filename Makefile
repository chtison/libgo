.PHONY: run

PWD := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
DIR := /root/go/src/github.com/chtison/libgo

run:
	docker run -it --rm -h docker -v $(PWD):$(DIR) -w $(DIR) chtison/workspace
