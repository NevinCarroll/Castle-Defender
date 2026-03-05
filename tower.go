package main

type Tower struct {
	Entity
	radius     int
	damage     int
	cost       int
	attackRate int
}

func NewTower(x int, y int) *Tower {
	tower := &Tower{
		radius:     3,
		damage:     1,
		cost:       100,
		attackRate: 1,
	}
	tower.position = [2]int{x, y}
	tower.symbol = "T"
	tower.pathIndex = -1 // Towers don't follow a path
	return tower
}

func (tower *Tower) getRadius() int {
	return tower.radius
}

func (tower *Tower) getDamage() int {
	return tower.damage
}

func (tower *Tower) getCost() int {
	return tower.cost
}

// isEnemyInRange checks if an enemy is within tower's attack range
func (tower *Tower) isEnemyInRange(enemy *Enemy) bool {
	towerPos := tower.getPosition()
	enemyPos := enemy.getPosition()

	// Manhattan distance
	distance := abs(towerPos[0]-enemyPos[0]) + abs(towerPos[1]-enemyPos[1])
	return distance <= tower.radius
}

func (tower *Tower) attackEnemy(enemy *Enemy) {
	if tower.isEnemyInRange(enemy) {
		damage := tower.damage
		enemy.takeDamage(&damage)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
