package resource

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/pkg/httpclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjectEventsRetentionLifecycle(t *testing.T) {
	var override *int64
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var body struct {
				RetentionMonths *int64 `json:"retention_months"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Errorf("decode request: %v", err)
				http.Error(w, "invalid request", http.StatusBadRequest)
				return
			}
			override = body.RetentionMonths
		case http.MethodGet:
		default:
			t.Errorf("unexpected method %s", r.Method)
		}
		if err := json.NewEncoder(w).Encode(map[string]any{"project_retention_months": override, "retention_months": 6}); err != nil {
			t.Errorf("encode response: %v", err)
		}
	}))
	defer server.Close()
	ctx := context.Background()
	client := httpclient.NewClient(server.Client(), server.URL, "test-key", "test")
	ops := ProjectEventsRetentionOps{}
	model := ProjectEventsRetentionModel{RetentionMonths: types.Int64Value(12)}
	model.InitializeProjectID("123")
	req, diags := ops.BuildCreateRequest(ctx, model)
	require.False(t, diags.HasError())
	result, err := ops.Create(ctx, client, model, req)
	require.NoError(t, err)
	require.False(t, ops.MapResponseToModel(ctx, result, &model).HasError())
	assert.Equal(t, int64(12), model.RetentionMonths.ValueInt64())
	assert.Equal(t, "123", model.GetID())

	plan := model
	plan.RetentionMonths = types.Int64Null()
	req, diags = ops.BuildUpdateRequest(ctx, plan, model)
	require.False(t, diags.HasError())
	result, _, err = ops.Update(ctx, client, plan, req)
	require.NoError(t, err)
	require.False(t, ops.MapResponseToModel(ctx, result, &model).HasError())
	assert.True(t, model.RetentionMonths.IsNull())
	assert.Nil(t, override)

	override = new(int64(12))
	result, _, err = ops.Read(ctx, client, model)
	require.NoError(t, err)
	require.False(t, ops.MapResponseToModel(ctx, result, &model).HasError())
	assert.Equal(t, int64(12), model.RetentionMonths.ValueInt64())
	_, err = ops.Delete(ctx, client, model)
	require.NoError(t, err)
	require.NotNil(t, override)
	assert.Equal(t, int64(12), *override)
}
