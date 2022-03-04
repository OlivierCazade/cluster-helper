.PHONY: build
build:
	go build -o cluster-helper cmd/cluster-helper.go

.PHONY: lint
lint:
	golangci-lint run ./...
