package main

import "fmt"

func main() {
	num1 := 10
	num2 := 10
	fmt.Println(countOperations(num1, num2))
}
func countOperations(num1 int, num2 int) int {
	operation := 0
	for num1 > 0 && num2 > 0 {
		if num1 >= num2 {
			num1 = num1 - num2
			operation++
		} else {
			num2 = num2 - num1
			operation++
		}
	}

	return operation
}
