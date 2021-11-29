package main

import "fmt"

func types() {

	// primitives
	// var example int
	// example = 23
	// fmt.Println(example)

	// var ex2 float64 = 3.2
	// fmt.Println(ex2)

	// namer := "wtf is happening"
	// fmt.Println(namer)

	// bgoolean :=true
	// // fmt.Println(bgoolean)

	// comple := complex(3,4)
	// // fmt.Println(comple)

	// r, i := real(comple), imag(comple)
	// // fmt.Println(r, i)

	// pointers
	// var justpointers *string = new(string)
	// fmt.Println(justpointers)

	// *justpointers = "dereferencing"

	// fmt.Println(*justpointers)

	addressofpointer := "tester in the hood"
	fmt.Println(addressofpointer)

	ptr := &addressofpointer
	fmt.Println(ptr)
	fmt.Println(*ptr)

	addressofpointer = "tester in the books"
	fmt.Println(addressofpointer)
	fmt.Println(ptr)
	fmt.Println(*ptr)


}
