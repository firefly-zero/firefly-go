package firefly

type TouchPad struct {
	X int16
	Y int16
}

func ReadLeft() TouchPad {
	raw := readLeft()
	return TouchPad{
		X: int16(raw >> 16),
		Y: int16(raw),
	}
}
