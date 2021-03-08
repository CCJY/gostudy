package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 5; i++ {
		jobs <- i
	}

	close(jobs)

	for j := 0; j < 5; j++ {
		fmt.Println(<-results)
	}

}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}

}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func loopChannelMain() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			c2 <- "Every two seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}

func channelCountMain() {
	c := make(chan string)
	go channelCount("sheep", c)
	for msg := range c {
		fmt.Println(msg)
	}
}

func countMain() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		count("sheep")
		wg.Done()
	}()

	wg.Wait()

}

func count(thing string) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func channelCount(thing string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
