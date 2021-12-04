package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func FMT() {
	// var first string
	// var last string
	// fmt.Println("whats ur name?")
	// fmt.Scanf("%s %s", &first, &last)
	// fmt.Printf("hello %v %v\n", first, last)

	f, err := os.Open("family.csv")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		var age int
		var name string
		n, err := fmt.Sscanf(scanner.Text(), "%v is %v yrs old", &name, &age)

		if err != nil {
			panic(err)
		}

		if n == 2 {
			fmt.Printf("%v,%v\n", name, age)
		}

	}

}
