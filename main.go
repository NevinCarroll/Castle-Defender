package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var test Entity

	var x int = 20
	var y int = 5
	setPosition(&test, &x, &y)

	fmt.Println(test.position)
}
