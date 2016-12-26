package main

import (
	"fmt"
	"time"
	
	"golang.org/x/net/websocket"
)

type game struct {
	m maps
	p player
}

func StartGame(c *websocket.Conn) {
	var g game
	
	b := make([]byte, 2)
	g.m = NewMap()
	g.p = NewPlayer(&g.m)
	fmt.Println(g.p)
	c.Write([]byte(fmt.Sprintf("%v", g.m)))
	dl := time.Now().Add(time.Millisecond*100)
	for {
		c.SetReadDeadline(dl)
		c.Read(b)
		if dl.Before(time.Now()) {
			dl = time.Now().Add(time.Millisecond*100)
			g.p.ChangeDir(string(b))
			g.p.MovePlayer(&g.m)
			x := g.p.pos%g.m.x
			y := g.p.pos/g.m.x
			c.Write([]byte(fmt.Sprintf("%v %v", x, y)))
		}
	}
}
