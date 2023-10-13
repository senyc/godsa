package prefixsumarray

// Generate returns a prefix sum array of the input argument `a`
func Generate(a []int) []int {
	if len(a) == 0 || len(a) == 1 {
		return a
	}
	result := make([]int, len(a))
	result[0] = a[0]
	for i := 1; i < len(a); i++ {
		result[i] = a[i] + result[i-1]
	}
	return result
}
