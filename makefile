APP=flusher

run:
	@go run cmd/main.go

test:
	ginkgo -r -v

build:
	@go build -o $(APP) cmd/main.go

clean:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean
	@rm $(APP) 2> /dev/null | true

vet:
	@go vet ./...

fmt:
	@go fmt ./...
