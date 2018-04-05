package main

import "fmt"

func main() {
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  " Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	if name, ok := elements["No"]; ok {
		fmt.Println(name, ok)
	}
}

// 1. sum is a function that takes a slice of numbers and adds them together. What
// would its function signature look like in Go?

// func sum(){
// 	return
// }

// func main() {
// sum [](1,2,3,4)
// 	fmt.Println(sum)
// }

// 2. Write a function that takes an integer and halves it and returns true if it was even
// or false if it was odd. For example, half(1) should return (0, false) and
// half(2) should return (1, true).
// 3. Write a function with one variadic parameter that finds the greatest number in a
// list of numbers.
// 4. Using makeEvenGenerator as an example, write a makeOddGenerator function
// that generates odd numbers.
// 5. The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) =
// fib(n-1) + fib(n-2). Write a recursive function that can find fib(n).
// 6. What are defer, panic, and recover? How do you recover from a runtime panic?
// 7. How do you get the memory address of a variable?
// 8. How do you assign a value to a pointer?
// 9. How do you create a new pointer?
// >>>>>>> 10. What is the value of x after running this program:
// func square(x *float64) {
// 	*x = *x * *x
// }
// func main() {
// 	x := 1.5
// 	square(&x)
// 	fmt.Println(x)
// }

// 11. Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y)
// should give you x=2 and y=1).
