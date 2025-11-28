package piscine

func CamelToSnakeCase(s string) string {
        if len(s) == 0 {
                return s
        }

        // Validate: only letters allowed
        for _, r := range s {
                if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
                        return s
                }
        }

        // Check for invalid patterns: consecutive uppercase or ending with uppercase
        for i := 1; i < len(s); i++ {
                curr := s[i]
                prev := s[i-1]

                // Two consecutive uppercase letters
                if curr >= 'A' && curr <= 'Z' && prev >= 'A' && prev <= 'Z' {
                        return s
                }

                // Last character is uppercase
                if i == len(s)-1 && curr >= 'A' && curr <= 'Z' {
                        return s
                }
        }

        // Convert: add underscore before uppercase letters (except first)
        result := ""
        for i, r := range s {
                if i > 0 && r >= 'A' && r <= 'Z' {
                        result += "_"
                }
                result += string(r)
        }

        return result
}
