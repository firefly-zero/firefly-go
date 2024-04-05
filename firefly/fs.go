package firefly

import "unsafe"

type Font struct {
	raw []byte
}

func LoadFont(path string, byteSize int) Font {
	raw := make([]byte, 0, byteSize)
	loadRomFile(
		unsafe.Pointer(&path), uint32(len(path)),
		unsafe.Pointer(&raw), uint32(byteSize),
	)
	return Font{raw: raw}
}
