package main

import (
	"fmt"
	"mastery/concurrency"
)

func main() {
	done := make(chan struct{})
	concurrency.SendeDataBetweenTwoUnbufferedChannels(done)

	fmt.Println("\nMain is done...")
}
