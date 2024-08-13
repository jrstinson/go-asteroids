package util

import (
	"math"
)

type Vector struct {
	X float64
	Y float64
}

func (v *Vector) Add(v2 Vector) Vector {
	return Vector{X: v.X + v2.X, Y: v.Y + v2.Y}
}

func (v *Vector) Sub(v2 Vector) Vector {
	return Vector{X: v.X - v2.X, Y: v.Y - v2.Y}
}

func (v *Vector) Mult(s float64) Vector {
	return Vector{X: v.X * s, Y: v.Y * s}
}

func (v *Vector) Div(s float64) Vector {
	return Vector{X: v.X / s, Y: v.Y / s}
}

func (v *Vector) Rotate(angle float64) Vector {
	return Vector{X: v.X*math.Cos(angle) - v.Y*math.Sin(angle), Y: v.X*math.Sin(angle) + v.Y*math.Cos(angle)}
}

func (v *Vector) RotateAround(angle float64, origin Vector) Vector {
	r_v := v.Sub(origin)
	r_v = r_v.Rotate(angle)
	return r_v.Add(origin)
}
