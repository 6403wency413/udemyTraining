package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Print("Enter a number: ")
	fmt.Scan(&x)
	f := factorial(x)
	a := result(f)
	fmt.Println("Factorial is:", <-a)
}

func factorial(x int) chan int {
	out := make(chan int)
	go func() {
		for i := x; i > 0; i-- {
			out <- i
		}
		close(out)
	}()
	return out
}

func result(c chan int) chan int {
	out := make(chan int)
	go func() {
		factorial := 1
		for n := range c {
			factorial *= n
		}
		out <- factorial
		close(out)
	}()
	return out
}
