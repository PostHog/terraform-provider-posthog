package core

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Identifiable interface {
	GetID() int64
	HasValidID() bool
}

type BaseIdentifiable struct {
	ID types.Int64 `tfsdk:"id"`
}

func (b BaseIdentifiable) HasValidID() bool {
	return !b.ID.IsNull() && !b.ID.IsUnknown()
}

func (b BaseIdentifiable) GetID() int64 {
	return b.ID.ValueInt64()
}
