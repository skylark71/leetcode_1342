package main

import "fmt"

func main() {

	fmt.Println(numberOfSteps(14))
}

func numberOfSteps(num int) int {
	count := 0
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		count++
	}

	return count
}
