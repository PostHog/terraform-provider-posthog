package httpclient

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDashboardTilePatchItemMarshalJSON(t *testing.T) {
	t.Run("encodes boolean show_description when configured", func(t *testing.T) {
		value := false
		item := DashboardTilePatchItem{
			ID:              10,
			ShowDescription: &value,
		}

		data, err := json.Marshal(item)
		require.NoError(t, err)
		assert.JSONEq(t, `{"id":10,"show_description":false}`, string(data))
	})

	t.Run("encodes null show_description when clearing", func(t *testing.T) {
		item := DashboardTilePatchItem{
			ID:                   10,
			ClearShowDescription: true,
		}

		data, err := json.Marshal(item)
		require.NoError(t, err)
		assert.JSONEq(t, `{"id":10,"show_description":null}`, string(data))
	})
}
