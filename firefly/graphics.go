package firefly

import (
	"math"
	"unsafe"
)

const (
	// The screen width in pixels.
	Width = 240

	// The screen height in pixels.
	Height = 160
)

// A point on the screen.
//
// Typically, the upper-left corner of a bounding box of a shape.
type Point struct {
	X int
	Y int
}

// Convert the Point to a Size.
func (p Point) Size() Size {
	return Size{W: p.X, H: p.Y}
}

// Convert the Point to a Pad.
func (p Point) Pad() Pad {
	return Pad(p)
}

// Set X and Y to their absolute (non-negative) value.
func (p Point) Abs() Point {
	if p.X < 0 {
		p.X = -p.X
	}
	if p.Y < 0 {
		p.Y = -p.Y
	}
	return p
}

// Add together two points.
func (p Point) Add(r Point) Point {
	p.X += r.X
	p.Y += r.Y
	return p
}

// Subtract the given point from the current one.
func (p Point) Sub(r Point) Point {
	p.X -= r.X
	p.Y -= r.Y
	return p
}

// Set both X and Y to their minimum in the two given points.
func (p Point) ComponentMin(r Point) Point {
	if r.X < p.X {
		p.X = r.X
	}
	if r.Y < p.Y {
		p.Y = r.Y
	}
	return p
}

// Set both X and Y to their maximum in the two given points.
func (p Point) ComponentMax(r Point) Point {
	if r.X > p.X {
		p.X = r.X
	}
	if r.Y > p.Y {
		p.Y = r.Y
	}
	return p
}

// Size of a bounding box for a shape.
//
// The width and height must be positive.
type Size struct {
	// W is the width of the bounding box.
	W int
	// H is the height of the bounding box.
	H int
}

// Convert the Size to a Point.
func (s Size) Point() Point {
	return Point{X: s.W, Y: s.H}
}

// Convert the Size to a Pad.
func (s Size) Pad() Pad {
	return Pad{X: s.W, Y: s.H}
}

// Set W and H to their absolute (non-negative) value.
func (s Size) Abs() Size {
	if s.W < 0 {
		s.W = -s.W
	}
	if s.H < 0 {
		s.H = -s.H
	}
	return s
}

// Add two sizes.
func (s Size) Add(r Size) Size {
	s.W += r.W
	s.H += r.H
	return s
}

// Subtract the given size from the current one.
func (s Size) Sub(r Size) Size {
	s.W -= r.W
	s.H -= r.H
	return s
}

// Set both W and H to their minimum in the two given sizes.
func (s Size) ComponentMin(r Size) Size {
	if r.W < s.W {
		s.W = r.W
	}
	if r.H < s.H {
		s.H = r.H
	}
	return s
}

// Set both W and H to their maximum in the two given sizes.
func (s Size) ComponentMax(r Size) Size {
	if r.W > s.W {
		s.W = r.W
	}
	if r.H > s.H {
		s.H = r.H
	}
	return s
}

// An angle between two vectors.
//
// Used by [DrawArc] and [DrawSector].
// Constructed by [Dagrees] and [Radians].
type Angle struct {
	a float32
}

// Define an angle in radians where Tau (doubled Pi) is the full circle.
func Radians(a float32) Angle {
	return Angle{a}
}

// Define an angle in radians where 360.0 is the full circle.
func Degrees(a float32) Angle {
	return Angle{a * math.Pi / 180.0}
}

// Get the angle value in radians.
func (a Angle) Radians() float32 {
	return a.a
}

// Get the angle value in degrees.
func (a Angle) Degrees() float32 {
	return a.a / (math.Pi * 2)
}

// A pointer to a color in the color palette.
type Color uint8

const (
	// No color (100% transparency).
	ColorNone Color = 0

	// The first color in the palette. Typically, the darkest color.
	ColorDark Color = 1

	// The second color in the palette.
	ColorAccent Color = 2

	// The third color in the palette.
	ColorSecondary Color = 3

	// The last color in the palette. Typically, the brightest, almost white, color.
	ColorLight Color = 4
)

// The RGB value of a color in the palette.
type RGB struct {
	R uint8
	G uint8
	B uint8
}

// Style of a shape.
type Style struct {
	// The color to use to fill the shape.
	FillColor Color

	// The color to use for the shape stroke.
	StrokeColor Color

	// The width of the shape stroke.
	//
	// If zero, a solid shape without a stroke will be drawn.
	StrokeWidth int
}

// Convert the [Style] to a [LineStyle].
//
// [LineStyle] is the same as [Style] except it doesn't have a fill color.
func (s Style) LineStyle() LineStyle {
	return LineStyle{Color: s.StrokeColor, Width: s.StrokeWidth}
}

// The same as [Style] but without a fill color (only stroke color and width).
type LineStyle struct {
	Color Color
	Width int
}

