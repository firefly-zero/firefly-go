package firefly

import "math/bits"

// The peer ID.
type Peer uint8

// Peer value that can be passed to [ReadPad] and [ReadButtons]
// to get the combined input of all peers.
//
// Useful for single-player games that want in multi-player to handle
// inputs from all devices as one input.
const Combined Peer = 0xFF

// The map of peers online.
type Peers uint32

// Get the slice of all peers that are online.
func (peers Peers) Slice() []Peer {
	res := make([]Peer, 0, 32)
	for peer := Peer(0); peer < 32; peer++ {
		if peers.IsOnline(peer) {
			res = append(res, peer)
		}
	}
	return res
}

// Check if the given peer is online.
func (peers Peers) IsOnline(peer Peer) bool {
	return peers>>peer&1 != 0
}

// Get how many peers are online.
func (peers Peers) Len() int {
	// Should be converted by TinyGo in a single wasm instruction.
	return bits.OnesCount32(uint32(peers))
}

// Get the peer ID representing the local device.
func GetMe() Peer {
	return Peer(getMe())
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
