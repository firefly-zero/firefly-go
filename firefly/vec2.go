package firefly

import (
	"math"

	"github.com/orsinium-labs/tinymath"
)

const cmp_epsilon = 0.00001

// Vec2 is a utility class for dealing with float-based positions.
type Vec2 struct {
	X float32
	Y float32
}

// V is a shortcut for creating a [Vec2]
func V(x, y float32) Vec2 {
	return Vec2{X: x, Y: y}
}

// Convert a Vec2 to a [Point].
// The X and Y floats are truncated, meaning the floored value of positive numbers
// and the ceiling value of negative values.
func (v Vec2) Point() Point {
	return Point{X: int(v.X), Y: int(v.Y)}
}

// Round returns a new Vec2 with both X and Y rounded to the nearest integer.
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

// ComponentMin returns a Vec2 with both X and Y to their minimum in the two given Vec2s.
func (p Vec2) ComponentMin(r Vec2) Vec2 {
	if r.X < p.X {
		p.X = r.X
	}
	if r.Y < p.Y {
		p.Y = r.Y
	}
	return p
}

// ComponentMax returns a Vec2 with both X and Y to their maximum in the two given Vec2s.
func (p Vec2) ComponentMax(r Vec2) Vec2 {
	if r.X > p.X {
		p.X = r.X
	}
	if r.Y > p.Y {
		p.Y = r.Y
	}
	return p
}

// Check if the Vec2 is within the screen boundaries.
func (p Vec2) InBounds() bool {
	return p.X >= 0 && p.Y >= 0 && p.X < Width && p.Y < Height
}

// Scale returns a new Vec2 where both the X and Y value are individually
// multiplied by the scalar factor.
func (v Vec2) Scale(factor float32) Vec2 {
	return Vec2{X: v.X * factor, Y: v.Y * factor}
}

// Radius returns the vector length (aka magnitude).
func (v Vec2) Radius() float32 {
	return tinymath.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Radius returns the squared vector length (aka squared magnitude),
// which is simpler to calculate.
func (v Vec2) RadiusSquared() float32 {
	return v.X*v.X + v.Y*v.Y
}

// Azimuth returns the angle of the [polar coordinate] of the vector.
//
//   - [V](1, 0).Azimuth() == [Degrees](0)
//   - [V](0, 1).Azimuth() == [Degrees](90)
//   - [V](-1, 0).Azimuth() == [Degrees](180)
//   - [V](0, -1).Azimuth() == [Degrees](270)
func (v Vec2) Azimuth() Angle {
	r := math.Pi / 2. * tinymath.Atan2Norm(v.Y, v.X)
	return Radians(r)
}

// MoveTowards returns a vector that has moved towards "to" by the "delta"
// amount, but will not go past "to".
// Use negative "delta" value to move away.
func (v Vec2) MoveTowards(to Vec2, delta float32) Vec2 {
	vd := to.Sub(v)
	dist := vd.Radius()
	if dist <= delta || dist < cmp_epsilon {
		return to
	} else {
		return v.Add(vd.Scale(delta / dist))
	}
}
