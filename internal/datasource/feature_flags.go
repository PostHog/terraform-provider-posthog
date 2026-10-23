package datasource

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/posthog/terraform-provider/internal/data"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
)

func NewFeatureFlags() datasource.DataSource {
	return &FeatureFlagsDataSource{}
}

type FeatureFlagsDataSource struct {
	client                httpclient.PosthogClient
	defaultProjectID      string
	defaultOrganizationID string
}

type FeatureFlagsDataSourceModel struct {
	ProjectID      types.String `tfsdk:"project_id"`
	ProjectName    types.String `tfsdk:"project_name"`
	OrganizationID types.String `tfsdk:"organization_id"`
	Flags          types.List   `tfsdk:"flags"`
}

type featureFlagSummaryModel struct {
	ID       types.Int64  `tfsdk:"id"`
	Key      types.String `tfsdk:"key"`
	Name     types.String `tfsdk:"name"`
	Active   types.Bool   `tfsdk:"active"`
	Archived types.Bool   `tfsdk:"archived"`
}

var featureFlagSummaryAttrTypes = map[string]attr.Type{
	"id":       types.Int64Type,
	"key":      types.StringType,
	"name":     types.StringType,
	"active":   types.BoolType,
	"archived": types.BoolType,
}

var (
	_ datasource.DataSource              = (*FeatureFlagsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*FeatureFlagsDataSource)(nil)
)

func (d *FeatureFlagsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = fmt.Sprintf("%s_feature_flags", req.ProviderTypeName)
}

func (d *FeatureFlagsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "List the feature flags in a project. Use this to resolve the numeric `id` of each flag " +
			"that already exists, so a `for_each` `import {}` block can adopt them into `posthog_feature_flag` " +
			"resources. Deleted flags are left out, as are the flags PostHog creates internally for surveys and product " +
			"tours. Archived flags are included, because an archived flag still holds its key and creating a new flag " +
			"with that key fails.",
		Attributes: map[string]schema.Attribute{
			"project_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project ID (environment) to list flags in. Defaults to the provider-level project_id. Conflicts with `project_name`. When the project is looked up by `project_name` and none matches, this is null.",
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("project_name")),
				},
			},
			"project_name": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					// Names are compared trimmed, so a blank one would match a
					// project whose name is empty.
					stringvalidator.RegexMatches(regexp.MustCompile(`\S`), "must not be blank"),
				},
				MarkdownDescription: "Name of the project to list flags in, matched case-insensitively within `organization_id`. " +
					"Use this when the project ID is not known until apply, for example a project created in the same configuration. " +
					"When no project has this name, `flags` is empty rather than an error. More than one project with the name is an error. " +
					"If a project with this name already exists, import it into your `posthog_project` resource rather than creating a second one: " +
					"before the new project exists, this name matches the old one, and adopting its flags into resources pointed at the new project " +
					"plans to replace them. Terraform deletes the old flags before PostHog refuses the second project, so they are lost even though " +
					"the apply fails. Setting `lifecycle { prevent_destroy = true }` on the flag resources turns this into an error at plan time.",
			},
			"organization_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Organization to look `project_name` up in. Defaults to the provider-level organization_id. Requires `project_name`.",
				Validators: []validator.String{
					stringvalidator.AlsoRequires(path.MatchRoot("project_name")),
				},
			},
			"flags": schema.ListNestedAttribute{
				Computed:            true,
				MarkdownDescription: "The project's flags, sorted by key.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Numeric ID of the flag. Use `<project_id>/<id>` as the `id` in an `import {}` block.",
						},
						"key": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Flag key.",
						},
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Flag name (its description in the PostHog UI).",
						},
						"active": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "Whether the flag is enabled.",
						},
						"archived": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "Whether the flag is archived.",
						},
					},
				},
			},
		},
	}
}

