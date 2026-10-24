package resource

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFeatureFlagRuntimeRequests(t *testing.T) {
	for _, tc := range []struct {
		name  string
		value types.String
		want  interface{}
	}{
		{"server", types.StringValue("server"), "server"},
		{"client", types.StringValue("client"), "client"},
		{"all", types.StringValue("all"), "all"},
		{"omitted", types.StringNull(), nil},
		{"unknown", types.StringUnknown(), nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			model := FeatureFlagTFModel{Key: types.StringValue("test"), EvaluationRuntime: tc.value}
			create, diags := (FeatureFlagOps{}).BuildCreateRequest(context.Background(), model)
			require.False(t, diags.HasError(), diags)
			update, diags := (FeatureFlagOps{}).BuildUpdateRequest(context.Background(), model, model)
			require.False(t, diags.HasError(), diags)
			for _, req := range []httpclient.FeatureFlagRequest{create, update} {
				raw, err := json.Marshal(req)
				require.NoError(t, err)
				var body map[string]interface{}
				require.NoError(t, json.Unmarshal(raw, &body))
				assert.Equal(t, tc.want, body["evaluation_runtime"])
				if tc.want == nil {
					assert.NotContains(t, body, "evaluation_runtime")
				}
			}
		})
	}
}
