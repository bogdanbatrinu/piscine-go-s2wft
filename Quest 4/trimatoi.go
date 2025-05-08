package piscine

func TrimAtoi(s string) int {
	var (
		sign    = 1
		result  = 0
		started = false
	)
	for _, r := range s {
		if !started && r == '-' {
			sign = -1
		} else if r >= '0' && r <= '9' {
			started = true
			result = result*10 + int(r-'0')
		}
	}
	return sign * result
}
