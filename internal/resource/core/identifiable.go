package core

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
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
	return !b.ID.IsNull() && !b.ID.IsUnknown()
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
	return !b.ID.IsNull() && !b.ID.IsUnknown() && b.ID.ValueString() != ""
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
	if !b.ProjectID.IsNull() && !b.ProjectID.IsUnknown() && b.ProjectID.ValueString() != "" {
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
	if !b.OrganizationID.IsNull() && !b.OrganizationID.IsUnknown() && b.OrganizationID.ValueString() != "" {
		return b.OrganizationID.ValueString()
	}
	return b.defaultOrganizationID
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
