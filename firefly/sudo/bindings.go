package sudo

import "unsafe"

//go:wasmimport sudo list_dirs_buf_size
func listDirsBufSize(pathPtr unsafe.Pointer, pathLen uint32) uint32

//go:wasmimport sudo list_dirs
func listDirs(pathPtr unsafe.Pointer, pathLen uint32, bufPtr unsafe.Pointer, bufLen uint32) uint32

//go:wasmimport sudo run_app
func runApp(authorPtr unsafe.Pointer, authorLen uint32, appPtr unsafe.Pointer, appLen uint32)

//go:wasmimport sudo load_file
func loadFile(pathPtr unsafe.Pointer, pathLen uint32, bufPtr unsafe.Pointer, bufLen uint32) uint32

//go:wasmimport sudo get_file_size
func getFileSize(pathPtr unsafe.Pointer, pathLen uint32) uint32
