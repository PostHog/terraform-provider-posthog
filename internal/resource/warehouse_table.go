package resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
)

// NewWarehouseTable builds the posthog_warehouse_table resource, a
// self-managed data warehouse table read in place from the customer's bucket.
//
// PostHog reads the files to derive the table's columns when the table is
// created, but a PATCH to format or url_pattern leaves the old columns in
// place. Both therefore force replacement, which re-runs that read. The access
// key and secret are write-only, so the configured values are kept in state.
func NewWarehouseTable() resource.Resource {
	return core.NewGenericResource[WarehouseTableTFModel, httpclient.WarehouseTableRequest, httpclient.WarehouseTable](
		WarehouseTableOps{},
		core.ProjectScopedImportParser[WarehouseTableTFModel](),
	)
}

type WarehouseTableTFModel struct {
	core.BaseStringIdentifiable
	core.BaseProjectID
	Name         types.String `tfsdk:"name"`
	Format       types.String `tfsdk:"format"`
	URLPattern   types.String `tfsdk:"url_pattern"`
	AccessKey    types.String `tfsdk:"access_key"`
	AccessSecret types.String `tfsdk:"access_secret"`
	CreatedAt    types.String `tfsdk:"created_at"`
}

type WarehouseTableOps struct{}

func (o WarehouseTableOps) ResourceName() string {
	return "warehouse_table"
}

func (o WarehouseTableOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "A self-managed PostHog data warehouse table that queries files in place from an S3-compatible bucket you control " +
			"(Amazon S3, Google Cloud Storage via HMAC keys, Cloudflare R2, and similar). " +
			"PostHog reads the files when the table is created to work out its columns, so the files must exist and the credentials must be able to read them. " +
			"Changing `format` or `url_pattern` replaces the table; the table is queried by `name`, so saved queries keep working. " +
			"Tables synced by a data source are managed with `posthog_external_data_source` instead.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Warehouse table ID (UUID).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"name": schema.StringAttribute{
				Required: true,
				MarkdownDescription: "Name the table is queried by in HogQL. Must be unique within the project, start with a letter or underscore, " +
					"and contain only letters, numbers, and underscores.",
			},
			"format": schema.StringAttribute{
				Required: true,
				MarkdownDescription: "File format of every file the pattern matches: `Parquet`, `CSV`, `CSVWithNames`, `JSONEachRow`, `Delta`, or `DeltaS3Wrapper`. " +
					"PostHog decides which values it accepts. Changing it replaces the table.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"url_pattern": schema.StringAttribute{
				Required: true,
				MarkdownDescription: "HTTPS URL of the files to read, with `*` matching any part of a path segment " +
					"(e.g. `https://your-bucket.s3.us-east-1.amazonaws.com/orders/*.parquet`). All matched files are read as one table. " +
					"Changing it replaces the table.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"access_key": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
				MarkdownDescription: "Access key ID for the bucket (an AWS access key ID, a Google Cloud HMAC key, or the equivalent for another S3-compatible store). " +
					"It needs permission to list the bucket and read the files. PostHog never returns it, so Terraform keeps the configured value " +
					"and cannot detect a key changed outside Terraform; after an import, the first apply sends the configured key.",
			},
			"access_secret": schema.StringAttribute{
				Required:            true,
				Sensitive:           true,
				MarkdownDescription: "Secret for `access_key`. PostHog stores it encrypted and never returns it, so Terraform keeps the configured value.",
			},
			"created_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the table was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (o WarehouseTableOps) BuildCreateRequest(_ context.Context, model WarehouseTableTFModel) (httpclient.WarehouseTableRequest, diag.Diagnostics) {
	return httpclient.WarehouseTableRequest{
		Name:       model.Name.ValueString(),
		Format:     model.Format.ValueString(),
		URLPattern: model.URLPattern.ValueString(),
		Credential: httpclient.WarehouseTableCredential{
			AccessKey:    model.AccessKey.ValueString(),
			AccessSecret: model.AccessSecret.ValueString(),
		},
	}, nil
}

func (o WarehouseTableOps) BuildUpdateRequest(ctx context.Context, plan, _ WarehouseTableTFModel) (httpclient.WarehouseTableRequest, diag.Diagnostics) {
	return o.BuildCreateRequest(ctx, plan)
}

func (o WarehouseTableOps) MapResponseToModel(_ context.Context, resp httpclient.WarehouseTable, model *WarehouseTableTFModel) diag.Diagnostics {
	model.ID = types.StringValue(resp.ID)
	model.Name = types.StringValue(resp.Name)
	model.Format = types.StringValue(resp.Format)
	model.URLPattern = types.StringValue(resp.URLPattern)
	model.CreatedAt = types.StringValue(resp.CreatedAt)
	return nil
}

func (o WarehouseTableOps) Create(ctx context.Context, client httpclient.PosthogClient, model WarehouseTableTFModel, req httpclient.WarehouseTableRequest) (httpclient.WarehouseTable, error) {
	return client.CreateWarehouseTable(ctx, model.GetEffectiveProjectID(), req)
}

func (o WarehouseTableOps) Read(ctx context.Context, client httpclient.PosthogClient, model WarehouseTableTFModel) (httpclient.WarehouseTable, httpclient.HTTPStatusCode, error) {
	return client.GetWarehouseTable(ctx, model.GetEffectiveProjectID(), model.GetID())
}

func (o WarehouseTableOps) Update(ctx context.Context, client httpclient.PosthogClient, model WarehouseTableTFModel, req httpclient.WarehouseTableRequest) (httpclient.WarehouseTable, httpclient.HTTPStatusCode, error) {
	return client.UpdateWarehouseTable(ctx, model.GetEffectiveProjectID(), model.GetID(), req)
}

func (o WarehouseTableOps) Delete(ctx context.Context, client httpclient.PosthogClient, model WarehouseTableTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteWarehouseTable(ctx, model.GetEffectiveProjectID(), model.GetID())
}
