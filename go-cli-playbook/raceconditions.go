package main

import (
	"fmt"
	"sync"
)

func RaceConditions() {
	var waitGrp sync.WaitGroup

	for i := 0; i < 100; i++ {
		waitGrp.Add(1)
		go func() {
			fmt.Println(i)
			waitGrp.Done()
		}()
		waitGrp.Wait()
	}
}
