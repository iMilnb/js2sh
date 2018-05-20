all:	build

build:
	go build -v ./...

test:
	go test -v ./...

clean:
	go clean
