package main

import (
	"math"
)

type V2 struct {
	X, Y float64
}

func V2Add(a, b V2) V2 {
	return V2{a.X + b.X, a.Y + b.Y}
}

func V2Scale(a V2, b float64) V2 {
	return V2{a.X * b, a.Y * b}
}

func V2LengthSquared(a V2) float64 {
	return a.X*a.X + a.Y*a.Y
}

func V2Length(a V2) float64 {
	return math.Sqrt(V2LengthSquared(a))
}

func V2Normalize(a V2) V2 {
	len := V2Length(a)
	return V2Scale(a, 1/len)
}

func V2Lerp(a, b V2, t float64) V2 {
	return V2{mix(a.X, b.X, t), mix(a.Y, b.Y, t)}
}
