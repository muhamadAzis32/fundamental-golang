package main

import "fmt"

func main() {

	/*
		### NUMBER
	*/

	// Ada baiknya jika kita menentukan tipe data integer dengan jenis apa yang ingin kita pakai untuk menghindari boros memori
	var first uint8 = 89
	var second int8 = -12

	fmt.Printf("tipe data first : %T\n", first)
	fmt.Printf("tipe data secont : %T\n", second)

	// Float : Mengecilkan digit desimalnya
	var decimalNumber float32 = 3.63

	fmt.Printf("decimal Number: %f\n", decimalNumber)
	fmt.Printf("decimal Number: %.3f\n", decimalNumber)

	/*
		### BOOL
	*/

	var condition bool = true
	fmt.Printf("is it permitted? %t \n", condition)

	/*
		### STRING
	*/

	var message string = `Selamat datang broo
	"Salam kenal" saya azis`
	fmt.Print(message)

}
