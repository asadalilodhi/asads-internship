package main

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Triangle struct {
	Width  float64
	Height float64
	Length float64
}

func (t Triangle) Area() float64 {
	a := float64(t.Height)
	b := float64(t.Length)
	c := float64(t.Width)
	s := (a + b + c) / 2
	area := math.Sqrt(s * (s - a) * (s - b) * (s - c))
	return math.Round(area*100) / 100
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(triangle Triangle) float64 {
	a := float64(triangle.Height)
	b := float64(triangle.Length)
	c := float64(triangle.Width)
	return (a + b + c)
}
