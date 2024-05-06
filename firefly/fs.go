package firefly

import "unsafe"

type File struct {
	raw []byte
}

func (f File) Font() Font {
	return Font(f)
}

func (f File) Image() Image {
	return Image(f)
}

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

func LoadROMFile(path string) File {
	pathPtr := unsafe.Pointer(unsafe.StringData(path))
	fileSize := getRomFileSize(pathPtr, uint32(len(path)))
	raw := make([]byte, fileSize)
	rawPtr := unsafe.Pointer(unsafe.SliceData(raw))
	loadRomFile(
		pathPtr, uint32(len(path)),
		rawPtr, uint32(fileSize),
	)
	return File{raw}
}

func LoadDataFile(path string) File {
	pathPtr := unsafe.Pointer(unsafe.StringData(path))
	fileSize := getFileSize(pathPtr, uint32(len(path)))
	raw := make([]byte, fileSize)
	rawPtr := unsafe.Pointer(unsafe.SliceData(raw))
	loadFile(
		pathPtr, uint32(len(path)),
		rawPtr, uint32(fileSize),
	)
	return File{raw}
}

func DumpDataFile(path string, raw []byte) {
	pathPtr := unsafe.Pointer(unsafe.StringData(path))
	rawPtr := unsafe.Pointer(unsafe.SliceData(raw))
	dumpFile(
		pathPtr, uint32(len(path)),
		rawPtr, uint32(len(raw)),
	)
}

func RemoveDataFile(path string) {
	pathPtr := unsafe.Pointer(unsafe.StringData(path))
	removeFile(pathPtr, uint32(len(path)))
}
