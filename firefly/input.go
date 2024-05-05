package firefly

type TouchPad struct {
	X int
	Y int
}

func (p TouchPad) Point() Point {
	return Point(p)
}

func (p TouchPad) Size() Size {
	return Size{W: p.X, H: p.Y}
}

func ReadPad(p Player) (TouchPad, bool) {
	raw := readPad(uint32(p))
	pressed := raw != 0xffff
	if !pressed {
		return TouchPad{}, false
	}
	pad := TouchPad{
		X: int(int16(raw >> 16)),
		Y: int(int16(raw)),
	}
	return pad, true
}
