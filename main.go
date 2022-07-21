package main

import (
	"fmt"
)

func main() {

	var status [100]string

	// remove every 2nd number

	for i := 1; i < len(status); i++ {

		if i%2 == 0 {
			status[i] = "-"
		} else {
			status[i] = "+"
		}
	}

	for i := 4; i < len(status); i++ {

		if i%3 == 0 {
			status[i] = "-"
		}
	}

	for i := 8; i < len(status); i++ {

		if i%7 == 0 {
			status[i] = "-"
		}
	}

	for i := 10; i < len(status); i++ {

		if i%9 == 0 {
			status[i] = "-"
		}
	}

	for i := 1; i < len(status); i++ {
		fmt.Printf("%d : %s", i, status[i])
		fmt.Println()
	}
}
