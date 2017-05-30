#
# Simple Makefile
#

PROJECT = tmplfn

VERSION = $(shell grep -m1 "Version = " $(PROJECT).go | cut -d\" -f 2)

BRANCH = $(shell git branch | grep "* " | cut -d\   -f 2)

test:
	cd numbers && go test
	go test

status:
	git status

save:
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

website:
	./mk-website.bash

publish:
	./mk-website.bash
	./publish.bash

