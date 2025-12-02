PLUGIN_NAME ?= posthog
PLUGIN_VERSION ?= 0.0.0
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
BIN_DIR ?= $(CURDIR)/bin
BIN_PATH ?= $(BIN_DIR)/terraform-provider-$(PLUGIN_NAME)
PLAYGROUND_DIR ?= $(CURDIR)/playground
PLAYGROUND_TFRC ?= $(PLAYGROUND_DIR)/terraformrc
PLUGIN_FILENAME ?= terraform-provider-$(PLUGIN_NAME)_v$(PLUGIN_VERSION)

default: fmt lint install generate
.PHONY: fmt lint test testacc build install generate playground-binary playground-init playground-plan playground-apply playground-clean

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

playground-binary:
	mkdir -p $(BIN_DIR)
	go build -ldflags "-X main.version=$(PLUGIN_VERSION)" -o $(BIN_PATH) .

playground-tfrc: playground-binary
	@printf 'provider_installation {\n  dev_overrides {\n    "posthog/posthog" = "%s"\n  }\n\n  direct {\n    exclude = ["posthog/posthog"]\n  }\n}\n' "$(BIN_DIR)" > $(PLAYGROUND_TFRC)

playground-init: playground-tfrc
	cd $(PLAYGROUND_DIR) && TF_CLI_CONFIG_FILE=$(PLAYGROUND_TFRC) terraform init

playground-plan: playground-tfrc
	cd $(PLAYGROUND_DIR) && TF_CLI_CONFIG_FILE=$(PLAYGROUND_TFRC) terraform plan

playground-apply: playground-tfrc
	cd $(PLAYGROUND_DIR) && TF_CLI_CONFIG_FILE=$(PLAYGROUND_TFRC) terraform apply

playground-clean:
	rm -f $(PLAYGROUND_TFRC)
	rm -rf $(PLAYGROUND_DIR)/.terraform $(PLAYGROUND_DIR)/.terraform.lock.hcl
