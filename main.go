package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var test Entity = Entity{position: [2]int{1,2}, symbol: "@"}

	fmt.Println(test.position)
	fmt.Println(test.symbol)

	var enemy Enemy = Enemy{Entity: Entity{[2]int{30, 10}, "^"}, health: 2}

	fmt.Println(enemy.getPosition())
}


func readTrackFile() {
	
}