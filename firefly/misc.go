package firefly

import "unsafe"

func LogDebug(t string) {
	ptr := unsafe.Pointer(unsafe.StringData(t))
	logDebug(ptr, uint32(len(t)))
}

func LogError(t string) {
	ptr := unsafe.Pointer(unsafe.StringData(t))
	logError(ptr, uint32(len(t)))
}

func SetSeed(seed uint32) {
	setSeed(seed)
}

func GetRandom() uint32 {
	return getRandom()
}

func Quit() {
	quit()
}
