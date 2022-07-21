package main

import (
	"fmt"
)

func printLuckyNumbers(status []string) {
	for i := 0 + 1; i < len(status); i++ {
		if status[i] == "+" {
			fmt.Print(i)
			fmt.Print(" ")
		}
	}
	fmt.Println(" ")
}

func removeNumber(status []string, index int) {
	counter := 0
	for i := 1; i < len(status); i++ {

		if status[i] != "-" {
			counter++
		}
		if counter%index == 0 {
			status[i] = "-"
		}
	}
}

func main() {

	// status of numbers: + is lucky number, - is not
	status := make([]string, 50)

	// init status - set all to +
	for i := 1; i < len(status); i++ {
		status[i] = "+"
	}
	printLuckyNumbers(status)

	// start algorythm
	removeNumber(status, 2)
	printLuckyNumbers(status)
	removeNumber(status, 3)
	printLuckyNumbers(status)
	removeNumber(status, 7)
	printLuckyNumbers(status)
	removeNumber(status, 9)
}
