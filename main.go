package main

import "fmt"

func main() {
	result := SumTwoNumbers(10, 12)
	fmt.Println("Hello there from Golang", result)

}

func SumTwoNumbers(a int32, b int32) int32 {
	return a + b
}
