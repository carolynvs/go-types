package main

import (
	"github.com/davecgh/go-spew/spew"
)

type point struct {
	x, y float64
}

type circle struct {
	*point
	radius float64
}

func (c *circle) Move(dx, dy float64) {
	c.x += dx
	c.y += dy
}

func main() {
	c := circle{&point{0, 0}, 4}
	c.Move(1, 1)
	spew.Dump(c)
}
