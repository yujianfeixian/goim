package client

import (
	"fmt"
	"net"
	"time"
)

type Client struct {
	name           string
	addr           string
	c              chan string
	conn           net.Conn
	onLineTime     time.Time
	lastActiveTime time.Time
}

func NewClient(name string, addr string, conn net.Conn) *Client {

	client := &Client{
		name: name,
		addr: addr,
		c:    make(chan string),
		conn: conn,
	}
	go client.listenMessage()
	return client
}
func (thisClient *Client) listenMessage() {
	for {
		if thisClient.c == nil {
			break
		}
		msg := <-thisClient.c
		_, err := thisClient.conn.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Println("write failed, err:", err, "client:", thisClient.name)
			return
		}
	}
}

func (thisClient *Client) Name() string {
	return thisClient.name
}

func (thisClient *Client) SetName(name string) {
	thisClient.name = name
}

func (thisClient *Client) Addr() string {
	return thisClient.addr
}

func (thisClient *Client) SetAddr(addr string) {
	thisClient.addr = addr
}

func (thisClient *Client) C() chan string {
	return thisClient.c
}

func (thisClient *Client) SetC(c chan string) {
	thisClient.c = c
}

func (thisClient *Client) Conn() net.Conn {
	return thisClient.conn
}

func (thisClient *Client) SetConn(conn net.Conn) {
	thisClient.conn = conn
}

func (thisClient *Client) OnLineTime() time.Time {
	return thisClient.onLineTime
}

func (thisClient *Client) SetOnLineTime(onLineTime time.Time) {
	thisClient.onLineTime = onLineTime
}

func (thisClient *Client) LastActiveTime() time.Time {
	return thisClient.lastActiveTime
}

func (thisClient *Client) SetLastActiveTime(lastActiveTime time.Time) {
	thisClient.lastActiveTime = lastActiveTime
}
