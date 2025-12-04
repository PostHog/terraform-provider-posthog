package core

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Identifiable interface {
	GetID() string
	HasValidID() bool
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

type BaseStringIdentifiable struct {
	ID types.String `tfsdk:"id"`
}

func (b BaseStringIdentifiable) HasValidID() bool {
	return !b.ID.IsNull() && !b.ID.IsUnknown() && b.ID.ValueString() != ""
}

func (b BaseStringIdentifiable) GetID() string {
	return b.ID.ValueString()
}
