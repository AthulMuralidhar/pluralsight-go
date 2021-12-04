package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
	"time"
)

const (
	INFO int = 0 + iota
	WARNING
	ERROR
	FATAL
)

func Logs() {
	// 	writeLog(INFO, "this is testinf")
	// 	writeLog(WARNING, "warning test")
	// 	writeLog(ERROR, "error test")
	// 	writeLog(FATAL,"fatal test")

	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal("cant creaete file")
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal("failed to close trace file")
		}
	}()

	if err:= trace.Start(f); err != nil {
		log.Fatal("failed to start trace ")
	}
	defer trace.Stop()

	addRnd()

	// use go tool trace trace.out to view the trace

}

func addRnd(){
	f := rand.Intn(100)
	s := rand.Intn(100)
	time.Sleep(2 * time.Second)

	fmt.Printf("result %d\n", f*s)
}

// func writeLog(msgType int, msg string) {
// 	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.SetOutput(file)

// 	switch msgType {
// 	case INFO:
// 		logger := log.New(file, "INFO:", log.LUTC|log.Lshortfile)
// 		logger.Println(msg)
// 	case WARNING:
// 		logger := log.New(file, "WARNING:", log.LUTC|log.Lshortfile)
// 		logger.Println(msg)
// 	case ERROR:
// 		logger := log.New(file, "ERROR:", log.LUTC|log.Lshortfile)
// 		logger.Println(msg)
// 	case FATAL:
// 		logger := log.New(file, "FATAL:", log.LUTC|log.Lshortfile)
// 		logger.Fatal(msg)
// 	}

// }
