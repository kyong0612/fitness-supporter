include .env

PHONY: init
init:
	cp .env.sample .env
	direnv allow .

PHONY: lint
lint:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint run

PHONY: lint.fix
lint.fix:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint run --fix

PHONY: server.run
server.run:
	@go build -mod=readonly -v -o bin/server ./cmd/server && ./bin/server

PHONY: server.build
server.build:
	@CGO_ENABLED=0 go build -mod=readonly -v -o bin/server ./cmd/server

PHONY: compose.up
compose.up:
	@docker compose up -d


PHONY: deploy.apply
deploy.apply:
	gcloud deploy apply \
  		--file=.clouddeploy/clouddeploy.yaml \
  		--region=asia-northeast1 \
  		--project=kyong0612-lab

PHONY: deploy.release
deploy.release:
	gcloud deploy releases create munual-release-$(shell date +%Y%m%d%H%M%S) \
		--source=.clouddeploy \
  		--project=kyong0612-lab \
  		--region=asia-northeast1 \
  		--delivery-pipeline=fitness-support \
		--deploy-parameters="line_token=$(LINE_CHANNEL_TOKEN)"
