package main

import (
	"fmt"
	"mastery/concurrency"
)

func main() {
	alphabet := make(chan string)
	done := make(chan struct{})

	go concurrency.PrintingAlphabetWichChannels(alphabet, done)
	go func() {
		for x := range alphabet {
			fmt.Printf("%s\n", x)
		}
		//close(done)
	}()

	<-done
}
