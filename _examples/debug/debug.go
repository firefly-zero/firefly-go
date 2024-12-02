package main

import (
	"strconv"

	"github.com/firefly-zero/firefly-go/firefly"
)

var count = 0

func init() {
	firefly.Boot = boot
	firefly.Update = update
}

func boot() {
	firefly.LogDebug("hello from wasm, boot callback")
}

func update() {
	count++
	if count%60 == 0 {
		firefly.LogDebug("updates: " + strconv.Itoa(count))
	}
}
