package main

import "fmt"

func main() {

	for i := 0; i < 101; i++ {
		var x int
		x = i % 2
		if x == 0 {
			fmt.Printf("%v \n", i)
		}
	}
}
