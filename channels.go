package main

import (
	"fmt"
	"sync"
	"time"
)

var cwg sync.WaitGroup

func cleanupFromChannel() {

}

func getValFromChannel(c chan int, someValue int) {
	defer cwg.Done()
	c <- someValue * 5
	time.Sleep(time.Millisecond * 100)
}

func channelExampleDo() {
	footVal := make(chan int, 10)

	for i := 0; i < 10; i++ {
		cwg.Add(1)
		go getValFromChannel(footVal, i)
	}
	cwg.Wait()
	close(footVal)

	for item := range footVal {
		fmt.Println(item)
	}

}
