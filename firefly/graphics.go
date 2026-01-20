package firefly

import (
	"cmp"
	"math"
	"unsafe"

	"github.com/orsinium-labs/tinymath"
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

// Shortcut for creating a [Point].
func P(x, y int) Point {
	return Point{X: x, Y: y}
}

// Render a single pixel.
func (p Point) Draw(c Color) {
	DrawPoint(p, c)
}

// Convert the Point to a [Size].
func (p Point) Size() Size {
	return Size{W: p.X, H: p.Y}
}

// Convert the Point to a [Pad].
func (p Point) Pad() Pad {
	return Pad(p)
}

// Convert the Point to a [Vec2].
func (p Point) Vec2() Vec2 {
	return Vec2{X: float32(p.X), Y: float32(p.Y)}
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

// Check if the point is within the screen boundaries.
func (p Point) InBounds() bool {
	return p.X >= 0 && p.Y >= 0 && p.X < Width && p.Y < Height
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

// Shortcut for creating a [Size].
func S(w, h int) Size {
	return Size{W: w, H: h}
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
	return 180 * a.a / math.Pi
}

// Convert an Angle to a [Vec2].
func (a Angle) Vec2() Vec2 {
	return Vec2{
		X: tinymath.Cos(a.a),
		Y: -tinymath.Sin(a.a),
	}
}

func (a Angle) Neg() Angle {
	a.a = -a.a
	return a
}

func (a Angle) Add(r Angle) Angle {
	a.a += r.a
	return a
}

func (a Angle) Sub(r Angle) Angle {
	a.a -= r.a
	return a
}

// Ensure the angle is on the 0°-360° range.
func (a Angle) Normalize() Angle {
	for a.a >= math.Pi*2 {
		a.a -= math.Pi * 2
	}
	for a.a < 0 {
		a.a += math.Pi * 2
	}
	return a
}

// Angle difference to go from "a" to "to".
//
// Result will be in the range of [-[math.Pi], +[math.Pi]].
// When "a" and "to" are opposite,
// returns -[math.Pi] if "a" is smaller than "to", or [math.Pi] otherwise.
//
// Input angles do not need to be normalized.
// Based on the Godot [angle_difference] (licensed under MIT)
//
// [angle_difference]: https://github.com/godotengine/godot/blob/4.5.1-stable/core/math/math_funcs.h#L482-L489
func (a Angle) Difference(to Angle) Angle {
	// tinymath has "RemEuclid" https://github.com/orsinium-labs/tinymath/blob/v1.1.0/tinymath.go#L324-L332
	// but it has some math bugs, so we have to resort to the big math functions.
	diff := math.Mod(float64(to.Radians()-a.Radians()), 2*math.Pi)
	return Radians(float32(math.Mod(2*diff, 2*math.Pi) - diff))
}

// Rotates "a" toward "to" by the "delta" amount.
//
// Will not go past "to", but interpolated correctly when the angles
// wrap around [Radians](2*[math.Pi]) or [Degrees](360).
//
// If "delta" is negative, this function will rotate away from "to",
// towards the opposite angle, and will not go past the opposite angle.
//
// Based on the Godot [rotate_towards] (licensed under MIT)
//
// [rotate_towards]: https://github.com/godotengine/godot/blob/4.5.1-stable/core/math/math_funcs.h#L598-L609
func (a Angle) RotateTowards(to, delta Angle) Angle {
	diff := a.Difference(to).Radians()
	absDiff := tinymath.Abs(diff)
	return Radians(
		a.Radians() + clamp(delta.Radians(), absDiff-math.Pi, absDiff)*tinymath.Sign(diff),
	)
}

func clamp[T cmp.Ordered](val, minimum, maximum T) T {
	switch {
	case val < minimum:
		return minimum
	case val > maximum:
		return maximum
	default:
		return val
	}
}

// A pointer to a color in the color palette.
type Color uint8

const (
	// No color (100% transparency). Default.
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

// Get the color name in the default palette (SWEETIE-16).
//
// Implements [fmt.Stringer].
func (color Color) String() string {
	switch color {
	case ColorBlack:
		return "Black"
	case ColorBlue:
		return "Blue"
	case ColorCyan:
		return "Cyan"
	case ColorDarkBlue:
		return "DarkBlue"
	case ColorDarkGray:
		return "DarkGray"
	case ColorDarkGreen:
		return "DarkGreen"
	case ColorGray:
		return "Gray"
	case ColorGreen:
		return "Green"
	case ColorLightBlue:
		return "LightBlue"
	case ColorLightGray:
		return "LightGray"
	case ColorLightGreen:
		return "LightGreen"
	case ColorNone:
		return "None"
	case ColorOrange:
		return "Orange"
	case ColorPurple:
		return "Purple"
	case ColorRed:
		return "Red"
	case ColorWhite:
		return "White"
	case ColorYellow:
		return "Yellow"
	default:
		return "???"
	}
}

// The RGB value of a color in the palette.
type RGB struct {
	// Red component
	R uint8
	// Green component
	G uint8
	// Blue component
	B uint8
}

func NewRGB(r, g, b uint8) RGB {
	return RGB{R: r, G: g, B: b}
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

// Make [Style] for a solid shape (without stroke).
func Solid(c Color) Style {
	return Style{FillColor: c}
}

// Make [Style] for an outlined shape (without fill).
func Outlined(c Color, width int) Style {
	return Style{StrokeColor: c, StrokeWidth: width}
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

// A shortcut for creating a new [LineStyle].
func L(c Color, w int) LineStyle {
	return LineStyle{Color: c, Width: w}
}

// Draw a line from a to b.
func (s LineStyle) Draw(a, b Point) {
	DrawLine(a, b, s)
}

// A loaded font file.
//
// Can be loaded using [LoadFile].
type Font struct {
	raw []byte
}

// Render the given text.
func (f Font) Draw(t string, p Point, c Color) {
	DrawText(t, f, p, c)
}

// If the font is for ASCII encoding (English alphabet).
func (f Font) IsASCII() bool {
	return f.raw[1] == 0
}

// Calculate width (in pixels) of the given text.
//
// This function does not account for newlines.
func (f Font) LineWidth(t string) int {
	return len(t) * f.CharWidth()
}

// Character width.
func (f Font) CharWidth() int {
	return int(f.raw[2])
}

// Character height.
func (f Font) CharHeight() int {
	return int(f.raw[3])
}

// A loaded image file.
//
// Can be loaded using [LoadFile].
type Image struct {
	raw []byte
}

// Render the image.
func (i Image) Draw(p Point) {
	DrawImage(i, p)
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

// Get color of a pixel in the image.
//
// Returns [ColorNone] if out of bounds.
func (i Image) GetPixel(point Point) Color {
	if point.X < 0 || point.Y < 0 {
		return ColorNone
	}
	size := i.Size()
	if point.X >= size.W || point.Y >= size.H {
		return ColorNone
	}
	bpp := i.raw[1]
	headerLen := 5 + (1 << (bpp - 1))
	body := i.raw[headerLen:]

	pixelIndex := point.X + point.Y*size.W
	bodyIndex := pixelIndex * int(bpp) / 8
	pixelValue := body[bodyIndex]

	switch bpp {
	case 1:
		byteOffset := 1 * (7 - pixelIndex%8)
		pixelValue = (pixelValue >> byte(byteOffset)) & 0b1
	case 2:
		byteOffset := 2 * (3 - pixelIndex%4)
		pixelValue = (pixelValue >> byte(byteOffset)) & 0b11
	case 4:
		byteOffset := 4 * (1 - pixelIndex%2)
		pixelValue = (pixelValue >> byte(byteOffset)) & 0b1111
	default:
		panic("invalid bpp")
	}

	return i.GetColor(pixelValue)
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
func (i Image) ReplaceColor(oldC, newC Color) {
	var p uint8
	for p = range 16 {
		if i.GetColor(p) == oldC {
			i.SetColor(p, newC)
		}
	}
}

// A subregion of an image. Constructed using [Image.Sub].
type SubImage struct {
	raw   []byte
	point Point
	size  Size
}

// Render the sub image at the given point.
func (i SubImage) Draw(p Point) {
	DrawSubImage(i, p)
}

// Image returns back the original parent [Image] from which this sub-image
// was created from.
func (i SubImage) Image() Image {
	return Image{raw: i.raw}
}

// Point returns the offset of this sub-image in the parent [Image].
func (i SubImage) Point() Point {
	return i.point
}

// Size returns the size of this sub-image.
func (i SubImage) Size() Size {
	return i.size
}

// Width returns the width of this sub-image.
func (i SubImage) Width() int {
	return i.size.W
}

// Height returns the height of this sub-image.
func (i SubImage) Height() int {
	return i.size.H
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

// Set this canvas as the target for all subsequent draw operations.
func (c Canvas) Set() {
	SetCanvas(c)
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

// Set all colors in the color palette.
func SetPalette(colors [16]RGB) {
	for c, v := range colors {
		setColor(int32(c), int32(v.R), int32(v.G), int32(v.B))
	}
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

// Render a QR code for the given text.
func DrawQR(t string, p Point, black, white Color) {
	ptr := unsafe.Pointer(unsafe.StringData(t))
	drawQR(
		ptr, uint32(len(t)),
		int32(p.X), int32(p.Y),
		int32(black), int32(white),
	)
}

// Render an image at the given point.
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

// Set the target image for all subsequent drawing operations.
func SetCanvas(c Canvas) {
	rawPtr := unsafe.Pointer(unsafe.SliceData(c.raw))
	setCanvas(rawPtr, uint32(len(c.raw)))
}

// Make all subsequent drawing operations target the screen instead of a canvas.
//
// Cancels the effect of [SetCanvas].
func UnsetCanvas() {
	unsetCanvas()
}
