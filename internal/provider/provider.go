package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	frameworkresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	internaldata "github.com/posthog/terraform-provider/internal/data"
	"github.com/posthog/terraform-provider/internal/httpclient"
	posthogresource "github.com/posthog/terraform-provider/internal/resource"
)

const (
	EnvPostHogAPIKey    = "POSTHOG_API_KEY"
	EnvPostHogHost      = "POSTHOG_HOST"
	EnvPostHogProjectId = "POSTHOG_PROJECT_ID"
)

var (
	_ provider.Provider                       = &PostHogProvider{}
	_ provider.ProviderWithFunctions          = &PostHogProvider{}
	_ provider.ProviderWithEphemeralResources = &PostHogProvider{}
)

type PostHogProviderModel struct {
	Host      types.String `tfsdk:"host"`
	APIKey    types.String `tfsdk:"api_key"`
	ProjectID types.String `tfsdk:"project_id"`
}

type PostHogProvider struct {
	version string
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &PostHogProvider{
			version: version,
		}
	}
}

func (p *PostHogProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "posthog"
	resp.Version = p.version
}

func (p *PostHogProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Optional: true,
				MarkdownDescription: fmt.Sprintf(
					"Base URL for the PostHog API. Defaults to `https://us.posthog.com`. Can be set via `%s`", EnvPostHogHost,
				),
			},
			"api_key": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				MarkdownDescription: fmt.Sprintf(
					"PostHog personal API key. Can be set via `%s` environment variable.", EnvPostHogAPIKey,
				),
			},
			"project_id": schema.StringAttribute{
				Optional: true,
				MarkdownDescription: fmt.Sprintf(
					"Default project ID (environment) to target. Can be set via `%s` environment variable.", EnvPostHogProjectId,
				),
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

	host := "https://us.posthog.com"
	if !data.Host.IsNull() {
		host = data.Host.ValueString()
	} else if envHost := os.Getenv(EnvPostHogHost); envHost != "" {
		host = envHost
	}

	apiKey := os.Getenv(EnvPostHogAPIKey)
	if !data.APIKey.IsNull() {
		apiKey = data.APIKey.ValueString()
	}
	if apiKey == "" {
		resp.Diagnostics.AddError("Missing API key",
			fmt.Sprintf("The provider requires a PostHog personal API key. Set it in configuration or as env var: %s.", EnvPostHogAPIKey))
		return
	}

	projectID := os.Getenv(EnvPostHogProjectId)
	if !data.ProjectID.IsNull() {
		projectID = data.ProjectID.ValueString()
	}
	if projectID == "" {
		resp.Diagnostics.AddError("Missing project ID",
			fmt.Sprintf("The provider requires a PostHog project ID. Set it in configuration or as env var: %s.", EnvPostHogProjectId))
		return
	}

	tflog.Debug(ctx, "configured PostHog provider", map[string]any{"host": host})

	providerData := internaldata.ProviderData{
		Client: httpclient.NewDefaultClient(host, apiKey, projectID),
	}
	resp.DataSourceData = providerData
	resp.ResourceData = providerData
}

func (p *PostHogProvider) Resources(_ context.Context) []func() frameworkresource.Resource {
	return []func() frameworkresource.Resource{
		posthogresource.NewDashboard,
		posthogresource.NewInsight,
		posthogresource.NewFeatureFlag,
	}
}

func (p *PostHogProvider) EphemeralResources(_ context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{}
}

func (p *PostHogProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *PostHogProvider) Functions(_ context.Context) []func() function.Function {
	return []func() function.Function{}
}
