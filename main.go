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

func nextIndex(status []string) int {

	for i := 1; i < len(status); i++ {
		if status[i] == "=" {
			status[i] = "+"
			// TODO: wenn i 1 ist, dann +1
			if i == 1 {
				return i + 1
			}
			return i
		}
	}
	//TODO: was wenn keine glückliche zahl mehr gefunden wurde
	return -1
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

	// status of numbers: + is lucky number, - is not, = not yet processed
	status := make([]string, 200)

	// init status - set all to =
	for i := 1; i < len(status); i++ {
		status[i] = "="
	}
	//	printLuckyNumbers(status)
	fmt.Println(status)

	for {
		// start algorythm
		index := nextIndex(status)
		if index == -1 {
			break
		}
		removeNumber(status, index)
		printLuckyNumbers(status)
	}

}
