package main

import (
	"fmt"
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

func (p *point) String() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

type line struct {
	start, stop *point
}

func (l *line) Move(dx, dy float64) {
	l.start.Move(dx, dy)
	l.stop.Move(dx, dy)
}

func (l *line) String() string {
	return fmt.Sprintf("%s - %s", l.start, l.stop)
}

func main() {
	var s shape
	s = &point{}

	p, ok := s.(*point)
	if ok {
		fmt.Printf("Yay! It's a point!\n%s\n", p)
	}

	_, ok = s.(*line)
	if !ok {
		fmt.Println("It is NOT a line")
	}
}
