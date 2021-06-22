package main

import "fmt"

func main() {

	var a int
	var b int = 15
	numbers := [6]int{1, 2, 3, 4, 5}

	for a := 0; a < 10; a++ {
		fmt.Printf("value of a: %d \n", a)
	}

	for a < b {
		fmt.Printf("Second loop a: %d\n", a)
		a++
	}

	for i, x := range numbers {
		fmt.Printf("Third loop x: %d, i: %d \n", x, i)
	}

}
