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

func main() {
	p1 := point{0, 0}
	p2 := point{0, 4}
	d := p2.Subtract(p1)
	fmt.Printf("distance: %v\n", d)
}
