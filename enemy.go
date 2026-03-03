package main

type Enemy struct {
	Entity
	health int
}

func (enemy *Enemy) getHealth() int {
	return enemy.health
}

func (enemy *Enemy) takeDamage(damage *int) {
	enemy.health -= *damage
}