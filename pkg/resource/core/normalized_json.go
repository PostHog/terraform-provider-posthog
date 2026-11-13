package core

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.StringTypable                    = NormalizedJSONType{}
	_ basetypes.StringValuableWithSemanticEquals = NormalizedJSON{}
)

// NormalizedJSONType is jsontypes.NormalizedType with numbers compared by
// value. PostHog can return a configured 2.0 as 2 (an insight query's
// decimalPlaces does), which jsontypes.Normalized treats as a change because it
// compares number text.
type NormalizedJSONType struct {
	jsontypes.NormalizedType
}

func (t NormalizedJSONType) String() string {
	return "core.NormalizedJSONType"
}

func (t NormalizedJSONType) ValueType(_ context.Context) attr.Value {
	return NormalizedJSON{}
}

func (t NormalizedJSONType) Equal(o attr.Type) bool {
	other, ok := o.(NormalizedJSONType)
	return ok && t.NormalizedType.Equal(other.NormalizedType)
}

func (t NormalizedJSONType) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return NormalizedJSON{Normalized: jsontypes.Normalized{StringValue: in}}, nil
}

func (t NormalizedJSONType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}
	stringValue, ok := attrValue.(basetypes.StringValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}
	stringValuable, diags := t.ValueFromString(ctx, stringValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting StringValue to StringValuable: %v", diags)
	}
	return stringValuable, nil
}

// NormalizedJSON is a JSON string attribute value that ignores key order,
// whitespace, escaping and number formatting (2.0 equals 2, 1e3 equals 1000)
// when Terraform compares it with the value read back from the API.
type NormalizedJSON struct {
	jsontypes.Normalized
}

func NewNormalizedJSONNull() NormalizedJSON {
	return NormalizedJSON{Normalized: jsontypes.NewNormalizedNull()}
}

func NewNormalizedJSONValue(value string) NormalizedJSON {
	return NormalizedJSON{Normalized: jsontypes.NewNormalizedValue(value)}
}

func (v NormalizedJSON) Type(_ context.Context) attr.Type {
	return NormalizedJSONType{}
}

func (v NormalizedJSON) Equal(o attr.Value) bool {
	other, ok := o.(NormalizedJSON)
	return ok && v.StringValue.Equal(other.StringValue)
}

func (v NormalizedJSON) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(NormalizedJSON)
	if !ok {
		diags.AddError(
			"Semantic Equality Check Error",
			fmt.Sprintf("Expected value type %T, got %T. Please report this to the provider developers.", v, newValuable),
		)
		return false, diags
	}

	current, err := decodeJSONKeepingNumbers(v.ValueString())
	if err == nil {
		var proposed any
		proposed, err = decodeJSONKeepingNumbers(newValue.ValueString())
		if err == nil {
			return jsonValuesEqual(current, proposed), diags
		}
	}

	diags.AddError(
		"Semantic Equality Check Error",
		"An unexpected error occurred while comparing JSON values. Please report this to the provider developers.\n\nError: "+err.Error(),
	)
	return false, diags
}

func decodeJSONKeepingNumbers(s string) (any, error) {
	dec := json.NewDecoder(strings.NewReader(s))
	dec.UseNumber()
	var v any
	if err := dec.Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}

// jsonValuesEqual compares decoded JSON. Numbers are compared as exact
// rationals. The API client decodes responses into float64, so a configured
// number float64 cannot hold exactly (an integer above 2^53) still reads back
// as a different value.
func jsonValuesEqual(a, b any) bool {
	switch av := a.(type) {
	case map[string]any:
		bv, ok := b.(map[string]any)
		if !ok || len(av) != len(bv) {
			return false
		}
		for key, aItem := range av {
			bItem, exists := bv[key]
			if !exists || !jsonValuesEqual(aItem, bItem) {
				return false
			}
		}
		return true
	case []any:
		bv, ok := b.([]any)
		if !ok || len(av) != len(bv) {
			return false
		}
		for i := range av {
			if !jsonValuesEqual(av[i], bv[i]) {
				return false
			}
		}
		return true
	case json.Number:
		bv, ok := b.(json.Number)
		if !ok {
			return false
		}
		aRat, aOK := new(big.Rat).SetString(av.String())
		bRat, bOK := new(big.Rat).SetString(bv.String())
		if !aOK || !bOK {
			return av == bv
		}
		return aRat.Cmp(bRat) == 0
	default:
		// string, bool or nil
		return a == b
	}
}
