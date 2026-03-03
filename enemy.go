package main

type Enemy struct {
	Entity
	health int
}

func takeDamage(enemy *Enemy, damage *int) {
	enemy.health -= *damage
}