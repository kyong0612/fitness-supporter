SHELL=/bin/bash
include .env

.PHONY: init
init:
	cp .env.sample .env
	direnv allow .

.PHONY: generate.buf
generate.buf:
	@go run github.com/bufbuild/buf/cmd/buf generate

.PHONY: lint.buf
lint.buf:
	@go run github.com/bufbuild/buf/cmd/buf lint
	@go run github.com/bufbuild/buf/cmd/buf format -w

.PHONY: lint.fix
lint.fix:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint run --fix
	@go run golang.org/x/vuln/cmd/govulncheck ./...
	@make lint.buf

.PHONY: server.run
server.run:
	@go build -mod=readonly -v -o bin/server ./cmd/server && ./bin/server

.PHONY: server.build
server.build:
	@CGO_ENABLED=0 go build -mod=readonly -v -o bin/server ./cmd/server

.PHONY: compose.up
compose.up:
	@docker compose up -d

.PHONY: deploy.all
deploy.all: 
	make deploy.apply
	make deploy.build 
	make deploy.release

.PHONY: deploy.apply
deploy.apply:
	gcloud deploy apply \
  		--file=.clouddeploy/clouddeploy.yaml \
  		--region=asia-northeast1 \
  		--project=kyong0612-lab

.PHONY: deploy.build
deploy.build:
	@-docker image rm asia-northeast1-docker.pkg.dev/kyong0612-lab/fitness-supporter/prd
	@docker buildx build . --platform linux/amd64 --no-cache --tag asia-northeast1-docker.pkg.dev/kyong0612-lab/fitness-supporter/prd:latest
	@docker push asia-northeast1-docker.pkg.dev/kyong0612-lab/fitness-supporter/prd:latest

.PHONY: deploy.release
deploy.release:
	gcloud deploy releases create munual-release-$(shell date +%Y%m%d%H%M%S) \
		--source=.clouddeploy \
  		--project=kyong0612-lab \
  		--region=asia-northeast1 \
  		--delivery-pipeline=fitness-support \
		--deploy-parameters="\
			line_secret_token=$(LINE_CHANNEL_SECRET),\
			line_access_token=$(LINE_CHANNEL_ACCESS_TOKEN),\
			gemini_key=$(GEMINI_API_KEY)\
			project_id={GCP_PROJECT_ID}"
