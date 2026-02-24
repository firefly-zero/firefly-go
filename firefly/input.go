package firefly

import (
	"math"

	"github.com/orsinium-labs/tinymath"
)

const (
	// The lowest possible value for [Pad.X].
	PadMinX = -1000

	// The lowest possible value for [Pad.Y].
	PadMinY = -1000

	// The highest possible value for [Pad.X].
	PadMaxX = 1000

	// The highest possible value for [Pad.Y].
	PadMaxY = 1000
)

// The minimum X or Y value when converting Pad into DPad8
// for the direction to be considered pressed.
const dPad8Threshold = 300

const dPad4Threshold = 300

// A finger position on the touch pad.
//
// Both X and Y are somewhere the range between -1000 and 1000 (both ends included).
// The 1000 X is on the right, the 1000 Y is on the top.
type Pad struct {
	X int
	Y int
}

// The distance from the pad center to the touch point.
func (p Pad) Radius() float32 {
	r := p.X*p.X + p.Y*p.Y
	return tinymath.Sqrt(float32(r))
}

// Radius returns the squared distance from the pad center to the touch point,
// which is simpler to calculate, and guarantees that the return value is an integer.
func (p Pad) RadiusSquared() int {
	return p.X*p.X + p.Y*p.Y
}

// The angle of the [polar coordinate] of the touch point.
//
//   - (Pad{X: 1, Y: 0}).Azimuth() == [Degrees](0)
//   - (Pad{X: 0, Y: 1}).Azimuth() == [Degrees](90)
//   - (Pad{X: -1, Y: 0}).Azimuth() == [Degrees](180)
//   - (Pad{X: 0, Y: -1}).Azimuth() == [Degrees](270)
//
// [polar coordinate]: https://en.wikipedia.org/wiki/Polar_coordinate_system
func (p Pad) Azimuth() Angle {
	r := math.Pi / 2. * tinymath.Atan2Norm(float32(p.Y), float32(p.X))
	return Radians(r)
}

// Convert the [Pad] into [Point].
func (p Pad) Point() Point {
	return Point(p)
}

// Convert the [Pad] into [Size].
func (p Pad) Size() Size {
	return Size{W: p.X, H: p.Y}
}

