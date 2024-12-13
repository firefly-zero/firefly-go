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

// Ensure that the loaded file exists.
func (f File) Must() File {
	if f.Raw == nil {
		panic("file not found")
	}
	return f
}

// Read a file.
//
// It will first lookup file in the app's ROM directory and then check
// the app writable data directory.
//
// If the file does not exist, the Raw value of the returned File will be nil.
//
// The second argument is the buffer in which the file should be loaded.
// If the buffer is smaller than the file content, it gets cut.
// If the buffer is nil, a new buffer of sufficient size will be allocated.
func LoadFile(path string, buf []byte) File {
	if buf == nil {
		return loadAllocFile(path)
	}
	return loadFileInto(path, buf)
}

// Allocate a new buffer and load the file into it.
func loadAllocFile(path string) File {
	pathPtr := unsafe.Pointer(unsafe.StringData(path))
	fileSize := getFileSize(pathPtr, uint32(len(path)))
	if fileSize == 0 {
		return File{nil}
	}
	buf := make([]byte, fileSize)
	bufPtr := unsafe.Pointer(unsafe.SliceData(buf))
	loadFile(
		pathPtr, uint32(len(path)),
		bufPtr, uint32(len(buf)),
	)
	return File{buf}
}

// Load the file into the given buffer.
func loadFileInto(path string, buf []byte) File {
	pathPtr := unsafe.Pointer(unsafe.StringData(path))
	bufPtr := unsafe.Pointer(unsafe.SliceData(buf))
	fileSize := loadFile(
		pathPtr, uint32(len(path)),
		bufPtr, uint32(len(buf)),
	)
	if fileSize == 0 {
		return File{nil}
	}
	return File{buf[:fileSize]}
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
