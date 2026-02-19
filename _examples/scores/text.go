package main

import (
	"strconv"

	"github.com/firefly-zero/firefly-go/firefly"
)

func init() {
	firefly.Boot = boot
	firefly.Update = update
	firefly.Render = render
}

const board firefly.Board = 1

var (
	font    firefly.Font
	players []*Player
)

type Player struct {
	peer firefly.Peer
	name string
	curr int16
	best int16
}

func boot() {
	font = firefly.LoadFile("font", nil).Font()
	peers := firefly.GetPeers().Slice()
	players = make([]*Player, len(peers))
	for i, peer := range peers {
		players[i] = &Player{
			peer: peer,
			name: firefly.GetName(peer),
			curr: 0,
			best: firefly.GetScore(peer, board),
		}
	}
}

func update() {
	for _, p := range players {
		btns := firefly.ReadButtons(p.peer)
		if btns.S {
			p.curr += 1
		}
		if btns.E {
			p.best = firefly.AddScore(p.peer, board, p.curr)
			p.curr = 0
		}
	}
}

func render() {
	firefly.ClearScreen(firefly.ColorWhite)
	c := firefly.ColorDarkBlue
	for i, p := range players {
		y := 10 + 10*i
		font.Draw(p.name, firefly.Point{X: 10, Y: y}, c)
		font.Draw(formatInt(p.curr), firefly.Point{X: 120, Y: y}, c)
		font.Draw(formatInt(p.best), firefly.Point{X: 150, Y: y}, c)
	}
}

func formatInt(i int16) string {
	return strconv.FormatInt(int64(i), 10)
}
