BINARY_NAME=httptool
SOURCE_NAME=httptool.go

all: build test

build:
	go build -o ${BINARY_NAME} ${SOURCE_NAME}

clean:
	go clean

test:
	go test ./... -v -race

