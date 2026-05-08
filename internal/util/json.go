package util

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
