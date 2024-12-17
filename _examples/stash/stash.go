package main

import (
	"encoding/binary"
	"strconv"

	"github.com/firefly-zero/firefly-go/firefly"
)

func init() {
	firefly.Boot = boot
	firefly.Update = update
	firefly.Render = render
	firefly.BeforeExit = beforeExit
}

type State struct {
	clicks  uint32
	pressed bool
}

var (
	states map[firefly.Peer]*State = make(map[firefly.Peer]*State)
	peers  firefly.Peers
	font   firefly.Font
)

func boot() {
	font = firefly.LoadFile("font", nil).Font()
	peers = firefly.GetPeers()
	buf := make([]byte, 4)
	for _, peer := range peers.Slice() {
		stash := firefly.LoadStash(peer, buf)
		if stash != nil {
			states[peer] = &State{
				clicks: binary.LittleEndian.Uint32(stash),
			}
		}
	}
}

func update() {
	for _, peer := range peers.Slice() {
		buttons := firefly.ReadButtons(peer)
		pressed := buttons.S
		state, found := states[peer]
		if !found {
			state = &State{}
			states[peer] = state
		}
		if !state.pressed && pressed {
			state.clicks++
		}
		state.pressed = pressed
	}
}

func render() {
	firefly.ClearScreen(firefly.ColorWhite)
	for i, peer := range peers.Slice() {
		state := states[peer]
		text := strconv.FormatUint(uint64(state.clicks), 10)
		point := firefly.Point{X: 60, Y: 60 + i*10}
		color := firefly.ColorBlack
		if state.pressed {
			color = firefly.ColorRed
		}
		firefly.DrawText(text, font, point, color)
	}
}

func beforeExit() {
	peer := firefly.GetMe()
	state := states[peer]
	buf := binary.LittleEndian.AppendUint32(nil, state.clicks)
	firefly.SaveStash(peer, buf)
}
