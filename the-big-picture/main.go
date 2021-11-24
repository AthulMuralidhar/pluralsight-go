package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	path := flag.String("path", "./the-big-picture/text.txt", "the path to the log file")
	log_level := flag.String("level", "error", "log level to search for")

	flag.Parse()

	file, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(line, *log_level) {
			fmt.Println(line)
		}
	}
}