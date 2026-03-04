package main

import (
	"bufio"
	"fmt"
	"os"
)

var track Track

func main() {
	fmt.Println("Hello, World!")

	var test Entity = Entity{position: [2]int{1, 2}, symbol: "@"}

	fmt.Println(test.position)
	fmt.Println(test.symbol)

	var enemy Enemy = Enemy{Entity: Entity{[2]int{30, 10}, "^"}, health: 2}

	fmt.Println(enemy.getPosition())

	readTrackFile("./tracks/track.txt")
}

// Reads a track file
func readTrackFile(filePath string) {
	var trackData, err = os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	defer trackData.Close()

	scanner := bufio.NewScanner(trackData)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
