package pow

func Pow(num int, pow int) int {
	if num == 0 && pow == 0 {
		return 1
	}
	if pow == 0 {
		return 1
	}
	return num * Pow(num, pow-1)
}
