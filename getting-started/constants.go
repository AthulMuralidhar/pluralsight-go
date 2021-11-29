package main

import "fmt"

// works
// const (
// 	test = iota
// 	test3 = iota
// )

// also works
const (
	test = iota
	test2
)

const (
	test3 = iota
)

func constants() {

	// const test = 1213.123123
	// fmt.Println(test)

	// const c = 3
	// fmt.Println(c + 3)
	// fmt.Println(c + 3.2)

	// const typeconversion int = 4
	// fmt.Println(typeconversion + 3)
	// // without the type conversion, the code errors
	// fmt.Println(float32(typeconversion) + 4.4)

	fmt.Println(test, test2, test3)
}