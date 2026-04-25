package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
)

func NewExternalDataSource() resource.Resource {
	return core.NewGenericResource[ExternalDataSourceTFModel, map[string]interface{}, httpclient.ExternalDataSource](
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
		MarkdownDescription: "PostHog external data warehouse source (e.g. Stripe, Hubspot, Postgres, MySQL, MSSQL, Snowflake, BigQuery, Salesforce, Zendesk, Vitally, Chargebee).",
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
				MarkdownDescription: "Source type. Values accepted by PostHog include " +
					"`Stripe`, `Hubspot`, `Postgres`, `MySQL`, `MSSQL`, `Snowflake`, `BigQuery`, " +
					"`Salesforce`, `Zendesk`, `Vitally`, `Chargebee`, `TemporalIO`. Cannot be changed after creation.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"prefix": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Optional prefix for synced table names (e.g. `stripe_prod_`). Useful when connecting multiple sources of the same type.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"job_inputs_json": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
				MarkdownDescription: "JSON-encoded connection configuration for the source. Shape depends on `source_type`. " +
					"For example Postgres expects `{host, port, database, user, password, schema}`; Stripe expects `{stripe_account_id, stripe_secret_key}`. " +
					"PostHog redacts secret values when reading, so the plan value is always preserved in state.",
			},
			"schemas": schema.ListAttribute{
				Required:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "List of table names to sync from the source (e.g. `[\"users\", \"orders\"]`). PostHog discovers available tables from the source; these must match discovered table names.",
			},
			"sync_frequency": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "How often to sync. Managed per-schema by PostHog (e.g. `5min`, `30min`, `1hour`, `6hour`, `12hour`, `day`, `week`, `never`). Reports the sync frequency of the first schema.",
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

func (o ExternalDataSourceOps) BuildCreateRequest(ctx context.Context, model ExternalDataSourceTFModel) (map[string]interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Parse job_inputs_json into a flat map — these become top-level keys in the payload.
	inputs, d := parseJobInputsJSON(model.JobInputsJSON)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}

	payload := make(map[string]interface{})
	for k, v := range inputs {
		payload[k] = v
	}

	if !model.Prefix.IsNull() && !model.Prefix.IsUnknown() {
		payload["prefix"] = model.Prefix.ValueString()
	}

	// Build schemas array from the list of table names.
	schemaNames, d := extractSchemaNames(ctx, model.Schemas)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}

	schemaEntries := make([]map[string]interface{}, 0, len(schemaNames))
	for _, name := range schemaNames {
		schemaEntries = append(schemaEntries, map[string]interface{}{
			"name":        name,
			"should_sync": true,
		})
	}
	payload["schemas"] = schemaEntries

	req := map[string]interface{}{
		"source_type": model.SourceType.ValueString(),
		"payload":     payload,
	}

	return req, diags
}

func (o ExternalDataSourceOps) BuildUpdateRequest(ctx context.Context, plan, state ExternalDataSourceTFModel) (map[string]interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	req := make(map[string]interface{})

	if !plan.Prefix.IsNull() && !plan.Prefix.IsUnknown() {
		req["prefix"] = plan.Prefix.ValueString()
	} else if core.ShouldClearString(plan.Prefix, state.Prefix) {
		req["prefix"] = ""
	}

	// Rebuild the payload with updated job_inputs if changed.
	inputs, d := parseJobInputsJSON(plan.JobInputsJSON)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}

	payload := make(map[string]interface{})
	for k, v := range inputs {
		payload[k] = v
	}

	if len(payload) > 0 {
		req["payload"] = payload
	}

	return req, diags
}

func (o ExternalDataSourceOps) MapResponseToModel(ctx context.Context, resp httpclient.ExternalDataSource, model *ExternalDataSourceTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.StringValue(resp.ID)
	if resp.SourceType != "" {
		model.SourceType = types.StringValue(resp.SourceType)
	}
	// PostHog may not echo prefix back in the response even though it was
	// accepted on create. Preserve the plan/state value when the API returns null.
	if resp.Prefix != nil && strings.TrimSpace(*resp.Prefix) != "" {
		model.Prefix = types.StringValue(strings.TrimSpace(*resp.Prefix))
	}
	model.Status = core.PtrToStringNullIfEmptyTrimmed(resp.Status)
	model.LastRunAt = core.PtrToStringNullIfEmptyTrimmed(resp.LastRunAt)
	model.CreatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.CreatedAt)

	// Derive sync_frequency from the first schema (PostHog stores it per-schema).
	if len(resp.Schemas) > 0 && resp.Schemas[0].SyncFrequency != nil {
		model.SyncFrequency = types.StringValue(*resp.Schemas[0].SyncFrequency)
	} else if model.SyncFrequency.IsNull() || model.SyncFrequency.IsUnknown() {
		model.SyncFrequency = types.StringNull()
	}

	// Map schema names from the response.
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

	// job_inputs_json contains credentials. PostHog redacts secret values on GET,
	// so echoing the response would produce spurious diffs and "inconsistent values
	// for sensitive attribute" errors. Keep the prior/planned value.
	// On import (state is null), fall back to whatever PostHog returns so the user
	// at least sees non-secret keys.
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

func (o ExternalDataSourceOps) Create(ctx context.Context, client httpclient.PosthogClient, model ExternalDataSourceTFModel, req map[string]interface{}) (httpclient.ExternalDataSource, error) {
	return client.CreateExternalDataSource(ctx, model.GetEffectiveProjectID(), req)
}

func (o ExternalDataSourceOps) Read(ctx context.Context, client httpclient.PosthogClient, model ExternalDataSourceTFModel) (httpclient.ExternalDataSource, httpclient.HTTPStatusCode, error) {
	return client.GetExternalDataSource(ctx, model.GetEffectiveProjectID(), model.GetID())
}

func (o ExternalDataSourceOps) Update(ctx context.Context, client httpclient.PosthogClient, model ExternalDataSourceTFModel, req map[string]interface{}) (httpclient.ExternalDataSource, httpclient.HTTPStatusCode, error) {
	return client.UpdateExternalDataSource(ctx, model.GetEffectiveProjectID(), model.GetID(), req)
}

func (o ExternalDataSourceOps) Delete(ctx context.Context, client httpclient.PosthogClient, model ExternalDataSourceTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteExternalDataSource(ctx, model.GetEffectiveProjectID(), model.GetID())
}

func parseJobInputsJSON(v types.String) (map[string]interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	if v.IsNull() || v.IsUnknown() {
		return nil, diags
	}
	raw := strings.TrimSpace(v.ValueString())
	if raw == "" {
		return map[string]interface{}{}, diags
	}
	var inputs map[string]interface{}
	if err := json.Unmarshal([]byte(raw), &inputs); err != nil {
		diags.AddError("Invalid job_inputs_json", fmt.Sprintf("job_inputs_json must be a valid JSON object: %s", err.Error()))
		return nil, diags
	}
	return inputs, diags
}

func extractSchemaNames(ctx context.Context, list types.List) ([]string, diag.Diagnostics) {
	var diags diag.Diagnostics
	if list.IsNull() || list.IsUnknown() {
		return nil, diags
	}
	var names []string
	diags.Append(list.ElementsAs(ctx, &names, false)...)
	return names, diags
}
