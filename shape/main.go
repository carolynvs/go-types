package main

import (
	"github.com/davecgh/go-spew/spew"
)

type shape interface {
	Move(dx, dy float64)
}

type point struct {
	x, y float64
}

func (p *point) Move(dx, dy float64) {
	p.x += dx
	p.y += dy
}

type line struct {
	start, stop *point
}

func (l *line) Move(dx, dy float64) {
	l.start.Move(dx, dy)
	l.stop.Move(dx, dy)
}

func main() {
	var s shape

	s = &point{}
	s.Move(1, 1)
	spew.Dump(s)

	s = &line{&point{0, 0}, &point{1, 1}}
	s.Move(2, 2)
	spew.Dump(s)
}
