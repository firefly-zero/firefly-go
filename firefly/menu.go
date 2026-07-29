package firefly

import "unsafe"

type MenuItem uint8

func OpenMenu() {
	openMenu()
}

func AddMenuItem(i MenuItem, name string) {
	ptr := unsafe.Pointer(unsafe.StringData(name))
	size := uint32(len(name))
	addMenuItem(uint32(i), ptr, size)
}

func RemoveMenuItem(i MenuItem) {
	removeMenuItem(uint32(i))
}
