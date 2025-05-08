package main

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	result := 1
	for i := 1; i <= nb; i++ {
		next := result * i
		if next/i != result {
			return 0
		}
		result = next
	}
	return result
}
