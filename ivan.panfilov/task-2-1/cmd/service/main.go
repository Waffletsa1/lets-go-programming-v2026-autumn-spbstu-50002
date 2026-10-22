package main

import "fmt"

func processDepartment(k int) {
	minTemp := 15
	maxTemp := 30

	for range k {
		var operator string

		var temp int

		if _, err := fmt.Scan(&operator, &temp); err != nil {
			fmt.Println("Invalid temperature setting")

			return
		}

		switch operator {
		case ">=":
			if temp > minTemp {
				minTemp = temp
			}
		case "<=":
			if temp < maxTemp {
				maxTemp = temp
			}
		default:
			fmt.Println("Invalid operator")

			continue
		}

		if minTemp > maxTemp {
			fmt.Println(-1)
		} else {
			fmt.Println(minTemp)
		}
	}
}

func main() {
	var n int

	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Invalid department number")

		return
	}

	for range n {
		var k int

		if _, err := fmt.Scan(&k); err != nil {
			fmt.Println("Invalid worker number")

			return
		}

		processDepartment(k)
	}
}
