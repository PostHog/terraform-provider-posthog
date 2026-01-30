package core

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestDefaultBoolFalse(t *testing.T) {
	ctx := context.Background()
	modifier := DefaultBoolFalse{}

	tests := map[string]struct {
		configValue types.Bool
		stateValue  types.Bool
		wantValue   types.Bool
	}{
		"config null, state null - defaults to false": {
			configValue: types.BoolNull(),
			stateValue:  types.BoolNull(),
			wantValue:   types.BoolValue(false),
		},
		"config null, state true - defaults to false (triggers restore)": {
			configValue: types.BoolNull(),
			stateValue:  types.BoolValue(true),
			wantValue:   types.BoolValue(false),
		},
		"config null, state false - defaults to false": {
			configValue: types.BoolNull(),
			stateValue:  types.BoolValue(false),
			wantValue:   types.BoolValue(false),
		},
		"config false - uses config value": {
			configValue: types.BoolValue(false),
			stateValue:  types.BoolValue(true),
			wantValue:   types.BoolValue(false),
		},
		"config true - uses config value": {
			configValue: types.BoolValue(true),
			stateValue:  types.BoolValue(false),
			wantValue:   types.BoolValue(true),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			req := planmodifier.BoolRequest{
				ConfigValue: tc.configValue,
				StateValue:  tc.stateValue,
				PlanValue:   tc.configValue, // Plan starts as config value
			}
			resp := &planmodifier.BoolResponse{
				PlanValue: tc.configValue,
			}

			modifier.PlanModifyBool(ctx, req, resp)

			assert.Equal(t, tc.wantValue, resp.PlanValue)
		})
	}
}
