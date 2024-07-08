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

// Get a random value.
func GetRandom() uint32 {
	return getRandom()
}

// Exit the app after the current update is finished.
func Quit() {
	quit()
}

// Restart the app.
func Restart() {
	restart()
}
