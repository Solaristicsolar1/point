package piscine

func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return s
	}

	// Check if valid camelCase
	for i, r := range s {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			return s
		}
		if i > 0 && r >= 'A' && r <= 'Z' && s[i-1] >= 'A' && s[i-1] <= 'Z' {
			return s
		}
		if i == len(s)-1 && r >= 'A' && r <= 'Z' {
			return s
		}
	}

	result := ""
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result += "_"
		}
		result += string(r)
	}
	return result
}
