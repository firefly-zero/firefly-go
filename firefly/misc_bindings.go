package firefly

import "unsafe"

//go:wasmimport misc log_debug
func logDebug(ptr unsafe.Pointer, len uint32)
