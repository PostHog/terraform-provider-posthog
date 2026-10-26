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
		{"empty", types.StringValue(""), strPtr(""), types.StringValue("")},
		{"omitted", types.StringUnknown(), strPtr(""), types.StringValue("")},
		{"import empty", types.StringNull(), strPtr(""), types.StringValue("")},
		{"null API", types.StringValue("old"), nil, types.StringNull()},
		{"unchanged spaces", types.StringValue(" padded "), strPtr(" padded "), types.StringValue(" padded ")},
		{"API trimmed spaces", types.StringValue(" padded \n"), strPtr("padded"), types.StringValue(" padded \n")},
		{"whitespace only", types.StringValue(" \t\n"), strPtr(""), types.StringValue(" \t\n")},
		{"remote edit", types.StringValue(" old "), strPtr("new"), types.StringValue("new")},
		{"remote inner whitespace edit", types.StringValue("a  b"), strPtr("a b"), types.StringValue("a b")},
		{"import exact API", types.StringNull(), strPtr(" padded "), types.StringValue(" padded ")},
	} {
		t.Run(tc.name, func(t *testing.T) {
			model := FeatureFlagTFModel{Name: tc.prior}
			diags := (FeatureFlagOps{}).MapResponseToModel(context.Background(), httpclient.FeatureFlag{ID: 1, Key: "test", Name: tc.api}, &model)
			require.False(t, diags.HasError(), diags)
			assert.Equal(t, tc.want, model.Name)
		})
	}
}
