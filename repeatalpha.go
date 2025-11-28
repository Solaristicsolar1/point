package piscine

func RepeatAlpha(s string) string {
    result := ""

    for _, ch := range s {
        // If character is lowercase
        if ch >= 'a' && ch <= 'z' {
            count := int(ch - 'a' + 1)
            for i := 0; i < count; i++ {
                result += string(ch)
            }
        } else if ch >= 'A' && ch <= 'Z' {
            // If character is uppercase
            count := int(ch - 'A' + 1)
            for i := 0; i < count; i++ {
                result += string(ch)
            }
        } else {
            // Non-alphabet = print once
            result += string(ch)
        }
    }

    return result
}
