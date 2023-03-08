package main

import "fmt"

func main() {

	// Variable menggunakan Type data
	var name string
	var age int = 23

	name = "Airell"
	age = 25

	fmt.Println("Ini adalah namanya ==>", name)
	fmt.Println("Ini adalah umurnya ==>", age)

	// Variable tanpa menggunakan type data
	var nama = "Azos"
	var umur = 25

	fmt.Printf("%T, %T\n", nama, umur)

	// short declaration: membuat variable tanpa menggunakan var
	alamat := "Wanayasa"

	fmt.Printf("%T\n", alamat)

	// Multiple Declaration Variable
	var student1, student2, student3 string = "satu", "dua", "tiga"

	var first, second, thirt int

	first, second, thirt = 1, 2, 3

	huruf, angka := "azis", 23

	fmt.Println(student1, student2, student3)
	fmt.Println(first, second, thirt)
	fmt.Println(huruf, angka)

	// Underscore Variable : digunakan menghilangkan error pada variable yang tidak digunakan
	var firstVariable string
	_ = firstVariable

	var namaa = "Ariel"
	var agee = 29
	var addresss = "jalan sudirman"

	fmt.Printf("Halo nama ku %s, umurku %d, dan aku tingga di %s", namaa, agee, addresss)

}
