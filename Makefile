GOBIN?=${GOPATH}/bin

all: lint install

lint-pre:
	@test -z $(gofmt -l .)
	@go mod verify

lint: lint-pre
	@golangci-lint run

lint-verbose: lint-pre
	@golangci-lint run -v --timeout=5m

install: go.sum
	GO111MODULE=on go install -v ./cmd/foggd

clean:
	rm -f ${GOBIN}/{foggd}

tests:
	@go test -v -coverprofile .testCoverage.txt ./...

run-fogg-service:
	@foggd
