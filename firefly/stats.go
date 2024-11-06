package firefly

type Badge uint8

// type Progress struct {
// 	Done uint16
// 	Goal uint16
// }

// func (p Progress) Earned() bool {
// 	return p.Done >= p.Goal
// }

// Add the given value to the progress ...
func AddProgress(p Peer, b Badge, v int16) {
	addProgress(uint32(p), uint32(b), int32(v))
	// return Progress{
	// 	Done: uint16(r >> 16),
	// 	Goal: uint16(r),
	// }
}

func AddScore(p Peer, b Badge, v int16) {
	addScore(uint32(p), uint32(b), int32(v))
}
