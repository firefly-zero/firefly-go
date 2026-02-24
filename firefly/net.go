package firefly

import (
	"iter"
	"math/bits"
	"unsafe"
)

// A peer obtained either from [Peers] ([GetPeers]) or from [GetMe].
type AnyPeer interface {
	peerID() uint8
}

// The peer representing the current device.
//
// Can be compared to [Peer] (using [Me.Eq]) or used to [GetSettings].
//
// **IMPORTANT:** Using this type may cause state drift between device in multiplayer.
// See [the docs] for more info.
//
// [the docs]: https://docs.fireflyzero.com/dev/net/
type Me struct {
	raw uint8
}

// Check if the given [Peer] represents the current device.
func (me Me) Eq(peer Peer) bool {
	return me.raw == peer.raw
}

// peerID implements [AnyPeer].
func (me Me) peerID() uint8 {
	return me.raw
}

// The peer ID.
//
// Can be obtained by getting the list of [Peers] using [GetPeers]
// and then iterating over it.
type Peer struct {
	raw uint8
}

// Check if the given peers represent the same device.
func (peer Peer) Eq(other Peer) bool {
	return peer.raw == other.raw
}

// peerID implements [AnyPeer].
func (peer Peer) peerID() uint8 {
	return peer.raw
}

// A combination of all connected peers.
//
// Can be passed in functions like [ReadPad] and [ReadButtons]
// to get the combined input of all peers.
//
// Useful for single-player games that want in multiplayer to handle
// inputs from all devices as one input.
var Combined = Peer{0xFF}

// The list of peers online.
//
// Can be obtained using [GetPeers].
type Peers uint32

// Stash is a serialized binary state of the app that you want to persist
// between app runs and to be available in multiplayer.
//
// For single-player purposes, you can save data in a regular file
// using [DumpFile]. File saved that way can be bigger (and you can create lots of them)
// but it cannot be accessed in multiplayer.
//
// It's your job to serialize data into a binary stash and later parse it.
// Stash can be saved using [SaveStash] and later read using [LoadStash].
type Stash = []byte

// Get the slice of all peers that are online.
func (peers Peers) Slice() []Peer {
	res := make([]Peer, 0, 32)
	for peerID := range uint8(32) {
		peer := Peer{peerID}
		if peers.Contains(peer) {
			res = append(res, peer)
		}
	}
	return res
}

// Iterate over all peers that are online.
//
// Uses the iterators API introduced in Go 1.23.
func (peers Peers) Iter() iter.Seq[Peer] {
	return func(yield func(Peer) bool) {
		for peerID := range uint8(32) {
			peer := Peer{peerID}
			if !peers.Contains(peer) {
				continue
			}
			exit := yield(peer)
			if !exit {
				return
			}
		}
	}
}

// Check if the given [Peer] is online.
func (peers Peers) Contains(peer Peer) bool {
	return peers>>peer.raw&1 != 0
}

// Get how many peers are online.
func (peers Peers) Len() int {
	// Should be converted by TinyGo in a single wasm instruction.
	return bits.OnesCount32(uint32(peers))
}

// Get the peer corresponding to the local device.
func GetMe() Me {
	return Me{uint8(getMe())}
}

// Get the list of peers that are currently online.
//
// Includes the local device.
//
// It can be used to detect if multiplayer is active:
// if there is more than 1 peer, you're playing with friends.
func GetPeers() Peers {
	return Peers(getPeers())
}

// Save the given [Stash].
//
// When called, the stash for the given peer will be stored in RAM.
// Calling [LoadStash] for the same peer will return that stash.
// On exit, the runtime will persist the stash in FS.
// Next time the app starts, calling [LoadStash] will restore the stash
// saved earlier.
func SaveStash(p Peer, b Stash) {
	ptr := unsafe.Pointer(unsafe.SliceData(b))
	saveStash(uint32(p.raw), ptr, uint32(len(b)))
}

// Load [Stash] saved earlier (in this or previous run) by [SaveStash].
//
// The buffer should be big enough to fit the stash.
// If it's not, the stash will be truncated.
// If there is no stash or it's empty, nil is returned.
//
// If the given buffer is nil, a new buffer will be allocated
// big enough to fit the biggest allowed stash. At the moment, it is 80 bytes.
func LoadStash(p Peer, buf []byte) Stash {
	if buf == nil {
		buf = make([]byte, 80)
	}
	ptr := unsafe.Pointer(unsafe.SliceData(buf))
	size := loadStash(uint32(p.raw), ptr, uint32(len(buf)))
	if size == 0 {
		return nil
	}
	return buf[:size]
}
