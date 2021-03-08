package main

import (
	"fmt"
	"sync"
	"time"
)

func cleanup() {
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup", r)
	}

}

func say(s string) {
	// defer wg.Done()
	defer cleanup()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		if i == 2 {
			panic("Oh dear, a 2")
		}
	}
}

var wg sync.WaitGroup

func CoroutinesDo() {
	wg.Add(1)
	go say("hey")
	wg.Add(1)
	go say("there")
	wg.Wait()
}
