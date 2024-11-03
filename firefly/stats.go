package firefly

type Badge uint8

type Progress struct {
	Done uint16
	Goal uint16
}

func (p Progress) Earned() bool {
	return p.Done >= p.Goal
}

func GetProgress(b Badge) Progress {
	p := getProgress(uint32(b))
	return Progress{
		Done: uint16(p >> 16),
		Goal: uint16(p),
	}
}

func AddProgress(b Badge, v int16) Progress {
	p := addProgress(uint32(b), int32(v))
	return Progress{
		Done: uint16(p >> 16),
		Goal: uint16(p),
	}
}
