// Structs for working with shapes as values.
//
// The main firefly module provides functions for drawing shapes.
// This modules provides useful struct for when you need to store
// or manipulate a shape before it can be drawn.
package shapes

import "github.com/life4/firefly-go/firefly"

type Shape interface {
	Draw()
}

// A wrapper for [firefly.DrawLine].
type Line struct {
	A     firefly.Point
	B     firefly.Point
	Style firefly.LineStyle
}

// Draw implements [Shape] interface.
func (s Line) Draw() {
	firefly.DrawLine(s.A, s.B, s.Style)
}

// A wrapper for [firefly.DrawRect].
type Rect struct {
	Point firefly.Point
	Size  firefly.Size
	Style firefly.Style
}

// Draw implements [Shape] interface.
func (s Rect) Draw() {
	firefly.DrawRect(s.Point, s.Size, s.Style)
}

// A wrapper for [firefly.DrawRoundedRect].
type RoundedRect struct {
	Point  firefly.Point
	Size   firefly.Size
	Corner firefly.Size
	Style  firefly.Style
}

// Draw implements [Shape] interface.
func (s RoundedRect) Draw() {
	firefly.DrawRoundedRect(s.Point, s.Size, s.Corner, s.Style)
}

// A wrapper for [firefly.DrawCircle].
type Circle struct {
	Point    firefly.Point
	Diameter int
	Style    firefly.Style
}

// Draw implements [Shape] interface.
func (s Circle) Draw() {
	firefly.DrawCircle(s.Point, s.Diameter, s.Style)
}

// A wrapper for [firefly.DrawEllipse].
type Ellipse struct {
	Point firefly.Point
	Size  firefly.Size
	Style firefly.Style
}

// Draw implements [Shape] interface.
func (s Ellipse) Draw() {
	firefly.DrawEllipse(s.Point, s.Size, s.Style)
}

// A wrapper for [firefly.DrawTriangle].
type Triangle struct {
	A     firefly.Point
	B     firefly.Point
	C     firefly.Point
	Style firefly.Style
}

// Draw implements [Shape] interface.
func (s Triangle) Draw() {
	firefly.DrawTriangle(s.A, s.B, s.C, s.Style)
}

// A wrapper for [firefly.DrawArc].
type Arc struct {
	Point    firefly.Point
	Diameter int
	Start    firefly.Angle
	Sweep    firefly.Angle
	Style    firefly.Style
}

// Draw implements [Shape] interface.
func (s Arc) Draw() {
	firefly.DrawArc(s.Point, s.Diameter, s.Start, s.Sweep, s.Style)
}

// A wrapper for [firefly.DrawSector].
type Sector struct {
	Point    firefly.Point
	Diameter int
	Start    firefly.Angle
	Sweep    firefly.Angle
	Style    firefly.Style
}

// Draw implements [Shape] interface.
func (s Sector) Draw() {
	firefly.DrawSector(s.Point, s.Diameter, s.Start, s.Sweep, s.Style)
}
