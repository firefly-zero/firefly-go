package firefly

type Badge uint8

type Progress struct {
	Done uint16
	Goal uint16
}

func (p Progress) Earned() bool {
	return p.Done >= p.Goal
}

func GetProgress(p Peer, b Badge) Progress {
	return AddProgress(p, b, 0)
}

// Add the given value to the progress for the badge.
func AddProgress(p Peer, b Badge, v int16) Progress {
	r := addProgress(uint32(p), uint32(b), int32(v))
	return Progress{
		Done: uint16(r >> 16),
		Goal: uint16(r),
	}
}

func AddScore(p Peer, b Badge, v int16) {
	addScore(uint32(p), uint32(b), int32(v))
}
