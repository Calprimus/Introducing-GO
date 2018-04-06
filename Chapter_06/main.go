package main

import "fmt"

// func main() {
// 	elements := map[string]string{
// 		"H":  "Hydrogen",
// 		"He": "Helium",
// 		"Li": "Lithium",
// 		"Be": "Beryllium",
// 		"B":  "Boron",
// 		"C":  "Carbon",
// 		"N":  " Nitrogen",
// 		"O":  "Oxygen",
// 		"F":  "Fluorine",
// 		"Ne": "Neon",
// 	}
// 	if name, ok := elements["No"]; ok {
// 		fmt.Println(name, ok)
// 	}
// }

// 1. sum is a function that takes a slice of numbers and adds them together. What
// would its function signature look like in Go?
// func sum(xs []int) int.

// 2. Write a function that takes an integer and halves it and returns true if it was even
// or false if it was odd. For example, half(1) should return (0, false) and
// half(2) should return (1, true).
// func half(x int) (int, bool) {
// 	return x / 2, x%2 == 0
// }

// func main() {
// 	fmt.Println(half(2))
// }

// 3. Write a function with one variadic parameter that finds the greatest number in a
// list of numbers.
// func max(xs ...int) int {
// 	var max int
// 	for i, x := range xs {
// 		if i == 0 || x > max {
// 			max = x
// 		}
// 	}
// 	return max
// }

// func main() {
// 	fmt.Println(max(2, 3, 6, 8, 16, 2))
// }

// 4. Using makeEvenGenerator as an example, write a makeOddGenerator function
// that generates odd numbers.
// func makeOddGenerator() func() uint {
// 	i := uint(1)
// 	return func() (ret uint) {
// 		ret = i
// 		i += 2
// 		return
// 	}
// }

// func main() {
// 	nextOdd := makeOddGenerator()
// 	fmt.Println(nextOdd())
// 	fmt.Println(nextOdd())
// 	fmt.Println(nextOdd())
// 	fmt.Println(nextOdd())
// }

// 5. The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) =
// fib(n-1) + fib(n-2). Write a recursive function that can find fib(n).
func fibonacci(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	fmt.Println(fibonacci(10))
}

// 6. What are defer, panic, and recover? How do you recover from a runtime panic?

// 7. How do you get the memory address of a variable?
// To get the memory address of a variable, use the & operator: y = &x.

// 8. How do you assign a value to a pointer?
// To assign a value to a pointer, use the * operator: *x = 5.

// 9. How do you create a new pointer?
// To create a pointer, use the new function: x = new(int).

// 10. What is the value of x after running this program:
// 2.25

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

func swap(x, y *int) {
	*x, *y = *y, *x
}
