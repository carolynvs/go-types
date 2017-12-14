package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y float64
}

func (p point) Subtract(o point) float64 {
	dx := p.x - o.x
	dy := p.y - o.y
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}

func (p *point) Move(dx, dy float64) {
	p.x += dx
	p.y += dy
}

func main() {
	p1 := point{0, 0}

	p1.Move(1, 1)
	fmt.Printf("%+v\n", p1)
}
