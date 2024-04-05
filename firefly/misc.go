package firefly

import "unsafe"

func LogDebug(t string) {
	bytes := []byte(t)
	ptr := unsafe.Pointer(unsafe.SliceData(bytes))
	logDebug(ptr, uint32(len(t)))
}