// Convert the [Pad] into [DPad4].
func (p Pad) DPad4() DPad4 {
	x := p.X
	y := p.Y
	absX := abs(x)
	absY := abs(y)
	switch {
	case y > dPad4Threshold && y > absX:
		return DPad4Up
	case y < -dPad4Threshold && -y > absX:
		return DPad4Down
	case x > dPad4Threshold && x > absY:
		return DPad4Right
	case x < -dPad4Threshold && -x > absY:
		return DPad4Left
	default:
		return DPad4None
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Convert the [Pad] into [DPad8].
func (p Pad) DPad8() DPad8 {
	return DPad8{
		Left:  p.X <= -dPad8Threshold,
		Right: p.X >= dPad8Threshold,
		Up:    p.Y <= -dPad8Threshold,
		Down:  p.Y >= dPad8Threshold,
	}
}

// 4-directional DPad-like representation of the [Pad].
//
// Constructed with [Pad.DPad4]. Useful for simple games and ports.
// The middle of the pad is a "dead zone" pressing which will not activate any direction.
//
// Implements all the same methods as [DPad8].
type DPad4 uint8

// Possible directions for [DPad4].
const (
	DPad4None  DPad4 = 0
	DPad4Right DPad4 = 1
	DPad4Up    DPad4 = 2
	DPad4Left  DPad4 = 3
	DPad4Down  DPad4 = 4
)

func (p DPad4) Any() bool {
	return p != DPad4None
}

func (p DPad4) JustPressed(old DPad4) DPad4 {
	if p == old {
		return DPad4None
	}
	return p
}

func (p DPad4) JustReleased(old DPad4) DPad4 {
	if p == old {
		return DPad4None
	}
	return old
}

func (p DPad4) Held(old DPad4) DPad4 {
	if p == old {
		return p
	}
	return DPad4None
}

// 8-directional DPad-like representation of the [Pad].
//
// Constructed with [Pad.DPad8]. Useful for simple games and ports.
// The middle of the pad is a "dead zone" pressing which will not activate any direction.
//
// Invariant: it's not possible for opposite directions (left and right, or down and up)
// to be active at the same time. However, it's possible for neighboring directions
// (like up and right) to be active at the same time if the player presses a diagonal.
//
// Implements all the same methods as [DPad4].
type DPad8 struct {
	Left  bool
	Right bool
	Up    bool
	Down  bool
}

func (p DPad8) Any() bool {
	return p.Left || p.Right || p.Up || p.Down
}

// Given the old state, get directions that were not pressed but are pressed now.
func (p DPad8) JustPressed(old DPad8) DPad8 {
	p.Left = p.Left && !old.Left
	p.Right = p.Right && !old.Right
	p.Up = p.Up && !old.Up
	p.Down = p.Down && !old.Down
	return p
}

// Given the old state, get directions that were pressed but aren't pressed now.
func (p DPad8) JustReleased(old DPad8) DPad8 {
	p.Left = !p.Left && old.Left
	p.Right = !p.Right && old.Right
	p.Up = !p.Up && old.Up
	p.Down = !p.Down && old.Down
	return p
}

// Given the old state, get directions that were pressed and are still pressed now.
func (p DPad8) Held(old DPad8) DPad8 {
	p.Left = p.Left && old.Left
	p.Right = p.Right && old.Right
	p.Up = p.Up && old.Up
	p.Down = p.Down && old.Down
	return p
}

// State of the buttons.
type Buttons struct {
	// South. The bottom button, like A on the X-Box controller.
	//
	// Typically used for confirmation, main action, jump, etc.
	S bool

	// East. The right button, like B on the X-Box controller.
	//
	// Typically used for cancellation, going to previous screen, etc.
	E bool

	// West. The left button, like X on the X-Box controller.
	//
	// Typically used for attack.
	W bool

	// North. The top button, like Y on the X-Box controller.
	//
	// Typically used for a secondary action, like charged attack.
	N bool

	// The menu button, almost always handled by the runtime.
	Menu bool
}

// Given the old state, get buttons that were not pressed but are pressed now.
func (p Buttons) JustPressed(old Buttons) Buttons {
	p.S = p.S && !old.S
	p.E = p.E && !old.E
	p.W = p.W && !old.W
	p.N = p.N && !old.N
	p.Menu = p.Menu && !old.N
	return p
}

// Given the old state, get buttons that were pressed but aren't pressed now.
func (p Buttons) JustReleased(old Buttons) Buttons {
	p.S = !p.S && old.S
	p.E = !p.E && old.E
	p.W = !p.W && old.W
	p.N = !p.N && old.N
	p.N = !p.N && old.N
	p.N = !p.N && old.N
	p.Menu = !p.Menu && old.Menu
	return p
}

// Given the old state, get buttons that were pressed and are still pressed now.
func (p Buttons) Held(old Buttons) Buttons {
	p.S = p.S && old.S
	p.E = p.E && old.E
	p.W = p.W && old.W
	p.N = p.N && old.N
	p.Menu = p.Menu && old.Menu
	return p
}

// Check if any button is currently pressed.
func (p Buttons) Any() bool {
	return p.S || p.E || p.W || p.N || p.Menu
}

// Get the current touch pad state.
//
// The peer can be [Combined] or one of the [GetPeers].
func ReadPad(p Peer) (Pad, bool) {
	raw := readPad(uint32(p.raw))
	pressed := raw != 0xffff
	if !pressed {
		return Pad{}, false
	}
	pad := Pad{
		X: int(int16(raw >> 16)),
		Y: int(int16(raw)),
	}
	return pad, true
}

// Get the currently pressed buttons.
//
// The peer can be [Combined] or one of the [GetPeers].
func ReadButtons(p Peer) Buttons {
	raw := readButtons(uint32(p.raw))
	return Buttons{
		S:    hasBitSet(raw, 0),
		E:    hasBitSet(raw, 1),
		W:    hasBitSet(raw, 2),
		N:    hasBitSet(raw, 3),
		Menu: hasBitSet(raw, 4),
	}
}

// Check if the given uint32 value has the given bit set.
func hasBitSet(val uint32, bit uint) bool {
	return (val>>bit)&0b1 != 0
}
