package main

import (
	"fmt"
	"strconv"

	"github.com/firefly-zero/firefly-go/firefly"
	"github.com/firefly-zero/firefly-go/firefly/random"
)

var count = 0

func init() {
	firefly.Boot = boot
	firefly.Update = update
}

func boot() {
	firefly.LogDebug("hello from wasm, boot callback")
	firefly.LogDebug(fmt.Sprintf("random int: %d", random.Int()))
	firefly.LogDebug(fmt.Sprintf("random float: %f", random.Float32()))
}

func update() {
	count++
	if count%60 == 0 {
		firefly.LogDebug("updates: " + strconv.Itoa(count))
	}
}
