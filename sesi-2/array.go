package main

import "fmt"

func main() {

	// array
	var number [4]int
	number = [4]int{1, 2, 3, 4}

	var strings = [3]string{"tia", "nadia", "yuli"}

	fmt.Printf("%#v\n", number)
	fmt.Printf("%#v\n", strings)

	// Array (modify element through index)
	var fruits = [3]string{"apel", "pisang", "mangga"}
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "manggo"

	fmt.Printf("%#v\n", fruits)

	// Array (loop through element)
	// cara pertama
	for i, v := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	fmt.Println("=================")

	// cara kedua
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, fruits[i])
	}

	fmt.Println("=================")

	// Array (multidimensional array)
	balance := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balance {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

	// array value
	arr := []interface{}{"test", 1, true, 9.2}
	// fmt.Printf("%#v\n", arr)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Index: %d, Value: %#v\n", i, arr[i])
	}

	

}
