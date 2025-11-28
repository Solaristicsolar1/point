package piscine

func Gcd(a, b uint) uint {
        // If either number is 0, return 0
        if a == 0 || b == 0 {
                return 0
        }

        // Keep dividing until one becomes 0
        for b != 0 {
                a, b = b, a%b
        }

        return a
}