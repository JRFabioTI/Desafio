package main

import "fmt"

func main() {
	num := 23

	if fibonacci(num) == num {
		fmt.Printf("O número %d pertence à sequência de Fibonacci.\n", num)
	} else {
		fmt.Printf("O número %d não pertence à sequência de Fibonacci.\n", num)
	}
}
