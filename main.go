package main

import (
	"fmt"
	"goChat/websocket"
	"net/http"
)

func serveWs(hub *websocket.Hub, w http.ResponseWriter, r *http.Request){
	conn, err:=websocket.Upgrade(w,r)
	if err!=nil{
		fmt.Println(err)
		return
	}

	client := &websocket.Client{
		Conn :conn,
		Hub : hub,
	}
	hub.Register<-client
	client.Read()
}


func setUpRoutes(){
	hub := websocket.NewHUb()
	go hub.Start()
	http.HandleFunc("/ws",func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
}

func main(){
	setUpRoutes()
	fmt.Println("hi there")
	http.ListenAndServe(":8000",nil)
}