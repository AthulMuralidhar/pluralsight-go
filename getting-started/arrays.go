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

	slice := []int{2,5}
	fmt.Println(slice)

	slice = append(slice, 32,234,234,1231,5234)
	fmt.Println(slice)

	// the "colon" like index slicing in python works the same in go
	

}