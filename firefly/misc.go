package firefly

import "unsafe"

// Log a debug message.
func LogDebug(t string) {
	ptr := unsafe.Pointer(unsafe.StringData(t))
	logDebug(ptr, uint32(len(t)))
}

// Log an error message.
func LogError(t string) {
	ptr := unsafe.Pointer(unsafe.StringData(t))
	logError(ptr, uint32(len(t)))
}

// Set the seed used to generate random values.
func SetSeed(seed uint32) {
	setSeed(seed)
}

// Get a random value from [0-[math.MaxUint32]]
//
// See also the [github.com/firefly-zero/firefly-go/firefly/random]
// package provides more ways to generate random numbers.
func GetRandom() uint32 {
	return getRandom()
}

// Get human-readable name of the given peer.
func GetName(p Peer) string {
	buf := [16]byte{}
	ptr := unsafe.Pointer(&buf)
	length := getName(uint32(p), ptr)
	return unsafe.String(&buf[0], length)
}

// Exit the app after the current update is finished.
func Quit() {
	quit()
}

// Restart the app.
func Restart() {
	restart()
}
