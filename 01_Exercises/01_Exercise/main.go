package main

import "fmt"

func main() {

	a := "Erwin"

	fmt.Print("Please enter your name:")
	fmt.Scan(&a)
	fmt.Print("Hello ")
	fmt.Printf("%v \n", a)
}
