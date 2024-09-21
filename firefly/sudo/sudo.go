package sudo

import (
	"unsafe"

	"github.com/firefly-zero/firefly-go/firefly"
)

// Get the name of all dirs in the given dir.
func ListDirs(path string) []string {
	pathPtr := unsafe.Pointer(unsafe.StringData(path))
	bufSize := listDirsBufSize(pathPtr, uint32(len(path)))
	buf := make([]byte, bufSize)
	bufPtr := unsafe.Pointer(unsafe.SliceData(buf))
	listDirs(
		pathPtr, uint32(len(path)),
		bufPtr, bufSize,
	)
	res := make([]string, 0)
	for len(buf) != 0 {
		size := buf[0]
		name := string(buf[1:size])
		res = append(res, name)
		buf = buf[size+1:]
	}
	return res
}

// Exit the current app and run the given app instead.
//
// Calling this function will NOT immediately terminate the app.
// The current update will be finished first.
func RunApp(author string, app string) {
	authorPtr := unsafe.Pointer(unsafe.StringData(author))
	appPtr := unsafe.Pointer(unsafe.StringData(app))
	runApp(
		authorPtr, uint32(len(author)),
		appPtr, uint32(len(app)),
	)
}

// Read a file from the FS.
func LoadFile(path string) firefly.File {
	pathPtr := unsafe.Pointer(unsafe.StringData(path))
	fileSize := getFileSize(pathPtr, uint32(len(path)))
	raw := make([]byte, fileSize)
	rawPtr := unsafe.Pointer(unsafe.SliceData(raw))
	loadFile(
		pathPtr, uint32(len(path)),
		rawPtr, fileSize,
	)
	return firefly.File{Raw: raw}
}
