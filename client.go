package main

import (
	"fmt"
	"net"
	"time"
)

type Client struct {
	Name           string
	Addr           string
	c              chan string
	conn           net.Conn
	onLineTime     time.Time
	lastActiveTime time.Time
}

func newClient(conn net.Conn) *Client {
	client := &Client{
		Name: conn.RemoteAddr().String(),
		Addr: conn.RemoteAddr().String(),
		c:    make(chan string),
		conn: conn,
	}
	go client.listenMessage()
	return client
}
func (thisClient *Client) online() {

}
func (thisClient *Client) listenMessage() {
	for {
		if thisClient.c == nil {
			break
		}
		msg := <-thisClient.c
		_, err := thisClient.conn.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Println("write failed, err:", err, "client:", thisClient.Name)
			return
		}
	}
}
