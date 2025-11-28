package piscine

func HashCode(dec string) string {
    result := ""
    size := len(dec)

    for i := 0; i < size; i++ {
        // Step 1: Calculate the hash value
        hash := (int(dec[i]) + size) % 127

        // Step 2: If character is unprintable, shift it up
        if hash < 33 {
            hash += 33
        }

        // Step 3: Add to result
        result += string(rune(hash))
    }

    return result
}
