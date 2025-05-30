.PHONY: build
build:
	@go build ./cmd/root

.PHONY: test
test:
	@go test -v ./...
