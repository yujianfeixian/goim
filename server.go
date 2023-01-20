package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip          string
	Port        int
	OnlineMap   map[string]*Client
	OmLock      sync.RWMutex
	MessageChan chan string
}

/**
 * A tcp server
 */
func newServer(ip string, port int) *Server {
	return &Server{
		Ip:          ip,
		Port:        port,
		OnlineMap:   map[string]*Client{},
		MessageChan: make(chan string),
	}
}
func (thisServer *Server) start() {
	tcpListener, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.ParseIP(thisServer.Ip),
		Port: thisServer.Port,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer tcpListener.Close()
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go thisServer.process(tcpConn)
	}

}

func (thisServer *Server) process(conn *net.TCPConn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read from conn failed, err:", err)
			return
		}
		fmt.Println("read from client, data:", string(buf[:n]))
	}
}
func (thisServer *Server) broadCast(client Client, msg string) {
	thisServer.OmLock.Lock()
	defer thisServer.OmLock.Unlock()
	message := Message{
		source: client,
		data:   msg,
	}
	thisServer.MessageChan <- message.jsonMessage()
}
