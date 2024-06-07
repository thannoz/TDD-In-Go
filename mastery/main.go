package main

import (
	"fmt"
	"mastery/interfaces"
)

func main() {
	var inCal interfaces.Calculator
	inCal = interfaces.SimpleCalculator{}
	fmt.Println(inCal.Add(2, 3))

}
