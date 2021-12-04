package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FileMaker() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "help" {
		fmt.Println("usage: ./<binary>  <input file name>")
	}

	fmt.Println("text options: 1: all caps, 2: title case, 3: lower case")

	var option int
	_, err := fmt.Scanf("%d", &option)

	if err != nil {
		fmt.Println("error durin scang")
	}

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		switch option {
		case 1:
			fmt.Println(strings.ToUpper(scanner.Text()))
		case 2:
			fmt.Println(strings.ToTitle(scanner.Text()))
		case 3:
			fmt.Println(strings.ToLower(scanner.Text()))
		}

	}
}