package main

import "fmt"

type circle struct {
	x, y   float64
	radius float64
}

func (c *circle) Move(dx, dy float64) {
	c.x += dx
	c.y += dy
}

func main() {
	c := circle{x: 0, y: 0, radius: 4}
	c.Move(1, 1)
	fmt.Printf("%#v\n", c)
}
