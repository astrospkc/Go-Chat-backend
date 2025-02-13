package websocket

import "fmt"

type Hub struct{
	Register  chan *Client
	Unregister chan *Client
	Clients   map[*Client]bool
	Broadcast  chan Message
}


func NewHUb() *Hub{
	return &Hub{
		Register: make(chan *Client),
		Unregister: make(chan *Client),
		Clients: make(map[*Client]bool),
		Broadcast: make(chan Message),
	}
}

func (h *Hub) Start(){

	for{
		select{
		case client := <-h.Register:
			// setting it true to make it confirm that the connection is established
			h.Clients[client] = true
			fmt.Println("total connections : ", len(h.Clients))
			// whenever the cient join the chat
			for k,_:=range h.Clients{
				fmt.Println(k)
				k.Conn.WriteJSON(Message{
					Type:1,
					Payload : "New user joined",
				})
			}
		
		case client := <-h.Unregister:
			delete(h.Clients, client)
			fmt.Println("total connections: ", len(h.Clients))
			// whenever the client leave the chat
			for k,_:=range h.Clients{
				fmt.Println(k)
				k.Conn.WriteJSON(Message{
					Type:1,
					Payload: "User has left",
				})
			}
		
		case msg := <-h.Broadcast:
			fmt.Println("broadcasting a message")
			for k,_:=range h.Clients{
				if err:=k.Conn.WriteJSON(msg); err!=nil{
					fmt.Println(err)
					return
				}
			}
		
		}
	}

}