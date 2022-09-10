package ws

import "github.com/gofiber/websocket/v2"

type Pool struct {
	Clients *[]websocket.Conn
}

func (pool Pool) addClient(client *websocket.Conn) {

}
