package main

import (
	"bufio"
	"fmt"
	"os"
)

func CreatingCLIApps() {
	// fmt.Println("current version of go:" + runtime.Version())
	// fmt.Print("current version of go:" + runtime.Version())
	// fmt.Print("\n")

	// fmt.Printf("the value of go is: %v \n", runtime.Version())

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("name? ")
	// text, _ := reader.ReadString('\n')
	// fmt.Printf("hello: %v \n", text)

	// var name string
	// fmt.Println("what is your name?")
	// inp, _ := fmt.Scanf("%s", &name)

	// switch inp {
	// case 0:
	// 	fmt.Println("enter values")
	// case 1:
	// 	fmt.Printf("hi %v, u entered: %v values \n", name, inp)
	// default:
	// 	fmt.Println("we only accept either 0 or 1 inputs for now")
	// }
	

	// buffio
	// scanner := bufio.NewScanner(os.Stdin)

	// for scanner.Scan() {
	// 	if scanner.Text() == "q"{
	// 		fmt.Println("quitting...")
	// 		os.Exit(3)
	// 	}
	// 	fmt.Println("typed: " + scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
	// 	fmt.Println(err)
	// }

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}


}
