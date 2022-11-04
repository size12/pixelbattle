package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"reflect"
	"sync"
)

type MessageReceive struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type MessageSend struct {
	Event string `json:"event"`
	X     int    `json:"x"`
	Y     int    `json:"y"`
}

type Point struct {
	X int
	Y int
}

type Field []Point

var field = Field{}
var command = MessageSend{}
var mu sync.Mutex

func main() {
	server := http.NewServeMux()
	server.Handle("/websocket", websocket.Handler(wsHandler))
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		log.Fatal(err)
	}
}

func wsHandler(ws *websocket.Conn) {
	//fmt.Println("new connection")
	go wsReceiver(ws)
	go wsWriter(ws)
	c := make(chan int)
	<-c
}

func indexOf(element Point, data []Point) int {
	for k, v := range data {
		if reflect.DeepEqual(v, element) {
			return k
		}
	}
	return -1 //not found.
}

func wsReceiver(ws *websocket.Conn) {
	var in []byte
	mess := MessageReceive{}

	for {
		if err := websocket.Message.Receive(ws, &in); err != nil {
			return
		}
		err := json.Unmarshal(in, &mess)
		if err != nil {
			fmt.Println(err)
			return
		}

		point := Point{mess.X, mess.Y}

		mu.Lock()
		index := indexOf(point, field)
		if index != -1 {
			field = append(field[:index], field[index+1:]...)
			command = MessageSend{Event: "clear", X: point.X, Y: point.Y}
		} else {
			field = append(field, point)
			command = MessageSend{Event: "fill", X: point.X, Y: point.Y}
		}
		mu.Unlock()

	}

}

func wsWriter(ws *websocket.Conn) {
	var lastCommand MessageSend

	for _, point := range field {
		sendData, err := json.Marshal(MessageSend{Event: "fill", X: point.X, Y: point.Y})
		if err != nil {
			fmt.Println(err)
			return
		}
		websocket.Message.Send(ws, string(sendData))
	}

	for {
		if !reflect.DeepEqual(lastCommand, command) {
			mu.Lock()
			sendData, err := json.Marshal(command)
			mu.Unlock()
			if err != nil {
				fmt.Println(err)
				return
			}
			websocket.Message.Send(ws, string(sendData))
			lastCommand = command
		}
	}
}
