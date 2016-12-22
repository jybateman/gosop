package main

import (
	"fmt"
	"net/http"
	"html/template"
	
	"golang.org/x/net/websocket"
)

func index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	tpl.Execute(w, nil)
}

func JSHandler(c *websocket.Conn) {
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/ws", websocket.Handler(StartGame))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
        http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.ListenAndServe(":4243", nil)
}
