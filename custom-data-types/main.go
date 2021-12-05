package main

import (
	"fmt"
	"pluralsight-go/custom-data-types/org"
)

func main() {
	// one way - editable and public
	// p := org.Person{}
	// fmt.Println(p.ID())

	// second way - private and not ediable
	// this way protects the struct from changing
	// in a way makes it immutable and points to the
	// interface rather than the stuct itself
	// to make changes, we'll need to impliment getters and setters
	// var p org.Identify = org.Person{} 
	// fmt.Println(p.ID())

	p := org.NewPerson("121312","das", org.NewEUSecurity("12-12312-12312","Some country"))
	fmt.Println(p.GetID())
	fmt.Println(p.Country())
	fmt.Println(p.GetFirst())
	fmt.Println(p.GetLast())

	err := p.SetTwitter("@dsaas")
	fmt.Printf("%T\n", org.TwitterHandler("t"))

	if err != nil {
		fmt.Printf("error: %s \n",err.Error())
	}

	fmt.Println(p.GetTwitter())
	fmt.Println(p.GetTwitter().RedirectURL())

}