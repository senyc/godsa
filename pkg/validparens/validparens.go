package validparens

func ValidParens(p string) bool {
	s := make([]rune, 0)
	m := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	for _, v := range p {
		if p, ok := m[v]; ok {
			s = append(s, p)
		} else if len(s) > 0 && v == s[len(s)-1] {
			s = s[:len(s)-1]
		} else {
			return false
		}
	}
	return len(s) == 0
}
