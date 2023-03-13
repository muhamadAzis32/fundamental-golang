package main

import "fmt"

func main() {
	// slice
	var fruits = []string{"apple", "banna", "mango"}
	_ = fruits

	// Slice (make function)
	var buah = make([]string, 3)
	// fmt.Printf("%#v \n", buah)

	// Slice (append function)
	buah = append(buah, "apple", "banana", "mago")
	fmt.Printf("%#v \n", buah)

	// Slice (append function with ellipsis)
	var fruits1 = []string{"apple", "banna", "mango"}
	var fruits2 = []string{"durian", "pineapple", "starfruit"}

	fruits1 = append(fruits1, fruits2...)
	fmt.Printf("%#v \n", fruits1)

	fmt.Println("================= \n")

	// Slice (copy function)
	nn := copy(fruits1, fruits2)
	fmt.Println("fruits 1 => ", fruits1)
	fmt.Println("fruits 2 => ", fruits2)
	fmt.Println("Copy elements => ", nn)

	fmt.Println("================= \n")

	// Slice (slicing)
	var fruits3 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	var fruits4 = fruits3[1:4]
	fmt.Printf("%#v\n", fruits4)

	var fruits5 = fruits3[0:]
	fmt.Printf("%#v\n", fruits5)

	var fruits6 = fruits3[:3]
	fmt.Printf("%#v\n", fruits6)

	var fruits7 = fruits3[:]
	fmt.Printf("%#v\n", fruits7)

	fmt.Println("================= \n")

	// Slice (combining slicing and append)
	var fruits8 = []string{"apple", "banana", "mango", "durian", "pineapple"}
	fruits8 = append(fruits8[:3], "rambutan")
	fmt.Printf("%#v\n", fruits8)

	fmt.Println("================= \n")

	// Slice (backing array)
	var fruits9 = []string{"apple", "mango", "durian", "banana", "starfruit"}
	var fruits10 = fruits9[2:4]
	fruits10[0] = "rambutan"

	fmt.Println("fruits1 => ", fruits9)
	fmt.Println("fruits2 => ", fruits10)

	fmt.Println("================= \n")

	// Slice (cap function)
	var fruits11 = []string{"apple", "mango", "durian", "banana"}
	fmt.Println("fruits11 => ", cap(fruits11))
	fmt.Println("fruits11 => ", len(fruits11))
	fmt.Println("================= ")

	var fruits12 = fruits11[0:3]
	fmt.Println("fruits12 => ", cap(fruits12))
	fmt.Println("fruits12 => ", len(fruits12))
	fmt.Println("================= ")

	var fruits13 = fruits11[1:]
	fmt.Println("fruits13 => ", cap(fruits13))
	fmt.Println("fruits13 => ", len(fruits13))
}
