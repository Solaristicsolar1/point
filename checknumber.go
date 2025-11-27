package piscine

func CheckNumber(arg string) bool {
	// Check if the argument is a number
	for _, char := range arg {
		if char  >= '0' && char <= '9' {
			return true
		}
	}
	return false
}
