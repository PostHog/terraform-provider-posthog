package resource

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
	"github.com/posthog/terraform-provider/internal/util"
)

// NewExternalDataSource builds the posthog_external_data_source resource.
//
// The PostHog create endpoint (POST /api/projects/{id}/external_data_sources/)
// expects {source_type, prefix?, payload: {connection-fields..., schemas: [...]}}
// and hardcodes the per-schema sync cadence at 6 hours (5 minutes for CDC). The
// update endpoint reads `job_inputs` at the top level of the body and treats
// `prefix` and `schemas` as immutable for non-direct-postgres sources, so we
// model `source_type`, `prefix`, and `schemas` as RequiresReplace and limit
// updates to job_inputs only. `sync_frequency` is exposed read-only because the
// create endpoint cannot accept it; per-schema cadence must be edited via the
// PostHog UI or the schemas API.
func NewExternalDataSource() resource.Resource {
	return core.NewGenericResource[ExternalDataSourceTFModel, map[string]any, httpclient.ExternalDataSource](
		ExternalDataSourceOps{},
		core.ProjectScopedImportParser[ExternalDataSourceTFModel](),
	)
}

type ExternalDataSourceTFModel struct {
	core.BaseStringIdentifiable
	core.BaseProjectID
	SourceType    types.String `tfsdk:"source_type"`
	Prefix        types.String `tfsdk:"prefix"`
	JobInputsJSON types.String `tfsdk:"job_inputs_json"`
	Schemas       types.List   `tfsdk:"schemas"`
	SyncFrequency types.String `tfsdk:"sync_frequency"`
	Status        types.String `tfsdk:"status"`
	LastRunAt     types.String `tfsdk:"last_run_at"`
	CreatedAt     types.String `tfsdk:"created_at"`
}

type ExternalDataSourceOps struct{}

func (o ExternalDataSourceOps) ResourceName() string {
	return "external_data_source"
}

func (o ExternalDataSourceOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "PostHog external data warehouse source (e.g. Stripe, Hubspot, Postgres, MySQL, MSSQL, Snowflake, BigQuery, Salesforce, Zendesk, Vitally, Chargebee). " +
			"`source_type`, `prefix`, and `schemas` are immutable: changes trigger replacement. " +
			"Sync cadence is managed by PostHog (per-schema) and is not configurable on this resource.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "External data source ID (UUID).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": core.ProjectIDSchemaAttribute(),
			"source_type": schema.StringAttribute{
				Required: true,
				MarkdownDescription: "Source type recognised by the PostHog data warehouse (e.g. `Stripe`, `Postgres`, `Snowflake`, `BigQuery`, `Hubspot`). " +
					"PostHog defines the set of accepted values and may add new types over time; see the PostHog data warehouse docs for the current list. Cannot be changed after creation.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"prefix": schema.StringAttribute{
				Optional: true,
				MarkdownDescription: "Optional prefix for synced table names (e.g. `stripe_prod_`). Useful when connecting multiple sources of the same type. " +
					"The PostHog update endpoint silently ignores prefix changes for non-direct-postgres sources, so this resource treats it as RequiresReplace.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"job_inputs_json": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
				MarkdownDescription: "JSON-encoded connection configuration for the source. Shape depends on `source_type`. " +
					"For example Postgres expects `{host, port, database, user, password, schema}`; Stripe expects `{stripe_account_id, stripe_secret_key}`. " +
					"PostHog redacts secret values when reading, so the plan value is preserved in state. " +
					"On import, redacted secrets will appear in state until the next apply with real values.",
			},
			"schemas": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
				MarkdownDescription: "List of table names to sync from the source (e.g. `[\"users\", \"orders\"]`). PostHog discovers available tables from the source; these must match discovered table names. " +
					"The source-level update endpoint cannot edit the schema list, so changes force destroy+recreate.",
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
			},
			"sync_frequency": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: "Sync cadence reported by PostHog. Defaults to 6 hours (5 minutes for CDC schemas) and is not configurable on this resource — the create endpoint does not accept a sync frequency. " +
					"Reports the value of the first schema; sources with mixed schedules will see this flap. Modify per-schema cadence via the PostHog UI or schemas API.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"status": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Current status reported by PostHog (e.g. `Running`, `Completed`, `Error`).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_run_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp of the most recent sync run.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Timestamp when the source was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (o ExternalDataSourceOps) BuildCreateRequest(ctx context.Context, model ExternalDataSourceTFModel) (map[string]any, diag.Diagnostics) {
	var diags diag.Diagnostics

	inputs, d := util.ParseJSONStringMap("job_inputs_json", model.JobInputsJSON)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}

	schemaNames, d := util.StringListToSlice(ctx, model.Schemas)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}

	schemaEntries := make([]map[string]any, 0, len(schemaNames))
	for _, name := range schemaNames {
		schemaEntries = append(schemaEntries, map[string]any{
			"name":        name,
			"should_sync": true,
		})
	}

	payload := make(map[string]any, len(inputs)+1)
	for k, v := range inputs {
		payload[k] = v
	}
	payload["schemas"] = schemaEntries

	req := map[string]any{
		"source_type": model.SourceType.ValueString(),
		"payload":     payload,
	}
	if !model.Prefix.IsNull() && !model.Prefix.IsUnknown() {
		req["prefix"] = model.Prefix.ValueString()
	}

	return req, diags
}

