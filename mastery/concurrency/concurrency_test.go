package concurrency

import (
	"io"
	"os"
	"strings"
	"testing"
)

// TestGreetingWithChannel is a test function that tests the GreetingWithChannel function by capturing stdout,
// executing GreetingWithChannel on a list of names, and checking the output against expected greetings.
//
// t: a testing.T object for running test functions.
func TestGreetingWithChannel(t *testing.T) {
	names := []string{"John", "Doe", "Jane", "Dean"}

	// Capture what is being written to stdout
	// use os.Pipe() to read and (write) from it.
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	done := make(chan struct{})

	/*
		Start ONE goroutine in the background. This goroutine will call
		GreetingWithChannel() which then will be executed 4 times (length of names).
		When all (4) goroutines are done, close the done channel.
		When this happens, the main goroutine will unblock (<-done)
		and continue execution...
	*/
	go func() {
		GreetingWithChannel(names)
		defer close(done)
	}()

	<-done

	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)

	// Restore stdout to avoid problems.
	os.Stdout = stdOut

	expectedGreetings := []string{
		"Hello, John!\n",
		"Hello, Doe!\n",
		"Hello, Jane!\n",
		"Hello, Dean!\n",
	}

	for _, greeting := range expectedGreetings {
		if !strings.Contains(output, greeting) {
			t.Errorf("expected to find %s in output", greeting)
		}
	}
}
