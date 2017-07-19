package main

import "fmt"

func main() {
	xs := []int{34, 56, 13, 78, 42, 99, 4, 548, 3, 3452}
	fmt.Println(greatest(xs...))
}

func greatest(x ...int) int {
	var largest int
	for _, v := range x {
		if v > largest {
			largest = v
		}
	}
	return largest
}
