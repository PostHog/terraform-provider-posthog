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

// OrganizationIDSetter is implemented by models that need organization_id set during import.
type OrganizationIDSetter interface {
	SetOrganizationID(string)
}

func (b *BaseOrganizationID) SetOrganizationID(organizationID string) {
	b.OrganizationID = types.StringValue(organizationID)
}
