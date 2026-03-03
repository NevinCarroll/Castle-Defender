package main

type Entity struct {
	position [2]int
	symbol   string
}

func getPosition(entity *Entity) [2]int {
	return entity.position
}

func setPosition(entity *Entity, x *int, y *int) {
	entity.position[0] = *x
	entity.position[1] = *y
}
