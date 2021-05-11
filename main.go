package main

import "fmt"

func main() {
	var message = "Hello, world!"
	var square = calcMe(3)

	printMe(message, square)
}

func printMe(message string, number float32) {
	fmt.Println(message, number)
}

func calcMe(number float32) float32 {
	return number * number
}
