package piscine

func Compare(a, b string) int {
	la, lb := len(a), len(b)
	for i := 0; i < la && i < lb; i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}
	if la < lb {
		return -1
	} else if la > lb {
		return 1
	}
	return 0
}
