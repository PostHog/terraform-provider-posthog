PLUGIN_NAME ?= posthog
PLUGIN_VERSION ?= 0.0.0
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
BIN_DIR ?= $(CURDIR)/bin
BIN_PATH ?= $(BIN_DIR)/terraform-provider-$(PLUGIN_NAME)
PLAYGROUND_DIR ?= $(CURDIR)/playground
PLAYGROUND_TFRC ?= $(PLAYGROUND_DIR)/terraformrc
PLUGIN_FILENAME ?= terraform-provider-$(PLUGIN_NAME)_v$(PLUGIN_VERSION)
# PostHog OpenAPI Client Generation
# Uses openapi-generator with tag filtering for insights and dashboards endpoints
POSTHOG_SPEC_URL ?= https://app.posthog.com/api/schema/
POSTHOG_SWAGGER_DIR ?= internal/posthog/swagger
OPENAPI_GENERATOR_VERSION ?= v7.17.0

default: fmt lint install generate
.PHONY: fmt lint test testacc build install generate playground-binary playground-init playground-plan playground-apply playground-clean swagger-generate swagger-clean

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

swagger-generate:
	@echo "Generating PostHog Go client with openapi-generator..."
	@echo "  Spec URL: $(POSTHOG_SPEC_URL)"
	@echo "  Output: $(POSTHOG_SWAGGER_DIR)"
	@echo "  Filtering tags: insights, dashboards"
	@mkdir -p $(POSTHOG_SWAGGER_DIR)
	@docker run --rm --network="host" \
		-v "$(CURDIR)":/local \
		openapitools/openapi-generator-cli:$(OPENAPI_GENERATOR_VERSION) generate \
		--input-spec "$(POSTHOG_SPEC_URL)" \
		--generator-name go \
		--package-name posthogapi \
		--git-user-id posthog \
		--git-repo-id terraform-provider \
		--output /local/$(POSTHOG_SWAGGER_DIR) \
		--global-property models="Insight:DashboardTileBasic:UserBasic:EffectiveRestrictionLevelEnum:PaginatedInsightList:Dashboard:PatchedDashboard:DashboardCollaborator:PatchedInsight:SharingConfiguration:DashboardRestrictionLevel:CreationModeEnum:EffectivePrivilegeLevelEnum:PaginatedDashboardBasicList:DashboardBasic",apis,supportingFiles \
		--additional-properties=enumClassPrefix=true,structPrefix=false \
		--openapi-normalizer "FILTER=tag:insights|dashboards"
	@rm -f "$(POSTHOG_SWAGGER_DIR)/go.mod" "$(POSTHOG_SWAGGER_DIR)/go.sum"
	@echo "✓ Generated client at $(POSTHOG_SWAGGER_DIR)/"

swagger-clean:
	@echo "Cleaning generated swagger files..."
	@rm -rf $(POSTHOG_SWAGGER_DIR)
	@echo "Swagger files cleaned"
