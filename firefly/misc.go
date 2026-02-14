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

// Log a debug message from byte slice.
//
// Input slice must be valid UTF-8 formatted (without BOM).
//
// Useful when you are building strings in a byte array or slice, such as via a [bytes.Buffer].
// This function allows you to use such content without allocating a new string.
//
// It is allowed to modify the byte slice after the function call.
func LogDebugBytes(t []byte) {
	ptr := unsafe.Pointer(unsafe.SliceData(t))
	logDebug(ptr, uint32(len(t)))
}

// Log an error message from byte slice.
//
// Input slice must be valid UTF-8 formatted (without BOM).
//
// Useful when you are building strings in a byte array or slice, such as via a [bytes.Buffer].
// This function allows you to use such content without allocating a new string.
//
// It is allowed to modify the byte slice after the function call.
func LogErrorBytes(t []byte) {
	ptr := unsafe.Pointer(unsafe.SliceData(t))
	logError(ptr, uint32(len(t)))
}

// Set the seed used to generate random values.
func SetSeed(seed uint32) {
	setSeed(seed)
}

// Get a random value.
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
