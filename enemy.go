package main

type Enemy struct {
	Entity
	health int
}

func NewEnemy() *Enemy {
	enemy := &Enemy{
		health: 3, // Increased health - towers need to hit multiple times
	}
	enemy.symbol = "E"
	enemy.pathIndex = 0
	return enemy
}

func (enemy *Enemy) getHealth() int {
	return enemy.health
}

func (enemy *Enemy) takeDamage(damage *int) {
	enemy.health -= *damage
}

func (enemy *Enemy) move(track *Track) {
	path := track.getPath()

	if len(path) == 0 {
		return
	}

	// Move to next position in path (one step per turn in turn-based)
	if enemy.pathIndex < len(path)-1 {
		enemy.pathIndex++
		nextPos := path[enemy.pathIndex]
		enemy.position = nextPos
	}
}

func (enemy *Enemy) isAtEnd(track *Track) bool {
	path := track.getPath()
	if len(path) == 0 {
		return false
	}
	return enemy.pathIndex >= len(path)-1
}
