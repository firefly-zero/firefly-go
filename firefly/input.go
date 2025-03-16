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

// The minimum X or Y value when converting Pad into DPad
// for the direction to be considered pressed.
// We do that to provide a dead zone in the middle of the pad.
const dPadThreshold = 100

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

// The angle of the [polar coordinate] of the touch point.
//
// [polar coordinate]: https://en.wikipedia.org/wiki/Polar_coordinate_system
func (p Pad) Azimuth() Angle {
	r := math.Pi / 2. * tinymath.Atan2Norm(float32(p.Y), float32(p.X))
	return Radians(r)
}

// Convert the Pad into a Point.
func (p Pad) Point() Point {
	return Point(p)
}

// Convert the Pad into a Size.
func (p Pad) Size() Size {
	return Size{W: p.X, H: p.Y}
}

// Convert the Pad into DPad.
func (p Pad) DPad() DPad {
	return DPad{
		Left:  p.X <= -dPadThreshold,
		Right: p.X >= dPadThreshold,
		Up:    p.Y <= -dPadThreshold,
		Down:  p.Y >= dPadThreshold,
	}
}

// DPad-like representation of the [Pad].
//
// Constructed with [Pad.DPad]. Useful for simple games and ports.
// The middle of the pad is a "dead zone" pressing which will not activate any direction.
//
// Invariant: it's not possible for opposite directions (left and right, or down and up)
// to be active at the same time. However, it's possible for neighboring directions
// (like up and right) to be active at the same time if the player presses a diagonal.
type DPad struct {
	Left  bool
	Right bool
	Up    bool
	Down  bool
}

// Given the old state, get directions that were not pressed but are pressed now.
func (p DPad) JustPressed(old DPad) DPad {
	p.Left = p.Left && !old.Left
	p.Right = p.Right && !old.Right
	p.Up = p.Up && !old.Up
	p.Down = p.Down && !old.Down
	return p
}

// Given the old state, get directions that were pressed but aren't pressed now.
func (p DPad) JustReleased(old DPad) DPad {
	p.Left = !p.Left && old.Left
	p.Right = !p.Right && old.Right
	p.Up = !p.Up && old.Up
	p.Down = !p.Down && old.Down
	return p
}

// Given the old state, get directions that were pressed and are still pressed now.
func (p DPad) Held(old DPad) DPad {
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
func (p Buttons) AnyPressed() bool {
	return p.S || p.E || p.W || p.N || p.Menu
}

// Get the current touch pad state.
//
// The peer can be [Combined] or one of the [GetPeers].
func ReadPad(p Peer) (Pad, bool) {
	raw := readPad(uint32(p))
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
	raw := readButtons(uint32(p))
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
