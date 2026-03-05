package main

type Entity struct {
	position    [2]int // Current position on the track
	symbol      string
	pathIndex   int // Current position in the path (0 = start)
	moveCounter int // Counter to control movement speed
}

func (entity *Entity) getPosition() [2]int {
	return entity.position
}

func (entity *Entity) setPosition(x *int, y *int) {
	entity.position[0] = *x
	entity.position[1] = *y
}

func (entity *Entity) getPathIndex() int {
	return entity.pathIndex
}

func (entity *Entity) setPathIndex(index int) {
	entity.pathIndex = index
}

func (entity *Entity) getSymbol() string {
	return entity.symbol
}

func (entity *Entity) setSymbol(symbol string) {
	entity.symbol = symbol
}
