package main

import "fmt"

func main() {

	var marks int = 84
	var grade string = "B"

	switch {
	case marks < 35:
		grade = "D"
		fmt.Println("Failed")
	case marks < 35:
		grade = "C"
		fmt.Println("c Grade")
	case marks < 75:
		grade = "B"
		fmt.Println("B Grade")
	default:
		grade = "A"
		fmt.Printf("%s Grade", grade)
	}

}
