package core

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"

	"github.com/posthog/terraform-provider/pkg/httpclient"
)

const (
	testOrganizationA = "11111111-2222-3333-4444-555555555555"
	testOrganizationB = "99999999-8888-7777-6666-555555555555"
)

func TestOrganizationIDPlanModifier(t *testing.T) {
	ctx := context.Background()
	existingResource := tfsdk.Plan{Raw: tftypes.NewValue(tftypes.Object{}, map[string]tftypes.Value{})}

	tests := map[string]struct {
		state           types.String
		config          types.String
		wantPlan        types.String
		wantReplacement bool
	}{
		"legacy @current state moved to an id updates in place": {
			state:    types.StringValue(httpclient.CurrentOrganizationAlias),
			config:   types.StringValue(testOrganizationA),
			wantPlan: types.StringValue(testOrganizationA),
		},
		"legacy @current state left unset is recorded at apply": {
			state:    types.StringValue(httpclient.CurrentOrganizationAlias),
			config:   types.StringNull(),
			wantPlan: types.StringUnknown(),
		},
		"another organization id replaces the resource": {
			state:           types.StringValue(testOrganizationA),
			config:          types.StringValue(testOrganizationB),
			wantPlan:        types.StringValue(testOrganizationB),
			wantReplacement: true,
		},
		"the same organization id keeps the resource": {
			state:    types.StringValue(testOrganizationA),
			config:   types.StringValue(testOrganizationA),
			wantPlan: types.StringValue(testOrganizationA),
		},
		"an unset organization keeps the state value": {
			state:    types.StringValue(testOrganizationA),
			config:   types.StringNull(),
			wantPlan: types.StringValue(testOrganizationA),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			plan := tc.config
			if plan.IsNull() {
				plan = types.StringUnknown()
			}
			req := planmodifier.StringRequest{
				ConfigValue: tc.config,
				StateValue:  tc.state,
				PlanValue:   plan,
				Plan:        existingResource,
			}
			resp := &planmodifier.StringResponse{PlanValue: plan}

			OrganizationIDPlanModifier{}.PlanModifyString(ctx, req, resp)

			assert.Equal(t, tc.wantPlan, resp.PlanValue)
			assert.Equal(t, tc.wantReplacement, resp.RequiresReplace)
		})
	}
}

func TestNotCurrentOrganizationAlias(t *testing.T) {
	tests := map[string]struct {
		value     types.String
		wantError bool
	}{
		"@current is rejected": {value: types.StringValue(httpclient.CurrentOrganizationAlias), wantError: true},
		"a uuid is accepted":   {value: types.StringValue(testOrganizationA)},
		"unset is accepted":    {value: types.StringNull()},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			resp := &validator.StringResponse{}
			NotCurrentOrganizationAlias{}.ValidateString(context.Background(), validator.StringRequest{ConfigValue: tc.value}, resp)

			assert.Equal(t, tc.wantError, resp.Diagnostics.HasError())
		})
	}
}

func TestGetEffectiveOrganizationID_LegacyCurrentAlias(t *testing.T) {
	tests := map[string]struct {
		providerDefault string
		want            string
	}{
		"reads as the provider organization": {providerDefault: testOrganizationA, want: testOrganizationA},
		"stays the alias without a provider organization, so the client rejects it": {
			providerDefault: "",
			want:            httpclient.CurrentOrganizationAlias,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			model := BaseOrganizationID{
				OrganizationID:        types.StringValue(httpclient.CurrentOrganizationAlias),
				defaultOrganizationID: tc.providerDefault,
			}

			assert.Equal(t, tc.want, model.GetEffectiveOrganizationID())
		})
	}
}

func TestReplaceLegacyOrganizationIDFromPlan(t *testing.T) {
	tests := map[string]struct {
		state types.String
		want  types.String
	}{
		"legacy @current state takes the planned organization": {
			state: types.StringValue(httpclient.CurrentOrganizationAlias),
			want:  types.StringValue(testOrganizationB),
		},
		"a real organization id is kept": {
			state: types.StringValue(testOrganizationA),
			want:  types.StringValue(testOrganizationA),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			state := testOrganizationScopedModel{BaseOrganizationID: BaseOrganizationID{OrganizationID: tc.state}}
			plan := testOrganizationScopedModel{BaseOrganizationID: BaseOrganizationID{OrganizationID: types.StringValue(testOrganizationB)}}

			replaceLegacyOrganizationIDFromPlan(&state, &plan)

			assert.Equal(t, tc.want, state.OrganizationID)
		})
	}
}
