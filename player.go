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

func NewPlayer(m *maps) player {
	var p player

	xy := bytes.IndexByte(m.layout, 'P')
	fmt.Println(xy, m.x, bsize)
	p.pos = xy%m.x*bsize+xy/m.x*m.x*bsize*bsize
	txy := p.pos / bsize / m.x
	fmt.Println("txy:", txy, "xy:", xy)
	return p
}
