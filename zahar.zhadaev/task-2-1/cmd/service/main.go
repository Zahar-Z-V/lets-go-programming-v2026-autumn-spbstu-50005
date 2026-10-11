package main

import "fmt"

func main() {
	var testsCount int

	_, err := fmt.Scan(&testsCount)
	if err != nil || testsCount <= 0 {
		return
	}

	for range testsCount {
		processTest()
	}
}

func processTest() {
	var limitsCount int

	_, err := fmt.Scan(&limitsCount)
	if err != nil || limitsCount <= 0 {
		return
	}

	minimum := 15
	maximum := 30

	for range limitsCount {
		var sign string

		var degree int

		_, err := fmt.Scan(&sign, &degree)
		if err != nil {
			return
		}

		if sign == ">=" && degree > minimum {
			minimum = degree
		}

		if sign == "<=" && degree < maximum {
			maximum = degree
		}

		if minimum <= maximum {
			fmt.Println(minimum)
		} else {
			fmt.Println(-1)
		}
	}
}
