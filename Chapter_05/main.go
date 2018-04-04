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

func main() {
	x := []int{
		48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17,
	}
	var sNumber int
	for i, value := range x {
		if i == 0 || value < sNumber {
			sNumber = value
		}
	}
	fmt.Println(sNumber)
}
