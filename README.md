# httptool

Tool to make http GET requests and print MD5 hash of the content


## Build

Use makefile to build httptool

```bash
# to make binary
make build

# to run tests
make test

# to make binary and run tests
make

# cleanup
make clean
```

## Usage

```bash
# to run
go run httptool.go test_url1 test_url2 ...

# to run in parallel
go run parallel httptool.go test_url1 test_url2 ...
```