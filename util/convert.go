package util

import (
	"encoding/json"
)

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// Object is a helper routine that allocates a new generic string-
// object map to unmarshal the string value into and returns a
// pointer to it.
func Object(v string) *map[string]interface{} {
	data := new(map[string]interface{})
	json.Unmarshal([]byte(v), data)
	return data
}
