SPEC_URL ?= https://us.posthog.com/api/schema/
GEN_DIR ?= internal/posthog/swagger
PKG_NAME ?= posthogapi
GIT_USER_ID ?= posthog
GIT_REPO_ID ?= terraform-provider
MODULE_ROOT ?= github.com/posthog/terraform-provider

default: fmt lint install generate
.PHONY: fmt lint test testacc build install generate

build:
	go build -v ./...

install: build
	go install -v ./...

lint:
	golangci-lint run

generate:
	cd tools; go generate ./...

fmt:
	gofmt -s -w -e .

test:
	go test -v -cover -timeout=120s -parallel=10 ./...

testacc:
	TF_ACC=1 go test -v -cover -timeout 120m ./...

.PHONY: generate-client
generate-client:
	@docker run --rm -v "$$PWD":/local openapitools/openapi-generator-cli:v7.17.0 generate \
		--input-spec "$(SPEC_URL)" \
		--generator-name go \
		--skip-validate-spec \
		--package-name "$(PKG_NAME)" \
		--git-user-id "$(GIT_USER_ID)" \
		--git-repo-id "$(GIT_REPO_ID)" \
		--additional-properties=isGoSubmodule=false,modulePath="$(MODULE_ROOT)/$(GEN_DIR)" \
		--output "/local/$(GEN_DIR)"
	@rm -f "$(GEN_DIR)/go.mod" "$(GEN_DIR)/go.sum"
