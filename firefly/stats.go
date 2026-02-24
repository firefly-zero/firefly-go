package firefly

// A badge (aka achievement) ID.
type Badge uint8

// A board (aka score board / leader board) ID.
type Board uint8

type Progress struct {
	// How many points the player already has.
	Done uint16
	// How many points the player needs to earn the badge.
	Goal uint16
}

// True if the player got enough points to unlock the badge.
func (p Progress) Earned() bool {
	return p.Done >= p.Goal
}

// Get the progress of earning the badge.
func GetProgress(p Peer, b Badge) Progress {
	return AddProgress(p, b, 0)
}

// Add the given value to the progress for the badge.
//
// May be negative if you want to decrease the progress.
// If zero, does not change the progress.
//
// If the Peer is [Combined], the progress is added to every peer
// and the returned value is the lowest progress.
func AddProgress(p Peer, b Badge, v int16) Progress {
	r := addProgress(uint32(p.raw), uint32(b), int32(v))
	return Progress{
		Done: uint16(r >> 16),
		Goal: uint16(r),
	}
}

// Get the personal best of the player.
func GetScore(p Peer, b Board) int16 {
	return AddScore(p, b, 0)
}

// Add the given score to the board.
//
// May be negative if you want the lower scores
// to rank higher. Zero value is not added to the board.
//
// If the Peer is [Combined], the score is added for every peer
// and the returned value is the lowest of their best scores.
func AddScore(p Peer, b Board, v int16) int16 {
	s := addScore(uint32(p.raw), uint32(b), int32(v))
	return int16(s)
}
