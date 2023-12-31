PHONY: init
init:
	cp .env.sample .env
	direnv allow .

PHONY: lint
lint:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint run


PHONY: build.rest
build.rest:
	CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o bin/rest ./cmd/rest

