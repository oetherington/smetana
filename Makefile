test:
	go test -v -coverprofile=coverage.cov

lint:
	go vet ./...

fmt:
	go fmt ./...

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

# Create a new release with `VERSION=v0.1.0 make release`
release:
	if [ -z ${VERSION} ]; then exit 1; else echo "Using version ${VERSION}"; fi
	git tag ${VERSION}
	git push origin ${VERSION}
	GOPROXY=proxy.golang.org go list -m github.com/oetherington/smetana@${VERSION}
