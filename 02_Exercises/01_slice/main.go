package main

import "fmt"

func half(x float64) (float64, bool) {
	return x / 2, x%2.0 == 0
}

func main() {
	z, even := half(5)
	fmt.Println(z, even)

}
