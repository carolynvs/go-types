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

func printType(s shape) {
	switch v := s.(type) {
	case *point:
		fmt.Printf("I'm a point! %s\n", v)
	case *line:
		fmt.Printf("I'm a line! %s\n", v)
	}
}

func main() {
	printType(&point{1, 1})
	printType(&line{&point{0, 0}, &point{1, 1}})
}
