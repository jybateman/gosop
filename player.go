package main

import (
	"fmt"
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

func (p *player) ChangeDir(key string) {
	fmt.Println(p.pos%bsize, p.pos, bsize)
	if (p.pos%bsize) == (bsize/2) {
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
}

func (p *player) MovePlayer(m *maps) {
	tmp := p.pos
	switch p.dir {
	case up:
		p.pos -= m.x*bsize*bsize
	case right:
		p.pos += bsize
	case down:
		p.pos += m.x*bsize*bsize
	case left:
		p.pos -= bsize
	}
	if m.isObstacle(p.pos) {
		p.pos = tmp
	}
}

func NewPlayer(m *maps) player {
	var p player

	p.dir = right
	p.speed = 5
	xy := bytes.IndexByte(m.layout, 'P')
	// p.pos = (xy%m.x*bsize+bsize/2)+xy/m.x*m.x*bsize*(bsize+bsize/2)
	p.pos = (xy%m.x*bsize+bsize/2)+(xy/m.x*m.x*bsize*bsize+bsize/2*bsize*m.x)
	// p.pos = xy/m.x*m.x*bsize*(bsize+bsize/2)
	// txy := (p.pos / (bsize * m.x) * m.x + p.pos % (bsize * m.x)) / bsize
	// txy := p.pos % (bsize * m.x) / bsize
	// fmt.Println("txy:",txy, "xy:", xy)
	return p
}
