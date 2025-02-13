package websocket

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct{
	Conn *websocket.Conn
	Hub *Hub
	mu sync.Mutex
}

type Message struct{
	Type int `json:"type"`
	Payload string `json:"payload"`
}

// // this is for chat
// payload:{
// 	type:"chat"
// 	message: "jfakfiejkla"
// 	userId: "fjakfjia"
// }

// // this is for canvas
// payload:{
// 	type:"canvas"
// 	shapeType:{
// 		type:"circle"
// 		x:9
// 		y:0
// 		centerX:89
// 		centerY:78
// 	}
// }

// here read is Client type , the messages will Client type and all the operation will according to this CLient
func (c *Client) Read(){
	defer func(){
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for{
		msgType, msg, err := c.Conn.ReadMessage()
		if err!=nil{
			fmt.Println(err)
			return
		}
		m:=Message{
			Type: msgType,
			Payload: string(msg),
		}
		c.Hub.Broadcast<-m
		fmt.Println("message received: ", m)
	}
}