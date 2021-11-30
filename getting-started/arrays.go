package main

import "fmt"

func Arrays() {
	// arrays are const length
	// var arr [3] int
	// fmt.Println(arr)

	// arr[0] = 1
	// arr[1] = 2
	// fmt.Println(arr)
	// fmt.Println(arr[1])

	// arr := [5] int{1,2,3}
	// fmt.Println(arr)

	// // slices have flexible lengths based on arrays

	// slice := arr[:3]
	// fmt.Println(slice)

	// arr[0] = 54
	// slice[0] = 234

	// fmt.Println(arr)
	// fmt.Println(slice)

	// slice := []int{2,5}
	// fmt.Println(slice)

	// slice = append(slice, 32,234,234,1231,5234)
	// fmt.Println(slice)

	// the "colon" like index slicing in python works the same in go

	// maps == dict
	// m := map[string] string{"test": "asd"}
	// fmt.Println(m) 
	// fmt.Println(m["test"])

	// delete(m, "test")
	// fmt.Println(m) 

	// structs - the only hetrogenrous collection in go

	type example struct {
		id int
		firstname string
		lastname string 
	}

	var e1 example
	fmt.Println(e1)

	e1.firstname = "asdad"
	e1.lastname = "asdasda"
	e1.id = 1

	fmt.Println(e1)

	e2 := example{id: 2,
		lastname: "12dsasda", 
		firstname: "asdasd",
	}
	fmt.Println(e2)


}