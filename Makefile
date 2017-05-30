#
# Simple Makefile
#

PROJECT = tmplfn

VERSION = $(shell grep -m1 "Version = " $(PROJECT).go | cut -d\" -f 2)

BRANCH = $(shell git branch | grep "* " | cut -d\   -f 2)

test:
	cd numbers && go test
	go test

website:
	./mk-website.bash

publish:
	./mk-website.bash
	./publish.bash

