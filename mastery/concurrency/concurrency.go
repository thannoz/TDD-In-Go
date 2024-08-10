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

func PrintingAlphabetWichChannels(alphabetStream chan string, done chan struct{}) {
	for i := 'A'; i <= 'F'; i++ {
		alphabetStream <- string(i)
	}
	close(alphabetStream)
	close(done)
}

func PrintDigitInBufferedChannel(digitStream chan int, done chan struct{}) {
	for i := 1; i <= 15; i++ {
		digitStream <- i
	}
	close(digitStream)
	close(done)
}

func sender(names []string, nameStream chan string) {

	for _, name := range names {
		nameStream <- name
	}
	close(nameStream)
}

func receiver(nameStream chan string, done chan struct{}) {

	for name := range nameStream {
		fmt.Printf("received from sender: %s\n", name)
	}
	done <- struct{}{}
}

func SendeDataBetweenTwoUnbufferedChannels(done chan struct{}) {
	dataStream := make(chan string, 3)
	data := []string{"Carlos", "Kalonji", "Nzolani", "Konzo"}

	go sender(data, dataStream)
	go receiver(dataStream, done)

	<-done
}
