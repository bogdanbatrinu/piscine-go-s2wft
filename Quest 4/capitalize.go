package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	inWord := false

	for i, r := range runes {
		isLetter := (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
		isDigit := r >= '0' && r <= '9'

		if isLetter || isDigit {
			if !inWord {
				if r >= 'a' && r <= 'z' {
					runes[i] = r - ('a' - 'A')
				}
				inWord = true
			} else {
				if r >= 'A' && r <= 'Z' {
					runes[i] = r + ('a' - 'A')
				}
			}
		} else {
			inWord = false
		}
	}

	return string(runes)
}
