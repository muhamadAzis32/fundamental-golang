package main

import "fmt"

func main() {
	var first bool = 2 < 3
	var second bool = "joey" == "Joey"
	var third bool = 10 != 2.3
	var fourt bool = 11 <= 11

	fmt.Println("Kondisi pertama", first)
	fmt.Println("Kondisi kedua", second)
	fmt.Println("Kondisi ketiga", third)
	fmt.Println("Kondisi keempat", fourt)
}
