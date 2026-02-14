package main

import (
	"bytes"
	"strconv"

	"github.com/firefly-zero/firefly-go/firefly"
)

var (
	count = 0
	buf   bytes.Buffer
)

func init() {
	firefly.Boot = boot
	firefly.Update = update

	buf.Grow(32)
}

func boot() {
	firefly.LogDebug("hello from wasm, boot callback")
}

func update() {
	count++
	if count%60 == 0 {
		firefly.LogDebug("updates: " + strconv.Itoa(count))
	}

	if count%10 == 0 {
		// Example of zero-allocation formatting
		// WARN: Not recommended to write code like this.
		// Only use this pattern in code paths where performance is critical.
		buf.Reset()
		buf.WriteString("... ")
		buf.Write(strconv.AppendInt(buf.AvailableBuffer(), int64(count), 10))
		buf.WriteString(" ...")
		firefly.LogDebugBytes(buf.Bytes())
	}
}
