package piscine

func LastWord(s string) string {
    i := len(s) - 1

    // Step 1: Skip trailing spaces
    for i >= 0 && s[i] == ' ' {
        i--
    }

    // If string was all spaces
    if i < 0 {
        return "\n"
    }

    // Step 2: Collect the last word (backwards)
    word := ""
    for i >= 0 && s[i] != ' ' {
        word = string(s[i]) + word  // prepend character
        i--
    }

    return word + "\n"
}
