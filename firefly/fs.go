package firefly

import "unsafe"

type Font struct {
	raw []byte
}

func LoadFont(path string, byteSize int) Font {
	raw := make([]byte, byteSize)
	pathPtr := unsafe.Pointer(unsafe.StringData(path))
	rawPtr := unsafe.Pointer(unsafe.SliceData(raw))
	loadRomFile(
		pathPtr, uint32(len(path)),
		rawPtr, uint32(byteSize),
	)
	return Font{raw: raw}
}
