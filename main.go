package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var track Track
var enemies []*Enemy
var towers []*Tower

var lives int = 3
var gold int = 300

var gameOver bool = false
var manualQuit bool

var enemySpawnCounter int = 0
var waveCount int = 0

const ROW_COUNT = 10
const COLUMN_COUNT = 15
const ENEMIES_PER_WAVE = 1

func main() {
	readTrackFile("./tracks/track.txt")

	// Clear console
	clearConsole()

	// Show tutorial
	showTutorial()

	// Clear console
	clearConsole()

	// Game loop
	turn := 0
	for !gameOver {
		turn++

		// Render current state
		render()

		// Display turn info
		fmt.Printf("\n--- TURN %d ---\n", turn)
		fmt.Print("Commands: (n)ext turn, (p)lace tower, (q)uit: ")

		// Wait for player input
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		command := strings.ToLower(strings.TrimSpace(input))

		switch command {
		case "n", "":
			// Advance turn - enemies move
			advanceTurn()
		case "p":
			// Place tower
			placeTowerMenu()
		case "q":
			gameOver = true
			manualQuit = true
		default:
			fmt.Println("Unknown command")
			turn-- // Don't count invalid commands
		}

		// Check game over condition
		if lives <= 0 {
			gameOver = true
		}
	}

	// Display game over
	displayGameOver()
}

func advanceTurn() {
	// Move enemies
	for i := 0; i < len(enemies); i++ {
		enemies[i].move(&track)

		// Check if enemy reached the end
		if enemies[i].isAtEnd(&track) {
			loseLife()
			// Remove enemy from slice
			enemies = append(enemies[:i], enemies[i+1:]...)
			i--
			continue
		}
	}

	// Towers attack enemies - each tower targets the farthest enemy in range
	for _, tower := range towers {
		var targetEnemy *Enemy = nil
		var maxPathIndex = -1

		// Find the farthest enemy in range
		for _, enemy := range enemies {
			if tower.isEnemyInRange(enemy) && enemy.getPathIndex() > maxPathIndex {
				targetEnemy = enemy
				maxPathIndex = enemy.getPathIndex()
			}
		}

		// Attack the target enemy if found
		if targetEnemy != nil {
			tower.attackEnemy(targetEnemy)
		}
	}

	// Check for dead enemies and remove them
	for i := 0; i < len(enemies); i++ {
		if enemies[i].getHealth() <= 0 {
			earnGold(50) // Earn 50 gold per kill
			// Remove dead enemy
			enemies = append(enemies[:i], enemies[i+1:]...)
			i--
		}
	}

	// Auto-spawn enemies at end of turn
	enemySpawnCounter++
	if enemySpawnCounter >= 2 { // Spawn every 2 turns
		spawnEnemyWave()
		enemySpawnCounter = 0
	}
}

func placeTowerMenu() {
	clearConsole()
	render()

	fmt.Printf("\nTower cost: %d gold (You have: %d)\n", NewTower(0, 0).getCost(), gold)
	fmt.Print("Enter tower position (row 1-10, col 1-15) or -1 to cancel (Example Input - '5,12'): ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "-1" {
		return
	}

	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		fmt.Println("Invalid format")
		return
	}

	var row, col int
	fmt.Sscanf(parts[0], "%d", &row)
	fmt.Sscanf(parts[1], "%d", &col)

	// Convert from 1-based to 0-based coordinates
	row--
	col--

	placeTower(row, col)
}

func placeTower(row int, col int) {
	towerCost := NewTower(0, 0).getCost()

	// Check if enough gold
	if gold < towerCost {
		fmt.Println("Not enough gold!")
		return
	}

	// Check bounds
	if row < 0 || row >= ROW_COUNT || col < 0 || col >= COLUMN_COUNT {
		fmt.Println("Position out of bounds!")
		return
	}

	// Check if on path
	if track.getLayout()[row][col] == "1" || track.getLayout()[row][col] == "*" {
		fmt.Println("Cannot place tower on the path!")
		return
	}

	// Check if tower already exists at position
	for _, tower := range towers {
		pos := tower.getPosition()
		if pos[0] == row && pos[1] == col {
			fmt.Println("Tower already exists here!")
			return
		}
	}

	// Place tower
	newTower := NewTower(row, col)
	towers = append(towers, newTower)
	gold -= towerCost

	fmt.Printf("Tower placed at (%d, %d)! Gold remaining: %d\n", row, col, gold)
}

