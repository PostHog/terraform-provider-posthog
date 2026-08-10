package util

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ParseJSONStringMap parses a JSON-encoded object value supplied as a Terraform
// string into a map. Used for fields where the user encodes structured config
// as JSON (e.g. job_inputs_json). Null or unknown values produce a nil map; an
// empty string produces an empty map. The field name is interpolated into the
// diagnostic so the error points at the offending attribute.
func ParseJSONStringMap(field string, v types.String) (map[string]any, diag.Diagnostics) {
	var diags diag.Diagnostics
	if v.IsNull() || v.IsUnknown() {
		return nil, diags
	}
	raw := strings.TrimSpace(v.ValueString())
	if raw == "" {
		return map[string]any{}, diags
	}
	var out map[string]any
	if err := json.Unmarshal([]byte(raw), &out); err != nil {
		diags.AddError(
			fmt.Sprintf("Invalid %s", field),
			fmt.Sprintf("%s must be a valid JSON object: %s", field, err.Error()),
		)
		return nil, diags
	}
	return out, diags
}

// ParseJSONStringRaw parses any JSON value supplied as a Terraform string —
// object, array, or bare scalar — and returns it as raw JSON. Used for fields
// whose value is not constrained to one JSON shape (e.g. default_value_json).
// Null, unknown, and empty values produce a nil result, which callers omit from
// the request. The field name is interpolated into the diagnostic so the error
// points at the offending attribute.
func ParseJSONStringRaw(field string, v types.String) (*json.RawMessage, diag.Diagnostics) {
	var diags diag.Diagnostics
	if v.IsNull() || v.IsUnknown() {
		return nil, diags
	}
	raw := strings.TrimSpace(v.ValueString())
	if raw == "" {
		return nil, diags
	}
	if !json.Valid([]byte(raw)) {
		diags.AddError(
			fmt.Sprintf("Invalid %s", field),
			fmt.Sprintf("%s must be valid JSON, for example jsonencode(\"prod\") or jsonencode([\"a\", \"b\"])", field),
		)
		return nil, diags
	}
	msg := json.RawMessage(raw)
	return &msg, diags
}

// JSONNull returns raw JSON null, used to clear a nullable JSON field that the
// user removed from configuration. Omitting the field from a PATCH leaves the
// stored value untouched, so clearing has to be explicit.
func JSONNull() *json.RawMessage {
	msg := json.RawMessage("null")
	return &msg
}

// StripFields recursively returns a copy of v with any map keys named in
// denylist removed. v is expected to be a JSON-decoded value (map, slice,
// or scalar); other types are returned as-is.
func StripFields(v interface{}, denylist map[string]struct{}) interface{} {
	switch val := v.(type) {
	case map[string]interface{}:
		cleaned := make(map[string]interface{}, len(val))
		for key, value := range val {
			if _, isStripped := denylist[key]; isStripped {
				continue
			}
			cleaned[key] = StripFields(value, denylist)
		}
		return cleaned
	case []interface{}:
		cleaned := make([]interface{}, len(val))
		for i, item := range val {
			cleaned[i] = StripFields(item, denylist)
		}
		return cleaned
	default:
		return val
	}
}
