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
.PHONY: fmt lint test testacc build install generate playground-binary playground-init playground-plan playground-apply playground-clean release-alpha release-beta release

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
	go test -v -cover -timeout=120s -parallel=10 ./internal/...

testacc:
	@if [ -z "$$POSTHOG_API_KEY" ]; then \
		read -p "POSTHOG_API_KEY not set. Enter API key: " POSTHOG_API_KEY; \
		export POSTHOG_API_KEY; \
	fi; \
	if [ -z "$$POSTHOG_PROJECT_ID" ]; then \
		read -p "POSTHOG_PROJECT_ID not set. Enter Project ID: " POSTHOG_PROJECT_ID; \
		export POSTHOG_PROJECT_ID; \
	fi; \
	if [ -z "$$POSTHOG_HOST" ]; then \
		read -p "POSTHOG_HOST not set. Enter Posthog host: " POSTHOG_HOST; \
		export POSTHOG_HOST; \
	fi; \
	echo "Running acceptance tests..."; \
	TF_ACC=1 POSTHOG_API_KEY=$$POSTHOG_API_KEY POSTHOG_PROJECT_ID=$$POSTHOG_PROJECT_ID POSTHOG_HOST=$$POSTHOG_HOST go test -v -timeout 30m ./testacc/...

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

# Release targets - create signed tags and push to trigger GoReleaser
# Usage: make release-alpha VERSION=0.1.0 NUM=1  -> v0.1.0-alpha.1
#        make release VERSION=0.1.0              -> v0.1.0
VERSION ?= 0.0.1
NUM ?= 1

release-alpha:
	@echo "Creating alpha release v$(VERSION)-alpha.$(NUM)..."
	git tag -s "v$(VERSION)-alpha.$(NUM)" -m "v$(VERSION)-alpha.$(NUM)"
	git push origin "v$(VERSION)-alpha.$(NUM)"
	@echo "Release v$(VERSION)-alpha.$(NUM) pushed. Check GitHub Actions for release status."

release-beta:
	@echo "Creating beta release v$(VERSION)-beta.$(NUM)..."
	git tag -s "v$(VERSION)-beta.$(NUM)" -m "v$(VERSION)-beta.$(NUM)"
	git push origin "v$(VERSION)-beta.$(NUM)"
	@echo "Release v$(VERSION)-beta.$(NUM) pushed. Check GitHub Actions for release status."

release:
	@echo "Creating stable release v$(VERSION)..."
	git tag -s "v$(VERSION)" -m "v$(VERSION)"
	git push origin "v$(VERSION)"
	@echo "Release v$(VERSION) pushed. Check GitHub Actions for release status."
