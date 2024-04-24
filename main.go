package main

import (
	"fmt"

	"bensoer.com/mymath"
	"rsc.io/quote"
)

func main() {
	const greeting = "Hello World"
	fmt.Println(greeting)

	var expliciteGreeting string = "Hello World! Explicite!"
	fmt.Println(expliciteGreeting)

	const constExpliciteGreeting string = "Constant Explicite Hello World!"
	fmt.Println(constExpliciteGreeting)

	fmt.Println(quote.Go())

	printWithLoops(10)

	callMyMath()

}

func callMyMath() {

	var result int = mymath.AdditionPlusWild(5, 5)
	fmt.Println(result)
}

func printWithLoops(count int) {
	fmt.Println(count)

	for i := 0; i < count; i++ {
		fmt.Printf("Aw Yea %d\n", i)

	}

	iterations := 0
	breaker := false
	for {
		defer fmt.Println("For Loop Deferral")

		fmt.Println("This is an infinite loop!")
		fmt.Printf("Iteration: %d", iterations)

		if iterations < 10 {
			fmt.Println("Iteration 10 Reached. Trigger Breaker!")
			breaker = true
		}

		if breaker {
			fmt.Println("Breaking!")
			break
		}
	}

	fmt.Println("Done Here!")
}
