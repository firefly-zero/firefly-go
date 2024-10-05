package firefly

import "unsafe"

// A file loaded from the filesystem.
type File struct {
	Raw []byte
}

// Convert the File to a Font.
func (f File) Font() Font {
	return Font{f.Raw}
}

// Convert the File to an Image.
func (f File) Image() Image {
	return Image{f.Raw}
}

// A loaded font file.
//
// Can be loaded using [LoadROMFile].
type Font struct {
	raw []byte
}

// A loaded image file.
//
// Can be loaded using [LoadROMFile].
type Image struct {
	raw []byte
}

// Get a rectangle subregion of the image.
func (i Image) Sub(p Point, s Size) SubImage {
	return SubImage{raw: i.raw, point: p, size: s}
}

// A subregion of an image. Constructed using [Image.Sub].
type SubImage struct {
	raw   []byte
	point Point
	size  Size
}

// Read a file.
//
// It will first lookup file in the app's ROM directory and then check
// the app writable data directory.
func LoadFile(path string) File {
	pathPtr := unsafe.Pointer(unsafe.StringData(path))
	fileSize := getFileSize(pathPtr, uint32(len(path)))
	raw := make([]byte, fileSize)
	rawPtr := unsafe.Pointer(unsafe.SliceData(raw))
	loadFile(
		pathPtr, uint32(len(path)),
		rawPtr, fileSize,
	)
	return File{raw}
}

// Write a file into the app data dir.
func DumpFile(path string, raw []byte) {
	pathPtr := unsafe.Pointer(unsafe.StringData(path))
	rawPtr := unsafe.Pointer(unsafe.SliceData(raw))
	dumpFile(
		pathPtr, uint32(len(path)),
		rawPtr, uint32(len(raw)),
	)
}

// Remove a file from the app data dir.
func RemoveFile(path string) {
	pathPtr := unsafe.Pointer(unsafe.StringData(path))
	removeFile(pathPtr, uint32(len(path)))
}
