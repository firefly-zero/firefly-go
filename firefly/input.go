package firefly

type TouchPad struct {
	X i32
	Y i32
}

func ReadPad(p Player) (TouchPad, bool) {
	raw := readPad(uint32(p))
	pressed := raw != 0xffff
	pad := TouchPad{
		X: i32(int16(raw >> 16)),
		Y: i32(int16(raw)),
	}
	return pad, pressed
}
