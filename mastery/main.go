package main

import (
	"fmt"
	"mastery/concurrency"
)

func main() {
	alphabet := make(chan string)
	done := make(chan struct{})

	go concurrency.EnumerateAlphabet(alphabet, done)
	go func() {
		for x := range alphabet {
			fmt.Printf("%s\t", x)
		}
		fmt.Println("\n")
		close(done)
	}()

	<-done
	fmt.Println("Main done...")
}
