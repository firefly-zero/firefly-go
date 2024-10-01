package main

import (
	"github.com/firefly-zero/firefly-go/firefly"
	"github.com/firefly-zero/firefly-go/firefly/audio"
)

func init() {
	firefly.Boot = boot
}

func boot() {
	audio.Out.AddSine(audio.A4, 0.)
}
