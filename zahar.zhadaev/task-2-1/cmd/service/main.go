package main

import "fmt"

func processQueries(countK int) {
	var sign string
	var degree int

	minimum := 15
	maximum := 30

	for range countK {
		_, err := fmt.Scan(&sign, &degree)
		if err != nil {
			continue
		}

		if sign == "-" {
			if degree > minimum {
				minimum = degree
			}
		}

		if sign == "+" {
			if degree < maximum {
				maximum = degree
			}
		}
	}

	if minimum <= maximum {
		fmt.Println(minimum)
	} else {
		fmt.Println(-1)
	}
}

func main() {
	var countN int
	var countK int

	_, err := fmt.Scan(&countN)
	if err != nil || countN <= 0 {
		return
	}

	for range countN {
		_, err = fmt.Scan(&countK)
		if err != nil || countK <= 0 {
			continue
		}

		processQueries(countK)
	}
}
