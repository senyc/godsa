package fibonacci

import "errors"

func Generate(n int) ([]int, error) {
	arr := make([]int, n)
	if n == 0 {
		return arr, errors.New("Cannot create fibonacci sequence with size 0")
	}
	if n == 1 {
		arr[0] = 1
		return arr, nil
	}
	if n == 2 {
		arr[0] = 1
		arr[1] = 1
		return arr, nil
	}
	arr[0] = 1
	arr[1] = 1

	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr, nil
}
