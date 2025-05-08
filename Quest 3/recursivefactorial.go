package main

var maxNB int

func init() {
	fact := 1
	for i := 1; ; i++ {
		next := fact * i
		if next/i != fact {
			maxNB = i - 1
			return
		}
		fact = next
	}
}

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb > maxNB {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}

	prev := RecursiveFactorial(nb - 1)
	if prev == 0 {
		return 0
	}
	result := nb * prev
	if result/nb != prev {
		return 0
	}
	return result
}
