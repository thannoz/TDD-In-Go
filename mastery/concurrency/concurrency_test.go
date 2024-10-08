package concurrency

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

// TestGreetingWithChannel is a test function that tests the GreetingWithChannel function by capturing stdout,
// executing GreetingWithChannel on a list of names, and checking the output against expected greetings.
//
// t: a testing.T object for running test functions.
func TestGreetingWithChannel(t *testing.T) {
	names := []string{"John", "Doe", "Jane", "Dean"}

	// Capture what is being written to the standard-output
	// and redirect it to pipe using os.Pipe().
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

	// close the pipe
	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)

	// Restore standart-output to avoid problems.
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

// TestCountingDigitsWithChannels is a test function that tests the CountingDigitsWithChannels
// function by capturing stdout, executing CountingDigitsWithChannels on a list of numbers,
// and checking the output against expected digits.
//
// t: a testing.T object for running test functions.
func TestCountingDigitsWithChannels(t *testing.T) {
	exptectedNumbers := []int{1, 2, 3, 4, 5}

	stdOut := os.Stdout
	_, w, _ := os.Pipe()
	os.Stdout = w

	done := make(chan struct{})
	digitStream := make(chan int)

	go func() {
		CountingDigitsWithChannels(digitStream, done)
	}()

	// Read the result from the digitStream channel & store them in
	// a string builder
	outputChan := make(chan string)
	go func() {
		var builder strings.Builder
		for digit := range digitStream {
			builder.WriteString(fmt.Sprintf("%d\n", digit))
		}
		outputChan <- builder.String()
	}()

	<-done

	_ = w.Close()
	// Restore stdout to avoid problems.
	os.Stdout = stdOut

	result := <-outputChan
	output := strings.TrimSpace(result)

	for _, expectedNumber := range exptectedNumbers {
		if !strings.Contains(output, strconv.Itoa(expectedNumber)) {
			t.Errorf("expected to find %d in output but got %s", expectedNumber, output)
		}
	}
}

func TestPrintingLettersWithChannels(t *testing.T) {
	letters := []string{"A", "B", "C", "D", "E", "F"}

	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	done := make(chan struct{})
	alphabetStream := make(chan string)

	go func() {
		PrintingAlphabetWichChannels(alphabetStream, done)
	}()

	go func() {
		var alphabet string
		for alphabet = range alphabetStream {
			fmt.Printf("%s\n", alphabet)
		}
	}()

	<-done
	_ = w.Close()

	os.Stdout = stdOut

	result, _ := io.ReadAll(r)
	output := string(result)

	for _, letter := range letters {
		if !strings.Contains(output, letter) {
			t.Errorf("expected to find %s in output but got %s", letter, output)
		}
	}
}
