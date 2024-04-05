package firefly

import "unsafe"

//go:wasmimport fs load_rom_file
func loadRomFile(
	pathPtr unsafe.Pointer, pathLen uint32,
	bufPtr unsafe.Pointer, bufLen uint32,
) uint32
