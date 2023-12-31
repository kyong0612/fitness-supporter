PHONY: init
init:
	cp .env.sample .env
	direnv allow .

PHONY: lint
lint:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint run

PHONY: server.run
server.run:
	go build -mod=readonly -v -o bin/server ./cmd/server && ./bin/server

PHONY: server.build
server.build:
	CGO_ENABLED=0 go build -mod=readonly -v -o bin/server ./cmd/server

PHONY: compose.up
compose.up:
	docker compose up -d

