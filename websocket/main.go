package main

import (
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func echoHandler(ws *websocket.Conn) {

	type T struct {
		Msg   string
		Count float64
	}

	// receive JSON type T
	var data T
	websocket.JSON.Receive(ws, &data)

	log.Printf("data=%#v\n", data)

	// send JSON type T
	websocket.JSON.Send(ws, data)
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/echo",
		func(w http.ResponseWriter, req *http.Request) {
			s := websocket.Server{Handler: websocket.Handler(echoHandler)}
			s.ServeHTTP(w, req)
		})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}
