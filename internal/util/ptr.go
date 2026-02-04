package util

// PtrToString safely dereferences a string pointer, returning empty string if nil.
func PtrToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
