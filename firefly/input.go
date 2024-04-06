package firefly

type TouchPad struct {
	X i32
	Y i32
}

func ReadLeft() TouchPad {
	raw := readLeft()
	return TouchPad{
		X: i32(int16(raw >> 16)),
		Y: i32(int16(raw)),
	}
}
