all:	build

build:
	go build -v ./...

clean:
	go clean
