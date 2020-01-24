package util

// ExpandStringList converts a schema.Set of strings from []interface{} to []string
func ExpandStringList(configured []interface{}) []string {
	vs := make([]string, 0, len(configured))

	for _, v := range configured {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, val)
		}
	}

	return vs
}
