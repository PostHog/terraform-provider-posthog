package resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
	"github.com/posthog/terraform-provider/internal/util"
)

const warehouseTableCSVDoubleQuotesOption = "csv_allow_double_quotes"

var warehouseTableFormats = []string{"CSV", "CSVWithNames", "Parquet", "JSONEachRow", "Delta", "DeltaS3Wrapper"}

// NewWarehouseTable builds the posthog_warehouse_table resource.
//
// PostHog infers the table columns from the files only on create. A PATCH to
// `url_pattern` or `format` keeps the old column snapshot, so this resource
// treats both as RequiresReplace. The API never returns the credential, so
// `access_key` and `access_secret` keep their planned values in state.
func NewWarehouseTable() resource.Resource {
	return core.NewGenericResource[WarehouseTableTFModel, map[string]any, httpclient.WarehouseTable](
		WarehouseTableOps{},
		core.ProjectScopedImportParser[WarehouseTableTFModel](),
	)
}

type WarehouseTableTFModel struct {
	core.BaseStringIdentifiable
	core.BaseProjectID
	Name                 types.String `tfsdk:"name"`
	Format               types.String `tfsdk:"format"`
	URLPattern           types.String `tfsdk:"url_pattern"`
	AccessKey            types.String `tfsdk:"access_key"`
	AccessSecret         types.String `tfsdk:"access_secret"`
	CSVAllowDoubleQuotes types.Bool   `tfsdk:"csv_allow_double_quotes"`
	HogQLName            types.String `tfsdk:"hogql_name"`
	CreatedAt            types.String `tfsdk:"created_at"`
}

type WarehouseTableOps struct{}

func (o WarehouseTableOps) ResourceName() string {
	return "warehouse_table"
}

func (o WarehouseTableOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "PostHog data warehouse table that reads files directly from an S3-compatible bucket (AWS S3, Google Cloud Storage, Cloudflare R2, and similar). " +
			"`format` and `url_pattern` are immutable: changes trigger replacement, because PostHog infers the table columns from the files only when it creates the table. " +
			"For connector sources such as Stripe or Postgres, use `posthog_external_data_source`.",
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
				MarkdownDescription: "Name to query the table by in HogQL. Must be unique within the project, start with a letter or underscore, " +
					"and contain only letters, numbers, and underscores.",
			},
			"format": schema.StringAttribute{
				Required: true,
				MarkdownDescription: "File format of the objects that `url_pattern` matches. All matched files must have this format. " +
					"One of `CSV`, `CSVWithNames`, `Parquet`, `JSONEachRow`, `Delta`, `DeltaS3Wrapper`. Cannot be changed after creation.",
				Validators: []validator.String{
					stringvalidator.OneOf(warehouseTableFormats...),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"url_pattern": schema.StringAttribute{
				Required: true,
				MarkdownDescription: "HTTPS URL of the files to read. `*` matches any part of a path segment " +
					"(e.g. `https://your-bucket.s3.amazonaws.com/orders/*.parquet`). PostHog reads all matched files as one table. Cannot be changed after creation.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"access_key": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
				MarkdownDescription: "Access key ID for the bucket (an AWS access key ID, a Google Cloud HMAC key, or the equivalent for another S3-compatible store). " +
					"PostHog never returns this value, so the provider keeps the configured value in state and cannot detect changes made outside Terraform.",
			},
			"access_secret": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
				MarkdownDescription: "Secret for `access_key`. PostHog stores it encrypted and never returns it, " +
					"so the provider keeps the configured value in state and cannot detect changes made outside Terraform.",
			},
			"csv_allow_double_quotes": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "For CSV formats only: set to `true` when the files quote fields with doubled quotes.",
			},
			"hogql_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name the table is queried by in HogQL, as reported by PostHog.",
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

func buildWarehouseTableOptions(model WarehouseTableTFModel) map[string]any {
	options := map[string]any{}
	if v := util.BoolPtrFromValue(model.CSVAllowDoubleQuotes); v != nil {
		options[warehouseTableCSVDoubleQuotesOption] = *v
	}
	return options
}

func (o WarehouseTableOps) BuildCreateRequest(_ context.Context, model WarehouseTableTFModel) (map[string]any, diag.Diagnostics) {
	return map[string]any{
		"name":        model.Name.ValueString(),
		"format":      model.Format.ValueString(),
		"url_pattern": model.URLPattern.ValueString(),
		"credential": map[string]any{
			"access_key":    model.AccessKey.ValueString(),
			"access_secret": model.AccessSecret.ValueString(),
		},
		"options": buildWarehouseTableOptions(model),
	}, nil
}

func (o WarehouseTableOps) BuildUpdateRequest(_ context.Context, plan, state WarehouseTableTFModel) (map[string]any, diag.Diagnostics) {
	req := map[string]any{
		"name":    plan.Name.ValueString(),
		"options": buildWarehouseTableOptions(plan),
	}

	credential := map[string]any{}
	if !plan.AccessKey.Equal(state.AccessKey) {
		credential["access_key"] = plan.AccessKey.ValueString()
	}
	if !plan.AccessSecret.Equal(state.AccessSecret) {
		credential["access_secret"] = plan.AccessSecret.ValueString()
	}
	if len(credential) > 0 {
		req["credential"] = credential
	}

	return req, nil
}

func (o WarehouseTableOps) MapResponseToModel(_ context.Context, resp httpclient.WarehouseTable, model *WarehouseTableTFModel) diag.Diagnostics {
	model.ID = types.StringValue(resp.ID)
	model.Name = types.StringValue(resp.Name)
	model.Format = types.StringValue(resp.Format)
	model.URLPattern = types.StringValue(resp.URLPattern)
	model.HogQLName = core.PtrToStringNullIfEmptyTrimmed(resp.HogQLName)
	model.CreatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.CreatedAt)

	if v, ok := resp.Options[warehouseTableCSVDoubleQuotesOption].(bool); ok {
		model.CSVAllowDoubleQuotes = types.BoolValue(v)
	} else {
		model.CSVAllowDoubleQuotes = types.BoolNull()
	}

	return nil
}

func (o WarehouseTableOps) Create(ctx context.Context, client httpclient.PosthogClient, model WarehouseTableTFModel, req map[string]any) (httpclient.WarehouseTable, error) {
	return client.CreateWarehouseTable(ctx, model.GetEffectiveProjectID(), req)
}

func (o WarehouseTableOps) Read(ctx context.Context, client httpclient.PosthogClient, model WarehouseTableTFModel) (httpclient.WarehouseTable, httpclient.HTTPStatusCode, error) {
	return client.GetWarehouseTable(ctx, model.GetEffectiveProjectID(), model.GetID())
}

func (o WarehouseTableOps) Update(ctx context.Context, client httpclient.PosthogClient, model WarehouseTableTFModel, req map[string]any) (httpclient.WarehouseTable, httpclient.HTTPStatusCode, error) {
	return client.UpdateWarehouseTable(ctx, model.GetEffectiveProjectID(), model.GetID(), req)
}

func (o WarehouseTableOps) Delete(ctx context.Context, client httpclient.PosthogClient, model WarehouseTableTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteWarehouseTable(ctx, model.GetEffectiveProjectID(), model.GetID())
}
