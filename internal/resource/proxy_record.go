package resource

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
	"github.com/posthog/terraform-provider/internal/util"
)

func NewProxyRecord() resource.Resource {
	return core.NewGenericResource[ProxyRecordTFModel, httpclient.ProxyRecordRequest, httpclient.ProxyRecord](
		ProxyRecordOps{},
		core.OrganizationScopedImportParser[ProxyRecordTFModel](),
	)
}

type ProxyRecordTFModel struct {
	core.BaseStringIdentifiable
	core.BaseOrganizationID
	Domain      types.String `tfsdk:"domain"`
	TargetCNAME types.String `tfsdk:"target_cname"`
	Status      types.String `tfsdk:"status"`
	Message     types.String `tfsdk:"message"`
	CreatedAt   types.String `tfsdk:"created_at"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
	CreatedBy   types.Int64  `tfsdk:"created_by"`
}

type ProxyRecordOps struct{}

func (o ProxyRecordOps) ResourceName() string {
	return "proxy_record"
}

func (o ProxyRecordOps) Schema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an organization-scoped PostHog custom domain proxy record.",
		Attributes: map[string]schema.Attribute{
			"id":              proxyRecordComputedStringAttribute("The UUID of the proxy record."),
			"organization_id": core.OrganizationIDSchemaAttribute(),
			"domain": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The custom domain to provision in PostHog. Configured values are normalised to lowercase with no trailing dot, so `Example.COM.` and `example.com` plan as the same value.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					normalizeProxyRecordDomainPlanModifier{},
				},
			},
			"target_cname": proxyRecordComputedStringAttribute("The PostHog-managed CNAME target that your DNS record must point to."),
			"status":       proxyRecordComputedStringAttribute("The current provisioning status reported by PostHog."),
			"message":      proxyRecordComputedStringAttribute("Additional status detail returned by PostHog, when present."),
			"created_at":   proxyRecordComputedStringAttribute("Timestamp when the proxy record was created."),
			"updated_at":   proxyRecordComputedStringAttribute("Timestamp when the proxy record was last updated."),
			"created_by":   proxyRecordComputedInt64Attribute("Numeric identifier of the user who created the proxy record."),
		},
	}
}

func (o ProxyRecordOps) BuildCreateRequest(_ context.Context, model ProxyRecordTFModel) (httpclient.ProxyRecordRequest, diag.Diagnostics) {
	var diags diag.Diagnostics
	return httpclient.ProxyRecordRequest{
		Domain: normalizeProxyRecordDomain(model.Domain.ValueString()),
	}, diags
}

func (o ProxyRecordOps) BuildUpdateRequest(_ context.Context, _, _ ProxyRecordTFModel) (httpclient.ProxyRecordRequest, diag.Diagnostics) {
	var diags diag.Diagnostics
	diags.AddError(
		"Update not supported",
		"Proxy records do not support updates. Change the domain to replace the resource instead.",
	)
	return httpclient.ProxyRecordRequest{}, diags
}

func (o ProxyRecordOps) MapResponseToModel(_ context.Context, resp httpclient.ProxyRecord, model *ProxyRecordTFModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.StringValue(resp.ID)
	model.Domain = types.StringValue(resp.Domain)
	model.TargetCNAME = types.StringValue(resp.TargetCNAME)
	model.Status = types.StringValue(resp.Status)
	model.Message = core.PtrToStringNullIfEmptyTrimmed(resp.Message)
	model.CreatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.CreatedAt)
	model.UpdatedAt = core.PtrToStringNullIfEmptyTrimmed(resp.UpdatedAt)
	model.CreatedBy = util.PtrToInt64(resp.CreatedBy)

	return diags
}

func (o ProxyRecordOps) Create(ctx context.Context, client httpclient.PosthogClient, model ProxyRecordTFModel, req httpclient.ProxyRecordRequest) (httpclient.ProxyRecord, error) {
	return client.CreateProxyRecord(ctx, model.GetEffectiveOrganizationID(), req)
}

func (o ProxyRecordOps) Read(ctx context.Context, client httpclient.PosthogClient, model ProxyRecordTFModel) (httpclient.ProxyRecord, httpclient.HTTPStatusCode, error) {
	return client.GetProxyRecord(ctx, model.GetEffectiveOrganizationID(), model.GetID())
}

func (o ProxyRecordOps) Update(ctx context.Context, client httpclient.PosthogClient, _ ProxyRecordTFModel, _ httpclient.ProxyRecordRequest) (httpclient.ProxyRecord, httpclient.HTTPStatusCode, error) {
	return httpclient.ProxyRecord{}, 0, fmt.Errorf("proxy records do not support updates; this is a bug if you see this error")
}

func (o ProxyRecordOps) Delete(ctx context.Context, client httpclient.PosthogClient, model ProxyRecordTFModel) (httpclient.HTTPStatusCode, error) {
	return client.DeleteProxyRecord(ctx, model.GetEffectiveOrganizationID(), model.GetID())
}

func normalizeProxyRecordDomain(raw string) string {
	return strings.ToLower(strings.TrimSuffix(strings.TrimSpace(raw), "."))
}

// normalizeProxyRecordDomainPlanModifier rewrites the planned `domain` value
// through normalizeProxyRecordDomain before plan finalises. Without it, a config
// of `Example.COM.` would plan as `Example.COM.` while state holds the API's
// canonical form (`example.com`), which trips Terraform's "Provider produced
// inconsistent result after apply" check on first apply and triggers a
// permanent RequiresReplace loop on every subsequent plan.
type normalizeProxyRecordDomainPlanModifier struct{}

func (normalizeProxyRecordDomainPlanModifier) Description(_ context.Context) string {
	return "Normalises the domain to lowercase with no trailing dot."
}

func (m normalizeProxyRecordDomainPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (normalizeProxyRecordDomainPlanModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	normalized := normalizeProxyRecordDomain(req.ConfigValue.ValueString())
	if normalized == req.PlanValue.ValueString() {
		return
	}
	resp.PlanValue = types.StringValue(normalized)
}

func proxyRecordComputedStringAttribute(description string) schema.StringAttribute {
	return schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: description,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
}

func proxyRecordComputedInt64Attribute(description string) schema.Int64Attribute {
	return schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: description,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
	}
}
