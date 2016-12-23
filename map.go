package main

import (
	"os"
	"log"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

const (
	bsize = 20
)

type maps struct {
	x, y int
	layout []byte
}

func drawMap(x, y int, b *bufio.Reader) []byte {
	m := make([]byte, x*y)
	for c := 0; c < y; c++ {
		row, _ := b.ReadBytes('\n')
		copy(m[c*x:], row[:x])
	}
	fmt.Println(string(m))
	return m
}

func getInfo(b *bufio.Reader) (int, int) {
	str, err := b.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return -1, -1
	}
	str = strings.Trim(str, "\n")
	xy := strings.Split(str, ",")	
	x, err := strconv.Atoi(xy[0])
	if err != nil {
		fmt.Println(err)
		return -1, -1
	}
	y, err := strconv.Atoi(xy[1])
	if err != nil {
		fmt.Println(err)
		return -1, -1
	}	
	return x, y
}

func (m *maps) isObstacle(pos int) bool {
	xy := pos/(bsize*m.x)/bsize* m.x + pos%(bsize*m.x)/bsize
	if m.layout[xy] == 'x' {
		return true
	}
	return false
}

func NewMap() maps {
	var m maps
	
	f, err := os.Open("map/map1")
	b := bufio.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	m.x, m.y = getInfo(b)
	m.layout = drawMap(m.x, m.y, b)
	return m
}
