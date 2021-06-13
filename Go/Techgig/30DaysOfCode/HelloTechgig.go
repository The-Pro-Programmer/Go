/*
 *Read input from STDIN. Print your output to STDOUT
 *Use fmt.Scanf to read input from STDIN and fmt. Println to write output to STDOUT
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var strInput string

	fmt.Println("Hello Techgig")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		strInput = scanner.Text()

	}
	fmt.Println(strInput)

}
