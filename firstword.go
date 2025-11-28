package piscine

func FirstWord(s string) string {
    word := ""

    // Step 1: skip leading spaces
    i := 0
    for i < len(s) && s[i] == ' ' {
        i++
    }

    // Step 2: collect characters until next space
    for i < len(s) && s[i] != ' ' {
        word += string(s[i])
        i++
    }

    return word + "\n"
}
