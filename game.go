package main

import (
	"fmt"
	"time"
	
	"golang.org/x/net/websocket"
)

type game struct {
	m maps
	p []*player
}

func NewGame() *game {
	var g game
	
	g.m = NewMap()
	return &g
}

func (g *game) StartGame(c *websocket.Conn) {
	b := make([]byte, 2)
	p := NewPlayer(&g.m)
	g.p = append(g.p, &p)
	fmt.Println(p)
	c.Write([]byte(fmt.Sprintf("%v", g.m)))
	dl := time.Now().Add(time.Millisecond*100)
	for {
		c.SetReadDeadline(dl)
		c.Read(b)
		if dl.Before(time.Now()) {
			dl = time.Now().Add(time.Millisecond*100)
			p.ChangeDir(string(b))
			p.MovePlayer(&g.m)
			for _, players := range g.p {
				x := players.pos%g.m.x
				y := players.pos/g.m.x
				c.Write([]byte(fmt.Sprintf("%v %v", x, y)))
			}
		}
	}
}
