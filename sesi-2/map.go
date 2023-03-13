package main

import "fmt"

func main() {
	// Map 1
	var person map[string]string
	person = map[string]string{}

	person["name"] = "Aril"
	person["age"] = "23"
	person["address"] = "jalan sudirman"

	fmt.Println("name:", person["name"])
	fmt.Println("age:", person["age"])
	fmt.Println("address:", person["address"])

	// Map 2
	var person2 = map[string]string{
		"name":    "Aril Noah",
		"age":     "21",
		"address": "jalan kiarapedes",
	}

	fmt.Println("name:", person2["name"])
	fmt.Println("age:", person2["age"])
	fmt.Println("address:", person2["address"])

	fmt.Println("==========================")
	// Map (looping with map)
	for key, value := range person2 {
		fmt.Println(key, ":", value)
	}

	// Map (deleting value)
	fmt.Println("Before deleting", person2)
	delete(person2, "age")
	fmt.Println("After deleting", person2)

	fmt.Println("==========================")
	// Map (detecting value)
	value, exist := person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value does'nt exist")
	}

	delete(person, "age")
	value, exist = person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value has been deleted")
	}

	fmt.Println("==========================")
	// Map (combining slice with map)
	var people = []map[string]string{
		{"name": "Airell", "age": "23"},
		{"name": "nanda", "age": "22"},
		{"name": "Mailo", "age": "15"},
	}

	for i, person := range people{
		fmt.Printf("Index: %d, name: %s, age: %s\n", i, person["name"], person["age"])
	}


}
