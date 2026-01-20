package firefly

import (
	"math"

	"github.com/orsinium-labs/tinymath"
)

const epsilon = 0.00001

// Utility type for dealing with float-based positions.
type Vec2 struct {
	X float32
	Y float32
}

// Shortcut for creating a [Vec2].
func V(x, y float32) Vec2 {
	return Vec2{X: x, Y: y}
}

// Convert a Vec2 to a [Point].
//
// The X and Y floats are truncated, meaning the floored value of positive numbers
// and the ceiling value of negative values.
func (v Vec2) Point() Point {
	return Point{X: int(v.X), Y: int(v.Y)}
}

// Get Vec2 with both X and Y rounded to the nearest integer.
func (v Vec2) Round() Vec2 {
	return Vec2{X: tinymath.Round(v.X), Y: tinymath.Round(v.Y)}
}

func (v Vec2) Abs() Vec2 {
	return Vec2{X: tinymath.Abs(v.X), Y: tinymath.Abs(v.Y)}
}

func (v Vec2) Add(rhs Vec2) Vec2 {
	return Vec2{X: v.X + rhs.X, Y: v.Y + rhs.Y}
}

func (v Vec2) Sub(rhs Vec2) Vec2 {
	return Vec2{X: v.X - rhs.X, Y: v.Y - rhs.Y}
}

func (v Vec2) Negate() Vec2 {
	return Vec2{X: -v.X, Y: -v.Y}
}

// Get Vec2 with both X and Y to their minimum in the two given Vec2s.
func (v Vec2) ComponentMin(r Vec2) Vec2 {
	if r.X < v.X {
		v.X = r.X
	}
	if r.Y < v.Y {
		v.Y = r.Y
	}
	return v
}

// Get Vec2 with both X and Y to their maximum in the two given Vec2s.
func (v Vec2) ComponentMax(r Vec2) Vec2 {
	if r.X > v.X {
		v.X = r.X
	}
	if r.Y > v.Y {
		v.Y = r.Y
	}
	return v
}

// Check if the Vec2 is within the screen boundaries.
func (v Vec2) InBounds() bool {
	return v.X >= 0 && v.Y >= 0 && v.X < Width && v.Y < Height
}

// Get Vec2 where both the X and Y value are individually multiplied by the scalar factor.
func (v Vec2) Scale(factor float32) Vec2 {
	return Vec2{X: v.X * factor, Y: v.Y * factor}
}

// Radius returns the vector length (aka magnitude).
func (v Vec2) Radius() float32 {
	return tinymath.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Get the squared vector length (aka squared magnitude), which is simpler to calculate.
func (v Vec2) RadiusSquared() float32 {
	return v.X*v.X + v.Y*v.Y
}

// The angle of the polar coordinate of the vector.
//
//   - [V](1, 0).Azimuth() == [Degrees](0)
//   - [V](0, 1).Azimuth() == [Degrees](90)
//   - [V](-1, 0).Azimuth() == [Degrees](180)
//   - [V](0, -1).Azimuth() == [Degrees](270)
func (v Vec2) Azimuth() Angle {
	r := math.Pi / 2. * tinymath.Atan2Norm(v.Y, v.X)
	return Radians(r)
}

// Get Vector that has moved towards "to" by the "delta" amount, but will not go past "to".
//
// Use negative "delta" value to move away.
func (v Vec2) MoveTowards(to Vec2, delta float32) Vec2 {
	vd := to.Sub(v)
	dist := vd.Radius()
	if dist <= delta || dist < epsilon {
		return to
	}
	return v.Add(vd.Scale(delta / dist))
}