func spawnEnemyWave() {
	waveCount++
	path := track.getPath()
	if len(path) == 0 {
		return
	}

	for i := 0; i < ENEMIES_PER_WAVE; i++ {
		enemy := NewEnemy()
		enemy.setPosition(&path[0][0], &path[0][1])
		enemies = append(enemies, enemy)
	}

	fmt.Printf("Wave %d spawned with %d enemies!\n", waveCount, ENEMIES_PER_WAVE)
}

func loseLife() {
	lives--
	fmt.Println("\n!!! ENEMY REACHED THE END !!!")
	fmt.Println("Lives remaining:", lives)
}

func earnGold(amount int) {
	gold += amount
	fmt.Printf("Enemy defeated! Earned %d gold. Total: %d\n", amount, gold)
}

func render() {
	clearConsole()

	// Create a copy of the layout for rendering
	var displayLayout [ROW_COUNT][COLUMN_COUNT]string
	originalLayout := track.getLayout()

	// Copy original layout
	for i := 0; i < ROW_COUNT; i++ {
		for j := 0; j < COLUMN_COUNT; j++ {
			displayLayout[i][j] = originalLayout[i][j]
		}
	}

	// Place towers on the layout
	for _, tower := range towers {
		pos := tower.getPosition()
		if pos[0] >= 0 && pos[0] < ROW_COUNT && pos[1] >= 0 && pos[1] < COLUMN_COUNT {
			displayLayout[pos[0]][pos[1]] = tower.getSymbol()
		}
	}

	// Place enemies on the layout
	for _, enemy := range enemies {
		pos := enemy.getPosition()
		if pos[0] >= 0 && pos[0] < ROW_COUNT && pos[1] >= 0 && pos[1] < COLUMN_COUNT {
			displayLayout[pos[0]][pos[1]] = enemy.getSymbol()
		}
	}

	// Render the layout
	for i := 0; i < ROW_COUNT; i++ {
		for j := 0; j < COLUMN_COUNT; j++ {
			print(displayLayout[i][j])
		}
		println("")
	}

	// Display game info
	println("")
	println("")
	fmt.Printf(" Lives: %-5d Gold: %-5d Enemies: %-1d \n", lives, gold, len(enemies))
	fmt.Printf(" Towers: %-5d Waves: %-5d         \n", len(towers), waveCount)
	println("")
}

func displayGameOver() {
	clearConsole()
	println("")

	if manualQuit {
		println("    GAME OVER     ")
	} else {
		println("    GAME OVER     ")
		println("                    ")
		println("  You lost all     ")
		println("     lives!       ")
	}

	println("")
	println("Final Stats:")
	println("Waves spawned:", waveCount)
	println("Final gold:", gold)
}

func clearConsole() {
	// Clear screen - works on Windows
	fmt.Printf("\033[2J\033[H")
}

func showTutorial() {
	clearConsole()
	println("")
	println("                   TOWER DEFENSE TUTORIAL                   ")
	println("")
	println("OBJECTIVE:")
	println("  Prevent enemies from reaching the end of the path (*)")
	println("  You have 3 lives. Lose all lives and it's GAME OVER!")
	println("")
	println("TILE TYPES:")
	println("  0  = Empty space (place towers here)")
	println("  1  = Enemy path (enemies follow this route)")
	println("  *  = End point (enemies reaching here cost a life)")
	println("")
	println("CONTROLS:")
	println("  (n) or Enter = Advance turn - enemies move, towers attack")
	println("  (p) = Place a tower (costs 100 gold)")
	println("  (q) = Quit the game")
	println("")
	println("RESOURCES:")
	println("  Starting Gold: 300")
	println("  Tower Cost: 100 gold each")
	println("  Enemy Reward: 50 gold per kill")
	println("")
	println("")
	fmt.Print("Type 'y' to start the game: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.ToLower(strings.TrimSpace(input))

	for input != "y" {
		fmt.Print("Invalid input. Type 'y' to start: ")
		input, _ = reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))
	}
}

// Reads a track file
func readTrackFile(filePath string) {
	var trackData, err = os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	defer trackData.Close()

	var countScanner = bufio.NewScanner(trackData)

	var rows int
	for countScanner.Scan() {
		rows++
	}

	if rows != ROW_COUNT {
		panic("Invalid row count for file")
	}

	trackData.Seek(0, 0)

	var trackLayout [ROW_COUNT][COLUMN_COUNT]string

	var scanner = bufio.NewScanner(trackData)
	var rowIndex int = 0
	for scanner.Scan() {
		var line []string = strings.Split(scanner.Text(), "")
		trackLayout[rowIndex] = [COLUMN_COUNT]string(line)
		rowIndex++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	track.setLayout(trackLayout)
}
