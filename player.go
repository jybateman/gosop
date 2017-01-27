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
	hist []string
	pos int
	dir int
	char int
	speed int
}

func (p *player) PeekObstacle(m *maps, key string) bool {
	switch key {
	case "38":
		return m.isObstacle(p.pos-m.x)
	case "39":
		return m.isObstacle(p.pos+1)
	case "40":
		return m.isObstacle(p.pos+m.x)
	case "37":
		return m.isObstacle(p.pos-1)
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
	m.layout[p.pos] = ' '
}

func NewPlayer(m *maps) player {
	var p player

	p.dir = right
	p.speed = 5
	p.pos = bytes.IndexByte(m.layout, 'P')
	return p
}
