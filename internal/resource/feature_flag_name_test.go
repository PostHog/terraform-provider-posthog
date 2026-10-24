package resource

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFeatureFlagNameRoundTrip(t *testing.T) {
	for _, tc := range []struct {
		name  string
		prior types.String
		api   *string
		want  types.String
	}{
		{"empty", types.StringValue(""), stringPtr(""), types.StringValue("")},
		{"omitted", types.StringUnknown(), stringPtr(""), types.StringValue("")},
		{"import empty", types.StringNull(), stringPtr(""), types.StringValue("")},
		{"null API", types.StringValue("old"), nil, types.StringNull()},
		{"unchanged spaces", types.StringValue(" padded "), stringPtr(" padded "), types.StringValue(" padded ")},
		{"API trimmed spaces", types.StringValue(" padded \n"), stringPtr("padded"), types.StringValue(" padded \n")},
		{"whitespace only", types.StringValue(" \t\n"), stringPtr(""), types.StringValue(" \t\n")},
		{"API trimmed control separator", types.StringValue("\u001ctest\u001f"), stringPtr("test"), types.StringValue("\u001ctest\u001f")},
		{"remote edit", types.StringValue(" old "), stringPtr("new"), types.StringValue("new")},
		{"import exact API", types.StringNull(), stringPtr(" padded "), types.StringValue(" padded ")},
	} {
		t.Run(tc.name, func(t *testing.T) {
			model := FeatureFlagTFModel{Name: tc.prior}
			diags := (FeatureFlagOps{}).MapResponseToModel(context.Background(), httpclient.FeatureFlag{ID: 1, Key: "test", Name: tc.api}, &model)
			require.False(t, diags.HasError(), diags)
			assert.Equal(t, tc.want, model.Name)
		})
	}
}

func stringPtr(s string) *string { return &s }
