package piscine

func DigitLen(n, base int) int {
        // Check if base is valid (must be between 2 and 36)
        if base < 2 || base > 36 {
                return -1
        }

        // If number is negative, make it positive
        if n < 0 {
                n = -n
        }

        // Special case: if number is 0, it has 1 digit
        if n == 0 {
                return 1
        }

        // Count how many times we can divide by base
        count := 0
        for n > 0 {
                n = n / base  // Divide by base
                count++       // Count this division
        }

        return count
}