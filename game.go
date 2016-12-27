package main

import (
	"io"
	"fmt"
	"time"
	
	"golang.org/x/net/websocket"
)

type game struct {
	m maps
	p []*player
}

func (g *game) RemovePlayer(p *player) {
	for idx, val := range g.p {
		if val == p {
			g.p = g.p[:idx+copy(g.p[idx:], g.p[idx+1:])]		
		}
	}
}

func (g *game) StartGame(c *websocket.Conn) {
	b := make([]byte, 2)
	p := NewPlayer(&g.m)
	g.p = append(g.p, &p)
	fmt.Println(p)
	c.Write([]byte(fmt.Sprintf("%v", g.m)))
	dl := time.Now()
	for {
		c.SetReadDeadline(dl)
		_, err := c.Read(b)
		if err == io.EOF {
			fmt.Println(err)
			g.RemovePlayer(&p)
			return
		}
		if dl.Before(time.Now()) {
			dl = time.Now().Add(time.Millisecond*200)
			if !p.PeekObstacle(&g.m, string(b)) {
				p.ChangeDir(string(b))
			}
			p.MovePlayer(&g.m)
			for _, players := range g.p {
				x := players.pos%g.m.x
				y := players.pos/g.m.x
				c.Write([]byte(fmt.Sprintf("%v %v", x, y)))
			}
		}
	}
}

func NewGame() *game {
	var g game
	
	g.m = NewMap()
	return &g
}
