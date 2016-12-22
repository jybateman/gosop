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

	g.m = NewMap()
	g.p = NewPlayer(&g.m)
	fmt.Println(g.p)
	c.Write([]byte(fmt.Sprintf("%v", g.m)))
	for {
		c.Write([]byte(fmt.Sprintf("%v", g.p.pos)))
		time.Sleep(time.Millisecond*17)
		g.p.pos += 10
	}
}