func (d *FeatureFlagsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(data.ProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected ProviderData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = providerData.Client
	d.defaultProjectID = providerData.DefaultProjectID
	d.defaultOrganizationID = providerData.DefaultOrganizationID
}

func (d *FeatureFlagsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config FeatureFlagsDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !config.ProjectName.IsNull() {
		config.OrganizationID = types.StringValue(d.organizationID(config))
	}

	projectID, diags := d.resolveProjectID(ctx, config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	config.ProjectID = projectID

	var flags []httpclient.FeatureFlag
	if !projectID.IsNull() {
		tflog.Debug(ctx, "Listing feature flags", map[string]any{"project_id": projectID.ValueString()})

		var err error
		flags, err = d.client.ListFeatureFlags(ctx, projectID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Error listing feature flags", fmt.Sprintf("Could not list the feature flags in project %s: %s", projectID.ValueString(), err))
			return
		}
	}

	list, diags := featureFlagsToList(ctx, flags)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	config.Flags = list

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// resolveProjectID returns the project to list flags in. It is null only when
// project_name names no project, which Read reports as no flags.
func (d *FeatureFlagsDataSource) resolveProjectID(ctx context.Context, config FeatureFlagsDataSourceModel) (types.String, diag.Diagnostics) {
	var diags diag.Diagnostics

	if config.ProjectName.IsNull() {
		projectID := config.ProjectID.ValueString()
		if projectID == "" {
			projectID = d.defaultProjectID
		}
		if projectID == "" {
			diags.AddError(
				"Missing project_id",
				"project_id or project_name must be set in the data source, or project_id at the provider level",
			)
			return types.StringNull(), diags
		}
		return types.StringValue(projectID), diags
	}

	orgID := d.organizationID(config)
	if orgID == "" {
		diags.AddError(
			"Missing organization_id",
			"project_name is looked up within an organization: set organization_id in the data source or at the provider level",
		)
		return types.StringNull(), diags
	}

	name := strings.TrimSpace(config.ProjectName.ValueString())

	tflog.Debug(ctx, "Looking up project by name", map[string]any{
		"organization_id": orgID,
		"project_name":    name,
	})

	projects, err := d.client.ListProjects(ctx, orgID)
	if err != nil {
		diags.AddError("Error listing projects", fmt.Sprintf("Could not list the projects in organization %s to find %q: %s", orgID, name, err))
		return types.StringNull(), diags
	}

	projectID, found, err := matchProjectByName(projects, name)
	if err != nil {
		diags.AddError("Ambiguous project_name", fmt.Sprintf("%s. Set project_id instead.", err))
		return types.StringNull(), diags
	}
	if !found {
		// Expected on a new environment's first plan, so not a warning; logged
		// so a misspelled name can still be told apart from it.
		tflog.Info(ctx, "No project matched project_name; returning no flags", map[string]any{
			"organization_id": orgID,
			"project_name":    name,
			"projects_seen":   len(projects),
		})
		return types.StringNull(), diags
	}
	return types.StringValue(projectID), diags
}

// organizationID is the organization project_name is looked up in: the data
// source's own, or the provider's.
func (d *FeatureFlagsDataSource) organizationID(config FeatureFlagsDataSourceModel) string {
	if orgID := config.OrganizationID.ValueString(); orgID != "" {
		return orgID
	}
	return d.defaultOrganizationID
}

// matchProjectByName finds the one project whose name equals name, ignoring
// case and surrounding whitespace. Two matches are an error rather than a
// guess, since listing the wrong project's flags would adopt the wrong flags.
func matchProjectByName(projects []httpclient.Project, name string) (string, bool, error) {
	var matches []string
	for _, p := range projects {
		if p.Name != nil && strings.EqualFold(strings.TrimSpace(*p.Name), name) {
			matches = append(matches, strconv.FormatInt(p.ID, 10))
		}
	}

	switch len(matches) {
	case 0:
		return "", false, nil
	case 1:
		return matches[0], true, nil
	default:
		return "", false, fmt.Errorf("%d projects in the organization are named %q (IDs %s)", len(matches), name, strings.Join(matches, ", "))
	}
}

// featureFlagsToList maps API flags onto the flags attribute, leaving out
// deleted flags and sorting by key so the list is stable between reads.
func featureFlagsToList(ctx context.Context, flags []httpclient.FeatureFlag) (types.List, diag.Diagnostics) {
	objectType := types.ObjectType{AttrTypes: featureFlagSummaryAttrTypes}

	kept := make([]httpclient.FeatureFlag, 0, len(flags))
	for _, f := range flags {
		if f.Deleted != nil && *f.Deleted {
			continue
		}
		kept = append(kept, f)
	}
	sort.Slice(kept, func(i, j int) bool { return kept[i].Key < kept[j].Key })

	summaries := make([]featureFlagSummaryModel, 0, len(kept))
	for _, f := range kept {
		summaries = append(summaries, featureFlagSummaryModel{
			ID:       types.Int64Value(f.ID),
			Key:      types.StringValue(f.Key),
			Name:     core.PtrToStringNullIfEmptyTrimmed(f.Name),
			Active:   types.BoolValue(f.Active != nil && *f.Active),
			Archived: types.BoolValue(f.Archived != nil && *f.Archived),
		})
	}

	return types.ListValueFrom(ctx, objectType, summaries)
}
