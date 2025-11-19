package provider

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	frameworkresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	internaldata "github.com/posthog/terraform-provider/internal/data"
	"github.com/posthog/terraform-provider/internal/examples"
	"github.com/posthog/terraform-provider/internal/posthog"
	posthogapi "github.com/posthog/terraform-provider/internal/posthog/swagger"
	posthogresource "github.com/posthog/terraform-provider/internal/resource"
)

var (
	_ provider.Provider                       = &PostHogProvider{}
	_ provider.ProviderWithFunctions          = &PostHogProvider{}
	_ provider.ProviderWithEphemeralResources = &PostHogProvider{}
)

// PostHogProviderModel describes the provider data model.
type PostHogProviderModel struct {
	Host      types.String `tfsdk:"host"`
	APIKey    types.String `tfsdk:"api_key"`
	ProjectID types.String `tfsdk:"project_id"`
}

// ProjectID can be an int

// PostHogProvider defines the provider implementation.
type PostHogProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &PostHogProvider{
			version: version,
		}
	}
}

func (p *PostHogProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "posthog"
	resp.Version = p.version
}

func (p *PostHogProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Base URL for the PostHog API. Defaults to `https://us.posthog.com`. Can be set via `POSTHOG_HOST`",
			},
			"api_key": schema.StringAttribute{
				Optional:            true,
				Sensitive:           true,
				MarkdownDescription: "PostHog personal API key. Can be set via `POSTHOG_API_KEY`.",
			},
			"project_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Default project ID (environment) to target. Can be set via `POSTHOG_PROJECT_ID`.",
			},
		},
	}
}

func (p *PostHogProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data PostHogProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }
	host := "https://us.posthog.com"
	if !data.Host.IsNull() {
		host = data.Host.ValueString()
	} else if envHost := os.Getenv("POSTHOG_HOST"); envHost != "" {
		host = envHost
	}

	apiKey := os.Getenv("POSTHOG_API_KEY")
	if !data.APIKey.IsNull() {
		apiKey = data.APIKey.ValueString()
	}
	if apiKey == "" {
		resp.Diagnostics.AddError("Missing API key", "The provider requires a PostHog personal API key. Set it in configuration or POSTHOG_API_KEY.")
		return
	}

	projectID := os.Getenv("POSTHOG_PROJECT_ID")
	if !data.ProjectID.IsNull() {
		projectID = data.ProjectID.ValueString()
	}

	tflog.Debug(ctx, "configured PostHog provider", map[string]any{
		"host": host,
	})

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}

	config := posthogapi.NewConfiguration()
	config.HTTPClient = httpClient
	// Remove scheme from host - swagger config expects just hostname
	config.Host = strings.TrimPrefix(strings.TrimPrefix(host, "https://"), "http://")
	config.Scheme = "https"
	config.UserAgent = "posthog/terraform-provider v0.0.0"
	// Only set Authorization header - Content-Type and Accept are set per-request by swagger client
	config.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	providerData := internaldata.ProviderData{
		Client:        posthog.NewDefaultClient(logger, host, apiKey, projectID),
		SwaggerClient: posthogapi.NewAPIClient(config),
		ProjectID:     projectID,
	}
	resp.DataSourceData = providerData
	resp.ResourceData = providerData
}

func (p *PostHogProvider) Resources(ctx context.Context) []func() frameworkresource.Resource {
	return []func() frameworkresource.Resource{
		examples.NewExampleResource,
		posthogresource.NewDashboard,
		posthogresource.NewInsight,
	}
}

func (p *PostHogProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{
		examples.NewExampleEphemeralResource,
	}
}

func (p *PostHogProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		examples.NewExampleDataSource,
	}
}

func (p *PostHogProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		examples.NewExampleFunction,
	}
}
