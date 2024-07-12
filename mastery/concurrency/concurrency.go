package concurrency

import "fmt"

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
