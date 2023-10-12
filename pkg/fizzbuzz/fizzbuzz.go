package fizzbuzz

import "fmt"

func FizzBuzz(l int) {
	for i := 1; i <= l; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Fizz")
		} else if i%3 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func FizzBuzzImproved(l int) {
	matchs := []struct {
		s   string
		div int
	}{
		{"Fizz", 5},
		{"Buzz", 3},
		{"Baz", 7},
	}
	tmp := ""
	for i := 1; i <= l; i++ {
		tmp = ""
		for _, v := range matchs {
			if i%v.div == 0 {
				tmp += v.s
			}
		}
		if len(tmp) == 0 {
			fmt.Println(i)
		} else {
			fmt.Println(tmp)
		}

	}
}
