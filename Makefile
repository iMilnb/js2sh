
VERSION := $(shell cat VERSION)
GITLOG := $(shell git log --format=%h -1)

all:	build

build:
	go build -v ./...

test:
	go test -v ./...

tag:
	git tag -a "${VERSION}-${GITLOG}" -m "${GITMSG}"

clean:
	go clean
