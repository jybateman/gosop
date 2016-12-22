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
	for {
		c.SetReadDeadline(time.Now().Add(time.Millisecond*33))
		c.Read(b)
		g.p.ChangeDir(string(b))
		c.Write([]byte(fmt.Sprintf("%v", g.p.pos)))
		g.p.MovePlayer(g.m.x)
	}
}
