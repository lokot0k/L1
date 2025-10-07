package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p *Point) GetX() float64  { return p.x }
func (p *Point) SetX(x float64) { p.x = x }
func (p *Point) GetY() float64  { return p.y }
func (p *Point) SetY(y float64) { p.y = y }

func (p *Point) Distance(other Point) float64 {
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2))
}

func main() {
	point := NewPoint(1, 1)
	otherPoint := NewPoint(10, 10)

	fmt.Printf(
		"Distance between %v and %v: %.5f\n",
		point,
		otherPoint,
		point.Distance(otherPoint),
	)

	otherPoint.SetX(4.0)
	otherPoint.SetY(5.0)

	fmt.Printf(
		"Distance between %v and %v: %.5f\n",
		point,
		otherPoint,
		point.Distance(otherPoint),
	)
}
