package piscine

func RetainFirstHalf(str string) string {
	n := len(str)

	if n == 0 {
		return ""
	}
	if n == 1 {
		return str
	}

	half := n / 2
	return str[:half]
}

