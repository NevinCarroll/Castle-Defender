package main

type Entity struct {
	position [2]int
	symbol   string
}

func (entity *Entity) getPosition() [2]int {
	return entity.position
}

func (entity *Entity) setPosition(x *int, y *int) {
	entity.position[0] = *x
	entity.position[1] = *y
}
