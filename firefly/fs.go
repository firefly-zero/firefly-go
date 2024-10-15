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

// Read a file from the app ROM.
func LoadROMFile(path string) File {
	pathPtr := unsafe.Pointer(unsafe.StringData(path))
	fileSize := getRomFileSize(pathPtr, uint32(len(path)))
	raw := make([]byte, fileSize)
	rawPtr := unsafe.Pointer(unsafe.SliceData(raw))
	loadRomFile(
		pathPtr, uint32(len(path)),
		rawPtr, fileSize,
	)
	return File{raw}
}

// Read a file from the app data dir.
func LoadDataFile(path string) File {
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
func DumpDataFile(path string, raw []byte) {
	pathPtr := unsafe.Pointer(unsafe.StringData(path))
	rawPtr := unsafe.Pointer(unsafe.SliceData(raw))
	dumpFile(
		pathPtr, uint32(len(path)),
		rawPtr, uint32(len(raw)),
	)
}

// Remove a file from the app data dir.
func RemoveDataFile(path string) {
	pathPtr := unsafe.Pointer(unsafe.StringData(path))
	removeFile(pathPtr, uint32(len(path)))
}
