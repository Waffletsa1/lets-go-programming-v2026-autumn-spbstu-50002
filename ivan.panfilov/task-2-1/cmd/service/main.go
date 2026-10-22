package main

import "fmt"

func processDepartment(K int) {
	minTemp := 15
	maxTemp := 30
	for range K {
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
	var N int
	if _, err := fmt.Scan(&N); err != nil {
		fmt.Println("Invalid department number")
		return
	}
	for range N {
		var K int
		if _, err := fmt.Scan(&K); err != nil {
			fmt.Println("Invalid worker number")
			return
		}
		processDepartment(K)
	}
}
