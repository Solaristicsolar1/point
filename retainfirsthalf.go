package piscine

func RetainFirstHalf(str string) string {
	var n = len(str)

	if n == 0 {
		return ""
	}
	if n == 1 {
		return str
	}

	var half = n / 2
	return str[:half]
}

