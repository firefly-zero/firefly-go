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
	// Black color: #1A1C2C.
	ColorBlack Color = 1
	// Purple color: #5D275D.
	ColorPurple Color = 2
	// Red color: #B13E53.
	ColorRed Color = 3
	// Orange color: #EF7D57.
	ColorOrange Color = 4
	// Yellow color: #FFCD75.
	ColorYellow Color = 5
	// Light green color: #A7F070.
	ColorLightGreen Color = 6
	// Green color: #38B764.
	ColorGreen Color = 7
	// Dark green color: #257179.
	ColorDarkGreen Color = 8
	// Dark blue color: #29366F.
	ColorDarkBlue Color = 9
	// Blue color: #3B5DC9.
	ColorBlue Color = 10
	// Light blue color: #41A6F6.
	ColorLightBlue Color = 11
	// Cyan color: #73EFF7.
	ColorCyan Color = 12
	// White color: #F4F4F4.
	ColorWhite Color = 13
	// Light gray color: #94B0C2.
	ColorLightGray Color = 14
	// Gray color: #566C86.
	ColorGray Color = 15
	// Dark gray color: #333C57.
	ColorDarkGray Color = 16
)

// The RGB value of a color in the palette.
type RGB struct {
	// Red component
	R uint8
	// Green component
	G uint8
	// Blue component
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

// A loaded font file.
//
// Can be loaded using [LoadFile].
type Font struct {
	raw []byte
}

// A loaded image file.
//
// Can be loaded using [LoadFile].
type Image struct {
	raw []byte
}

// Get a rectangle subregion of the image.
func (i Image) Sub(p Point, s Size) SubImage {
	return SubImage{raw: i.raw, point: p, size: s}
}

// Bits per pixel. One of: 1, 2, or 4.
func (i Image) BPP() uint8 {
	return i.raw[1]
}

// The color used for transparency. If no transparency, returns [ColorNone].
func (i Image) Transparency() Color {
	c := i.raw[4]
	if c > 15 {
		return ColorNone
	}
	return Color(c + 1)
}

// Set the color that should represent transparency.
//
// Pass ColorNone to disable transparency.
func (i Image) SetTransparency(c Color) {
	if c == ColorNone {
		i.raw[4] = 16
	}
	i.raw[4] = byte(c) - 1
}

// The number of pixels the image has.
func (i Image) Pixels() int {
	return len(i.raw) * 8 / int(i.BPP())
}

// The image width in pixels.
func (i Image) Width() int {
	return int(i.raw[2]) | int(i.raw[3])<<8
}

// The image height in pixels.
func (i Image) Height() int {
	w := i.Width()
	if w == 0 {
		return 0
	}
	return i.Pixels() / w
}

// The image size in pixels.
func (i Image) Size() Size {
	w := i.Width()
	if w == 0 {
		return Size{}
	}
	return Size{
		W: w,
		H: i.Pixels() / w,
	}
}

// Get the color used to represent the given pixel value.
func (i Image) GetColor(p uint8) Color {
	if p > 15 {
		return ColorNone
	}
	byteVal := i.raw[5+p/2]
	if p%2 == 0 {
		byteVal >>= 4
	}
	byteVal &= 0b1111
	transp := i.raw[4]
	if byteVal == transp {
		return ColorNone
	}
	return Color(byteVal + 1)
}

// Set color to be used to represent the given pixel value.
func (i Image) SetColor(p uint8, c Color) {
	if p > 15 || c == ColorNone {
		return
	}
	byteIdx := 5 + p/2
	byteVal := i.raw[byteIdx]
	colorVal := byte(c) - 1
	if p%2 == 0 {
		byteVal = (colorVal << 4) | (byteVal & 0b_0000_1111)
	} else {
		byteVal = (byteVal & 0b_1111_0000) | colorVal
	}
	i.raw[byteIdx] = byteVal
}

// Replace the old color with the new value.
func (i Image) ReplaceColor(old, new Color) {
	var p uint8
	for p = range 16 {
		if i.GetColor(p) == old {
			i.SetColor(p, new)
		}
	}
}

// A subregion of an image. Constructed using [Image.Sub].
type SubImage struct {
	raw   []byte
	point Point
	size  Size
}

// Canvas is an [Image] that can be drawn upon.
//
// Constructed by [NewCanvas].
type Canvas struct {
	raw []byte
}

func NewCanvas(s Size) Canvas {
	const headerSize = 5 + 8
	bodySize := s.W * s.H / 2
	raw := make([]byte, headerSize+bodySize)
	raw[0] = 0x21           // magic number
	raw[1] = 4              // BPP
	raw[2] = byte(s.W)      // width
	raw[3] = byte(s.W >> 8) // width
	raw[4] = 255            // transparency

	// color swaps
	var i byte
	for i = range 8 {
		raw[5+i] = ((i * 2) << 4) | (i*2 + 1)
	}
	return Canvas{raw}
}

// Represent the canvas as an [Image].
func (c Canvas) Image() Image {
	return Image(c)
}

// Fill the whole frame with the given color.
func ClearScreen(c Color) {
	clearScreen(int32(c))
}

// Set a color value in the palette.
func SetColor(c Color, v RGB) {
	setColor(int32(c), int32(v.R), int32(v.G), int32(v.B))
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
func DrawCircle(p Point, d int, s Style) {
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
func DrawArc(p Point, d int, start, sweep Angle, s Style) {
	drawArc(
		int32(p.X), int32(p.Y), int32(d),
		start.a, sweep.a,
		int32(s.FillColor), int32(s.StrokeColor), int32(s.StrokeWidth),
	)
}

// Draw a sector.
func DrawSector(p Point, d int, start, sweep Angle, s Style) {
	drawSector(
		int32(p.X), int32(p.Y), int32(d),
		start.a, sweep.a,
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
func DrawImage(i Image, p Point) {
	rawPtr := unsafe.Pointer(unsafe.SliceData(i.raw))
	drawImage(
		rawPtr, uint32(len(i.raw)),
		int32(p.X), int32(p.Y),
	)
}

// Draw a subregion of an image.
//
// Most often used to draw a sprite from a sprite atlas.
func DrawSubImage(i SubImage, p Point) {
	rawPtr := unsafe.Pointer(unsafe.SliceData(i.raw))
	drawSubImage(
		rawPtr, uint32(len(i.raw)),
		int32(p.X), int32(p.Y),
		int32(i.point.X), int32(i.point.Y),
		uint32(i.size.W), uint32(i.size.H),
	)
}

func SetCanvas(c Canvas) {
	rawPtr := unsafe.Pointer(unsafe.SliceData(c.raw))
	setCanvas(rawPtr, uint32(len(c.raw)))
}

func UnsetCanvas() {
	unsetCanvas()
}
