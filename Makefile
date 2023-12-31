PHONY: init
init:
	cp .env.sample .env
	direnv allow .

PHONY: run.rest
run.rest:
	go build -mod=readonly -v -o bin/rest ./cmd/rest && ./bin/rest

PHONY: build.rest
build.rest:
	CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o bin/rest ./cmd/rest

