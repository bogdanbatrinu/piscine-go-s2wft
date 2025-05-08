package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	ls, lf := len(s), len(toFind)
	if lf > ls {
		return -1
	}
	for i := 0; i <= ls-lf; i++ {
		if s[i:i+lf] == toFind {
			return i
		}
	}
	return -1
}
