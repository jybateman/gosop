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
}

func (p *player) move(xy int) {
	
}

func (p *player) MovePlayer(width int) {
	switch p.dir {
	case up:
		p.pos -= width*bsize
	case right:
		p.pos++
	case down:
		p.pos += width*bsize
	case left:
		p.pos--
		
	}
}

func NewPlayer(m *maps) player {
	var p player

	p.dir = up
	xy := bytes.IndexByte(m.layout, 'P')
	fmt.Println(xy, m.x, bsize)
	p.pos = xy%m.x*bsize+xy/m.x*m.x*bsize*bsize
	// txy := (p.pos / (bsize * m.x) * m.x + p.pos % (bsize * m.x)) / bsize
	txy := p.pos % (bsize * m.x) / bsize
	fmt.Println("txy:",txy, "xy:", xy)
	return p
}
