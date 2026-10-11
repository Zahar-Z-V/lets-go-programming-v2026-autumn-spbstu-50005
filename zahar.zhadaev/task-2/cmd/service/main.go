package main

import "fmt"

func main() {
	var N, K int
	_, err := fmt.Scan(&N)
	if err != nil || N <= 0 {
		return
	}

	for range N {
		_, err := fmt.Scan(&K)
		if err != nil || K <= 0 {
			continue
		}

		minimum := 15
		maximum := 30

		for range K {
			var sign string
			var degree int
			_, err := fmt.Scan(&sign, &degree)
			if err != nil {
				return
			}
			if sign == ">=" {
				if degree > minimum {
					minimum = degree
				}
			}
			if sign == "<=" {
				if degree < maximum {
					maximum = degree
				}
			}
			if minimum <= maximum {
				fmt.Println(minimum)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
