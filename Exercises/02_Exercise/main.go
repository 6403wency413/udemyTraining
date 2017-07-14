package main

import "fmt"

func main() {

	a := 0
	b := 0

	fmt.Print("Please enter a number:")
	fmt.Scan(&a)
	fmt.Print("Please enter a bigger number:")
	fmt.Scan(&b)
	fmt.Println(a, "%", b, "=", b%a)
}
