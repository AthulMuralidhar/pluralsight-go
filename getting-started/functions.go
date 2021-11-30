package main

import (
	"errors"
	"fmt"
)

func Functions() {
	port := 3000

	_, err := startWebServer(port, 10)
	fmt.Println(err)

}

func startWebServer(port, retries int) (int, error) {
	fmt.Println("starting web server...")

	fmt.Println("server started on port:", port)
	fmt.Println("number of retries:", retries)

	return 0, errors.New("erroring at start")
}