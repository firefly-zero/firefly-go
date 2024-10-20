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

// Read a file.
//
// It will first lookup file in the app's ROM directory and then check
// the app writable data directory.
//
// If the file does not exist, the Raw value of the returned File will be nil.
func LoadFile(path string) File {
	pathPtr := unsafe.Pointer(unsafe.StringData(path))
	fileSize := getFileSize(pathPtr, uint32(len(path)))
	if fileSize == 0 {
		return File{nil}
	}
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