func (o ExternalDataSourceOps) BuildUpdateRequest(_ context.Context, plan, _ ExternalDataSourceTFModel) (map[string]any, diag.Diagnostics) {
	var diags diag.Diagnostics

	inputs, d := util.ParseJSONStringMap("job_inputs_json", plan.JobInputsJSON)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}

	req := make(map[string]any)
	if len(inputs) > 0 {
		req["job_inputs"] = inputs
	}

	return req, diags
}

func (o ExternalDataSourceOps) MapResponseToModel(ctx context.Context, resp httpclient.ExternalDataSource, model *ExternalDataSourceTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.StringValue(resp.ID)
	if resp.SourceType != "" {
		model.SourceType = types.StringValue(resp.SourceType)
	}
	if resp.Prefix != nil && strings.TrimSpace(*resp.Prefix) != "" {
		model.Prefix = types.StringValue(strings.TrimSpace(*resp.Prefix))
	}
	model.Status = core.PtrToStringNullIfEmptyTrimmed(resp.Status)
	model.LastRunAt = core.PtrToStringNullIfEmptyTrimmed(resp.LastRunAt)
	model.CreatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.CreatedAt)

	if len(resp.Schemas) > 0 && resp.Schemas[0].SyncFrequency != nil {
		model.SyncFrequency = types.StringValue(*resp.Schemas[0].SyncFrequency)
	} else if model.SyncFrequency.IsNull() || model.SyncFrequency.IsUnknown() {
		model.SyncFrequency = types.StringNull()
	}

	if len(resp.Schemas) > 0 {
		names := make([]string, 0, len(resp.Schemas))
		for _, s := range resp.Schemas {
			names = append(names, s.Name)
		}
		listVal, d := types.ListValueFrom(ctx, types.StringType, names)
		diags.Append(d...)
		if !diags.HasError() {
			model.Schemas = listVal
		}
	} else if model.Schemas.IsNull() || model.Schemas.IsUnknown() {
		model.Schemas = types.ListNull(types.StringType)
	}

	// PostHog redacts sensitive credentials on GET, so echoing the response
	// would produce spurious diffs. Keep the prior plan value when it exists.
	// On import (state is null) fall back to the response so non-secret keys
	// are at least visible — the user must follow up with real secrets.
	if model.JobInputsJSON.IsNull() || model.JobInputsJSON.IsUnknown() {
		if len(resp.JobInputs) > 0 {
			bytes, err := json.Marshal(resp.JobInputs)
			if err != nil {
				diags.AddError("Failed to marshal job_inputs", err.Error())
				return diags
			}
			model.JobInputsJSON = types.StringValue(string(bytes))
		} else {
			model.JobInputsJSON = types.StringValue("{}")
		}
	}

	return diags
}

func (o ExternalDataSourceOps) Create(ctx context.Context, client httpclient.PosthogClient, model ExternalDataSourceTFModel, req map[string]any) (httpclient.ExternalDataSource, error) {
	return client.CreateExternalDataSource(ctx, model.GetEffectiveProjectID(), req)
}

func (o ExternalDataSourceOps) Read(ctx context.Context, client httpclient.PosthogClient, model ExternalDataSourceTFModel) (httpclient.ExternalDataSource, httpclient.HTTPStatusCode, error) {
	return client.GetExternalDataSource(ctx, model.GetEffectiveProjectID(), model.GetID())
}

func (o ExternalDataSourceOps) Update(ctx context.Context, client httpclient.PosthogClient, model ExternalDataSourceTFModel, req map[string]any) (httpclient.ExternalDataSource, httpclient.HTTPStatusCode, error) {
	return client.UpdateExternalDataSource(ctx, model.GetEffectiveProjectID(), model.GetID(), req)
}

func (o ExternalDataSourceOps) Delete(ctx context.Context, client httpclient.PosthogClient, model ExternalDataSourceTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteExternalDataSource(ctx, model.GetEffectiveProjectID(), model.GetID())
}

