package firefly

import "unsafe"

type Font struct {
	raw []byte
}

func LoadFont(path string, byteSize int) Font {
	raw := make([]byte, 0, byteSize)
	pathBytes := []byte(path)
	pathPtr := unsafe.Pointer(unsafe.SliceData(pathBytes))
	loadRomFile(
		pathPtr, uint32(len(path)),
		unsafe.Pointer(&raw), uint32(byteSize),
	)
	return Font{raw: raw}
}
