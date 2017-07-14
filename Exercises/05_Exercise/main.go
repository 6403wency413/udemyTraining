package main

import "fmt"

func main() {
	var z int
	for i := 1; i < 1000; i++ {
		if i%3 == 0 {
			z = z + i
		} else if i%5 == 0 {
			z = z + i
		}
	}
	fmt.Printf("%v \n", z)
}