// A mapping of colors in the image to the color palette.
type ImageColors struct {
	A Color
	B Color
	C Color
	D Color
}

// Fill the whole frame with the given color.
func ClearScreen(c Color) {
	clearScreen(int32(c))
}

// Set a color value in the palette.
func SetColor(c Color, v RGB) {
	setColor(int32(c), int32(v.R), int32(v.G), int32(v.B))
}

// Set the color palette.
func SetColors(a, b, c, d RGB) {
	setColors(
		int32(a.R), int32(a.G), int32(a.B),
		int32(b.R), int32(b.G), int32(b.B),
		int32(c.R), int32(c.G), int32(c.B),
		int32(d.R), int32(d.G), int32(d.B),
	)
}

// Set a single point (1 pixel is scaling is 1) on the frame.
func DrawPoint(p Point, c Color) {
	drawPoint(int32(p.X), int32(p.Y), int32(c))
}

// Draw a straight line from point a to point b.
func DrawLine(a, b Point, s LineStyle) {
	drawLine(
		int32(a.X), int32(a.Y),
		int32(b.X), int32(b.Y),
		int32(s.Color), int32(s.Width),
	)
}

// Draw a rectangle filling the given bounding box.
func DrawRect(p Point, b Size, s Style) {
	drawRect(
		int32(p.X), int32(p.Y),
		int32(b.W), int32(b.H),
		int32(s.FillColor), int32(s.StrokeColor), int32(s.StrokeWidth),
	)
}

// Draw a rectangle with rounded corners.
func DrawRoundedRect(p Point, b, c Size, s Style) {
	drawRoundedRect(
		int32(p.X), int32(p.Y),
		int32(b.W), int32(b.H),
		int32(c.W), int32(c.H),
		int32(s.FillColor), int32(s.StrokeColor), int32(s.StrokeWidth),
	)
}

// Draw a circle with the given diameter.
func DrawCircle(p Point, d uint32, s Style) {
	drawCircle(
		int32(p.X), int32(p.Y), int32(d),
		int32(s.FillColor), int32(s.StrokeColor), int32(s.StrokeWidth),
	)
}

// Draw an ellipse (oval).
func DrawEllipse(p Point, b Size, s Style) {
	drawEllipse(
		int32(p.X), int32(p.Y),
		int32(b.W), int32(b.H),
		int32(s.FillColor), int32(s.StrokeColor), int32(s.StrokeWidth),
	)
}

// Draw a triangle.
//
// The order of points doesn't matter.
func DrawTriangle(a, b, c Point, s Style) {
	drawTriangle(
		int32(a.X), int32(a.Y), int32(b.X), int32(b.Y), int32(c.X), int32(c.Y),
		int32(s.FillColor), int32(s.StrokeColor), int32(s.StrokeWidth),
	)
}

// Draw an arc.
func DrawArc(p Point, d uint32, start, sweep Angle, s Style) {
	drawArc(
		int32(p.X), int32(p.Y), int32(d),
		int32(start.a), int32(sweep.a),
		int32(s.FillColor), int32(s.StrokeColor), int32(s.StrokeWidth),
	)
}

// Draw a sector.
func DrawSector(p Point, d uint32, start, sweep Angle, s Style) {
	drawSector(
		int32(p.X), int32(p.Y), int32(d),
		int32(start.a), int32(sweep.a),
		int32(s.FillColor), int32(s.StrokeColor), int32(s.StrokeWidth),
	)
}

// Render text using the given font.
//
// Unlike in the other drawing functions, here [Point] points not to the top-left corner
// but to the baseline start position.
func DrawText(t string, f Font, p Point, c Color) {
	textPtr := unsafe.Pointer(unsafe.StringData(t))
	rawPtr := unsafe.Pointer(unsafe.SliceData(f.raw))
	drawText(
		textPtr, uint32(len(t)),
		rawPtr, uint32(len(f.raw)),
		int32(p.X), int32(p.Y), int32(c),
	)
}

// Render an image using the given colors.
func DrawImage(i Image, p Point, c ImageColors) {
	rawPtr := unsafe.Pointer(unsafe.SliceData(i.raw))
	drawImage(
		rawPtr, uint32(len(i.raw)),
		int32(p.X), int32(p.Y),
		int32(c.A), int32(c.B), int32(c.C), int32(c.D),
	)
}

// Draw a subregion of an image.
//
// Most often used to draw a sprite from a sprite atlas.
func DrawSubImage(i SubImage, p Point, c ImageColors) {
	rawPtr := unsafe.Pointer(unsafe.SliceData(i.raw))
	drawSubImage(
		rawPtr, uint32(len(i.raw)),
		int32(p.X), int32(p.Y),
		int32(i.point.X), int32(i.point.Y),
		uint32(i.size.W), uint32(i.size.H),
		int32(c.A), int32(c.B), int32(c.C), int32(c.D),
	)
}
