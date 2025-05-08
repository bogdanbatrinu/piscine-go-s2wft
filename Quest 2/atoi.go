package main

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	sign := 1
	i := 0
	switch s[0] {
	case '+':
		sign = 1
		i = 1
	case '-':
		sign = -1
		i = 1
	}

	if i == 1 && len(s) == 1 {
		return 0
	}
	result := 0
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0
		}
		result = result*10 + int(c-'0')
	}
	return sign * result
}
