package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(13))
}

func FindPrevPrime(nb int) int{
	for i := nb; i >= 2; i-- {
		if isPrime(i) {
			return i
		}
	}
	return 0

}

func isPrime(n int) bool{
	for i := 2; i*i <= n; i++ {
		if n%2 == 0 {
			return false
		}
	}
	return true
}
