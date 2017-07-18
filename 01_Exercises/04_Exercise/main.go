package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {
		var x, y int

		x = i % 3
		y = i % 5

		if x == 0 {
			if y == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if y == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Printf("%v \n", i)
		}
	}
}
