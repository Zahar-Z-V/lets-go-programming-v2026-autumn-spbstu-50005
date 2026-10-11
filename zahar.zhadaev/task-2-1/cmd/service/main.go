package main

import "fmt"

func main() {

	var totalTestCases int

	_, err := fmt.Scan(&totalTestCases)
	if err != nil || totalTestCases <= 0 {
		return
	}

	for range totalTestCases {
		
		var operationsCount int

		_, err = fmt.Scan(&operationsCount)
		if err != nil || operationsCount <= 0 {
			continue
		}

		minimum := 15
		maximum := 30

		for range operationsCount {
			
			var sign string
			var degree int

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
}
