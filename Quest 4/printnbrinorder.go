package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	var freq [10]int
	if n == 0 {
		freq[0] = 1
	} else {
		for n > 0 {
			freq[n%10]++
			n /= 10
		}
	}

	for d := 0; d < 10; d++ {
		for i := 0; i < freq[d]; i++ {
			z01.PrintRune(rune('0' + d))
		}
	}
}
