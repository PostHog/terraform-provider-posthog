package core

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/posthog/terraform-provider/pkg/httpclient"
)

type Identifiable interface {
	GetID() string
	HasValidID() bool
}

type IDSetter interface {
	SetID(string) error
}

type BaseInt64Identifiable struct {
	ID types.Int64 `tfsdk:"id"`
}

func (b BaseInt64Identifiable) HasValidID() bool {
	return IsKnown(b.ID)
}

func (b BaseInt64Identifiable) GetID() string {
	return strconv.FormatInt(b.ID.ValueInt64(), 10)
}

func (b *BaseInt64Identifiable) SetID(id string) error {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid int64 ID %q: %w", id, err)
	}
	b.ID = types.Int64Value(intID)
	return nil
}

type BaseStringIdentifiable struct {
	ID types.String `tfsdk:"id"`
}

func (b BaseStringIdentifiable) HasValidID() bool {
	return IsKnown(b.ID) && b.ID.ValueString() != ""
}

func (b BaseStringIdentifiable) GetID() string {
	return b.ID.ValueString()
}

func (b *BaseStringIdentifiable) SetID(id string) error {
	b.ID = types.StringValue(id)
	return nil
}

type ProjectIDInitializer interface {
	InitializeProjectID(string)
}

// BaseProjectID should be embedded in models that belong to a project.
type BaseProjectID struct {
	ProjectID        types.String `tfsdk:"project_id"`
	defaultProjectID string
}

func (b *BaseProjectID) GetEffectiveProjectID() string {
	if IsKnown(b.ProjectID) && b.ProjectID.ValueString() != "" {
		return b.ProjectID.ValueString()
	}
	return b.defaultProjectID
}

func (b *BaseProjectID) InitializeProjectID(defaultProjectID string) {
	b.defaultProjectID = defaultProjectID
	if b.ProjectID.IsNull() || b.ProjectID.IsUnknown() || b.ProjectID.ValueString() == "" {
		b.ProjectID = types.StringValue(defaultProjectID)
	}
}

type OrganizationIDInitializer interface {
	InitializeOrganizationID(string)
}

// BaseOrganizationID should be embedded in models that belong to an organization (e.g., Project).
type BaseOrganizationID struct {
	OrganizationID        types.String `tfsdk:"organization_id"`
	defaultOrganizationID string
}

func (b *BaseOrganizationID) GetEffectiveOrganizationID() string {
	value := b.OrganizationID.ValueString()
	if !IsKnown(b.OrganizationID) || value == "" {
		return b.defaultOrganizationID
	}
	// State written before the alias was rejected can still hold it until the next apply. Read it as the
	// provider's organization so refresh keeps working. With no provider organization the alias is
	// returned, and the client rejects it with an error that says what to set.
	if value == httpclient.CurrentOrganizationAlias && b.defaultOrganizationID != "" {
		return b.defaultOrganizationID
	}
	return value
}

// legacyOrganizationIDReplacer is implemented by models that embed BaseOrganizationID.
type legacyOrganizationIDReplacer interface {
	currentOrganizationID() types.String
	replaceLegacyOrganizationID(next types.String)
	GetEffectiveOrganizationID() string
}

// holdsLegacyOrganizationID reports whether model's organization_id is still the retired alias.
func holdsLegacyOrganizationID[TFModel any](model *TFModel) bool {
	replacer, ok := any(model).(legacyOrganizationIDReplacer)
	return ok && replacer.currentOrganizationID().ValueString() == httpclient.CurrentOrganizationAlias
}

// effectiveOrganizationID returns the organization model's requests go to, or "" for models that
// are not organization-scoped.
func effectiveOrganizationID[TFModel any](model *TFModel) string {
	if replacer, ok := any(model).(legacyOrganizationIDReplacer); ok {
		return replacer.GetEffectiveOrganizationID()
	}
	return ""
}

func (b *BaseOrganizationID) currentOrganizationID() types.String {
	return b.OrganizationID
}

// replaceLegacyOrganizationID sets next when b still holds the retired alias.
func (b *BaseOrganizationID) replaceLegacyOrganizationID(next types.String) {
	if b.OrganizationID.ValueString() == httpclient.CurrentOrganizationAlias {
		b.OrganizationID = next
	}
}

// replaceLegacyOrganizationIDFromPlan points a state that still holds the retired alias at the
// planned organization, so an update calls the API for the organization the plan names.
func replaceLegacyOrganizationIDFromPlan[TFModel any](state, plan *TFModel) {
	stateReplacer, ok := any(state).(legacyOrganizationIDReplacer)
	if !ok {
		return
	}
	if planReplacer, ok := any(plan).(legacyOrganizationIDReplacer); ok {
		stateReplacer.replaceLegacyOrganizationID(planReplacer.currentOrganizationID())
	}
}

func (b *BaseOrganizationID) InitializeOrganizationID(defaultOrganizationID string) {
	b.defaultOrganizationID = defaultOrganizationID
	if b.OrganizationID.IsNull() || b.OrganizationID.IsUnknown() || b.OrganizationID.ValueString() == "" {
		b.OrganizationID = types.StringValue(defaultOrganizationID)
	}
}

// BaseRoleID should be embedded in models that are scoped to a role (e.g., RoleMembership).
type BaseRoleID struct {
	RoleID types.String `tfsdk:"role_id"`
}

func (b *BaseRoleID) GetRoleID() string {
	return b.RoleID.ValueString()
}

// RoleIDSetter is implemented by models that need role_id set during import.
type RoleIDSetter interface {
	SetRoleID(string)
}

func (b *BaseRoleID) SetRoleID(roleID string) {
	b.RoleID = types.StringValue(roleID)
}

// UserUUIDSetter is implemented by models that need user_uuid set during import.
type UserUUIDSetter interface {
	SetUserUUID(string)
}
