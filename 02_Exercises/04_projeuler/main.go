package main

import "fmt"

func fib(limit int) {
	var i, x, y int
	z := 1
	xs := []int{}
	for i < limit {
		x = y + z
		y = z + x
		z = y + x
		i = z
		if x > limit {
		} else if y > limit {
			xs = append(xs, x)
		} else if z > limit {
			xs = append(xs, x, y)
		} else if z < limit {
			xs = append(xs, x, y, z)
		}
	}
	return
	fmt.Println(xs)
}

func main() {

}
