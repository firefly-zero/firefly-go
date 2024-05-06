package firefly

const (
	PadMinX = -1000
	PadMinY = -1000
	PadMaxX = 1000
	PadMaxY = 1000
)

const dPadThreshold = 100

type Pad struct {
	X int
	Y int
}

func (p Pad) Radius() float32 {
	r := p.X*p.X + p.Y*p.Y
	return sqrt(float32(r))
}

func (p Pad) Azimuth() Angle {
	r := atan(float32(p.Y) / float32(p.X))
	return Radians(r)
}

func (p Pad) Point() Point {
	return Point(p)
}

func (p Pad) Size() Size {
	return Size{W: p.X, H: p.Y}
}

func (p Pad) DPad() DPad {
	return DPad{
		Left:  p.X <= -dPadThreshold,
		Right: p.X >= dPadThreshold,
		Up:    p.Y <= -dPadThreshold,
		Down:  p.Y >= dPadThreshold,
	}
}

type DPad struct {
	Left  bool
	Right bool
	Up    bool
	Down  bool
}

func (p DPad) JustPressed(old DPad) DPad {
	p.Left = p.Left && !old.Left
	p.Right = p.Right && !old.Right
	p.Up = p.Up && !old.Up
	p.Down = p.Down && !old.Down
	return p
}

func (p DPad) JustReleased(old DPad) DPad {
	p.Left = !p.Left && old.Left
	p.Right = !p.Right && old.Right
	p.Up = !p.Up && old.Up
	p.Down = !p.Down && old.Down
	return p
}

func (p DPad) Held(old DPad) DPad {
	p.Left = p.Left && old.Left
	p.Right = p.Right && old.Right
	p.Up = p.Up && old.Up
	p.Down = p.Down && old.Down
	return p
}

type Buttons struct {
	// If "a" button is pressed.
	A bool

	// If "b" button is pressed.
	B bool

	// If "x" button is pressed.
	X bool

	// If "y" button is pressed.
	Y bool

	// If "menu" button is pressed.
	//
	// For singleplayer games, the button press is always intercepted by the runtime.
	Menu bool
}

func (p Buttons) JustPressed(old Buttons) Buttons {
	p.A = p.A && !old.A
	p.B = p.B && !old.B
	p.X = p.X && !old.X
	p.Y = p.Y && !old.Y
	p.Menu = p.Menu && !old.Y
	return p
}

func (p Buttons) JustReleased(old Buttons) Buttons {
	p.A = !p.A && old.A
	p.B = !p.B && old.B
	p.X = !p.X && old.X
	p.Y = !p.Y && old.Y
	p.Y = !p.Y && old.Y
	p.Y = !p.Y && old.Y
	p.Menu = !p.Menu && old.Menu
	return p
}

func (p Buttons) Held(old Buttons) Buttons {
	p.A = p.A && old.A
	p.B = p.B && old.B
	p.X = p.X && old.X
	p.Y = p.Y && old.Y
	p.Y = p.Y && old.Y
	p.Y = p.Y && old.Y
	p.Menu = p.Menu && old.Menu
	return p
}

func ReadPad(p Player) (Pad, bool) {
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

func ReadButtons(p Player) Buttons {
	raw := readButtons(uint32(p))
	return Buttons{
		A:    hasBitSet(raw, 0),
		B:    hasBitSet(raw, 1),
		X:    hasBitSet(raw, 2),
		Y:    hasBitSet(raw, 3),
		Menu: hasBitSet(raw, 4),
	}
}

func hasBitSet(val uint32, bit uint) bool {
	return (val>>bit)&0b1 != 0
}
