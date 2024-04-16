package firefly

type TouchPad struct {
	X i32
	Y i32
}

func ReadPad() (TouchPad, bool) {
	raw := readPad()
	pressed := raw != 0xffff
	pad := TouchPad{
		X: i32(int16(raw >> 16)),
		Y: i32(int16(raw)),
	}
	return pad, pressed
}
