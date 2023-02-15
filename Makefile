test:
	go test -v -coverprofile=coverage.out

lint:
	go vet .

fmt:
	go fmt .
	cd example && go fmt .

check-fmt:
	if [ "$$(gofmt -s -l . | wc -l)" -gt 0 ]; then exit 1; fi

run-example:
	go run example/main.go

build-example:
	cd example && go build

run-built-example:
	./example/example

docs:
	godoc -http=127.0.0.1:6060

check: test run-example lint check-fmt
