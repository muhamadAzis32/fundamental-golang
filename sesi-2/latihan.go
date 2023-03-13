package main

import "fmt"

func main() {
	// fmt.Printf("=========== Loopings (nested looping) \n\n")

	// for i := 0; i < 5; i++ {
	// 	for j := i; j < 5; j++ {
	// 		fmt.Print(j, " ")
	// 	}
	// 	fmt.Println()
	// }

	// for i := 0; i < 5; i++ {
	// 	for j := 1; j < i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// var arr = [3]int{2, 2, 2}
	// var total int

	// for _, data := range arr {
	// 	total += data
	// }
	// fmt.Print(total)

	// var arr1 = [3]int{1, 2, 3}
	// var arr2 = [3]string{"tiga", "dua", "satu"}

	arr1 := [3]int{1, 2, 3}
	arr2 := [3]string{"satu", "dua", "tiga"}

	for i, v := range arr1 {
		fmt.Println(v, ":", arr2[i])
	}

}
