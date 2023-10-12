package binarysearch

func BinarySearch(s []int, t int) int {
	var mid int
	for min, max := 0, len(s)-1; max >= min; {
		mid = min + (max-min)/2

		if s[mid] == t {
			return mid
		} else if s[mid] > t {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return -1
}
