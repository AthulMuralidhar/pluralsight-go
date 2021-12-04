package main

import (
	"flag"
	"fmt"
)

func CMDLineFlags() {

	archPtr := flag.String("arch", "x86", "CPU type")

	flag.Parse()

	switch *archPtr {
	case "x86":
		fmt.Println("running in 32 bit mode")
	case "AMD64":
		fmt.Println("running in 64 bit mode")
	default:
		fmt.Println("not recognised")
	}

	fmt.Println("process complete")

}