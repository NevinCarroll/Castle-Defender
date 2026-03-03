package main

type Track struct {
	layout [][]int
}

func (track *Track) getLayout() [][]int {
	return track.layout
}
