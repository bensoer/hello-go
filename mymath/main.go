package mymath

import (
	"fmt"
	"math/rand"
)

func Addition(firstNumber, secondNumber int) int {
	return firstNumber + secondNumber
}

func AdditionPlusWild(firstNumber, secondNumber int) int {
	wildcard := rand.Intn(10)
	fmt.Printf("Wildcard Value: %d\n", wildcard)
	return Addition(firstNumber, secondNumber) + wildcard
}
