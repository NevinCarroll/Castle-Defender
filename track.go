package main

type Track struct {
	layout [][]string
}

func (track *Track) getLayout() [][]string {
	return track.layout
}
