package main

import (
	"bufio"
	"fmt"
	"mastery/hashmaps"

	"os"
	"strings"
)

func main() {
	/* 	neighbours := hp.PrintNeighbouringCountries("Frankreich")
	   	for _, n := range neighbours {
	   		fmt.Printf("%s\n", n)
	   	} */

	author, err := hashmaps.PrintAuthorsName("Haben oder Sein")
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(author)
}

type language string

var phrasebook = map[language]string{
	"li": "Mbote, mokili",
	"en": "Hello, world",
	"fr": "Bonjour, le monde",
	"de": "Hallo, Welt",
}

func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}

func greetingFromTerminal() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\n")

	return text
}
