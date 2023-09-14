package reverse

func Reverse[T any](s []T) []T {
	for start, end := 0, len(s)-1; start < end; start, end = start+1, end-1 {
		s[start], s[end] = s[end], s[start]
	}
	return s
}
