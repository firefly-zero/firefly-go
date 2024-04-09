package firefly

import "unsafe"

type Font struct {
	raw []byte
}

type Image struct {
	raw []byte
}

func (i Image) Sub(p Point, s Size) SubImage {
	return SubImage{raw: i.raw, point: p, size: s}
}

type SubImage struct {
	raw   []byte
	point Point
	size  Size
}

func LoadFont(path string) Font {
	pathPtr := unsafe.Pointer(unsafe.StringData(path))
	fileSize := getRomFileSize(pathPtr, uint32(len(path)))
	raw := make([]byte, fileSize)
	rawPtr := unsafe.Pointer(unsafe.SliceData(raw))
	loadRomFile(
		pathPtr, uint32(len(path)),
		rawPtr, uint32(fileSize),
	)
	return Font{raw: raw}
}

func LoadImage(path string) Image {
	pathPtr := unsafe.Pointer(unsafe.StringData(path))
	fileSize := getRomFileSize(pathPtr, uint32(len(path)))
	raw := make([]byte, fileSize)
	rawPtr := unsafe.Pointer(unsafe.SliceData(raw))
	loadRomFile(
		pathPtr, uint32(len(path)),
		rawPtr, uint32(fileSize),
	)
	return Image{raw: raw}
}
