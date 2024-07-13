package concurrency

import (
	"fmt"
	"sync"
)

// GreetingWithChannel receives a list of names,
// iterate over the list and spawns a goroutine for each name
// and prints a greeting message to stdout using channels.
func GreetingWithChannel(names []string) {
	done := make(chan struct{})

	for _, name := range names {
		go func(name string) {
			defer func() { done <- struct{}{} }()
			fmt.Printf("Hello, %s!\n", name)
		}(name)
	}

	for range names {
		<-done
	}
}

func CountingDigitsWithWaitGroup(ch chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- i
	}

	close(ch)
}

func CountingDigitsWithChannels(digitStream chan int, done chan struct{}) {

	for i := 1; i <= 5; i++ {
		digitStream <- i
	}
	close(digitStream) // Close the digitStream channel after writing all numbers
	close(done)        // Close the done channel to signal completion
}

func EnumerateAlphabet(ch chan string, done chan struct{}) {
	for i := 'a'; i <= 'e'; i++ {
		ch <- string(i)
	}
	close(ch)
}
