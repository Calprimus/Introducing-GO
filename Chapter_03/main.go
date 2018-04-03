package main

import "fmt"

// var F int
// var C int
// var F = C*9/5+32
// var C = (F-32)*5/9

// func main() {
// 	fmt.Print("Enter a number: ")
// 	var input float64
// 	fmt.Scanf("%f", &input)

// 	output := ((input - 32) * 5 / 9)
// 	fmt.Println(output)
// }

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 0.3048
	fmt.Println(output)
}
