package firefly

const (
	PadMinX = -1000
	PadMinY = -1000
	PadMaxX = 1000
	PadMaxY = 1000
)

type Pad struct {
	X int
	Y int
}

func (p Pad) Point() Point {
	return Point(p)
}

func (p Pad) Size() Size {
	return Size{W: p.X, H: p.Y}
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
