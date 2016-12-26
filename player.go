package main

import (
	"bytes"
)

const (
	up = iota
	right = iota
	down = iota
	left = iota
)

type player struct {
	pos int
	dir int
	char int
	speed int
}

func (p *player) PeekObstacle(m *maps) bool {
	switch p.dir {
	case up:
		return m.isObstacle(p.pos-m.x*bsize*bsize)
	case right:
		return m.isObstacle(p.pos+bsize)
	case down:
		return m.isObstacle(p.pos+m.x*bsize*bsize)
	case left:
		return m.isObstacle(p.pos-bsize)
	}
	return false
}

func (p *player) ChangeDir(key string) {
	switch key {
	case "38":
		p.dir = up
	case "39":
		p.dir = right
	case "40":
		p.dir = down
	case "37":
		p.dir = left
	}
}

func (p *player) MovePlayer(m *maps) {
	tmp := p.pos
	switch p.dir {
	case up:
		p.pos -= m.x
	case right:
		p.pos++
	case down:
		p.pos += m.x
	case left:
		p.pos--
	}
	if m.isObstacle(p.pos) {
		p.pos = tmp
	}
}

func NewPlayer(m *maps) player {
	var p player

	p.dir = right
	p.speed = 5
	p.pos = bytes.IndexByte(m.layout, 'P')
	return p
}
