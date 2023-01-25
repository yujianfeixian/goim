package server

import (
	"fmt"
	messages "goim/message"
	"goim/onlineContext"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip            string
	Port          int
	OnlineContext *onlineContext.OnlineContext
	OmLock        sync.RWMutex
}

// NewServer /**
func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:            ip,
		Port:          port,
		OnlineContext: onlineContext.NewOnlineContext(),
	}
}
func (thisServer *Server) Start() {
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
	defer func(conn *net.TCPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("close conn failed or already closed, err:", err)
		}
	}(conn)
	clientName := conn.RemoteAddr().String() + "_" + time.Now().Format("2006-01-02 15:04:05")
	addr := conn.RemoteAddr().String()
	isActive := make(chan bool)
	thisServer.OmLock.Lock()
	thisServer.OnlineContext.AddNewClient(clientName, addr, conn)
	thisServer.OmLock.Unlock()
	go func() {
		for {
			buf := make([]byte, 512)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("read from conn failed, err:", err)
				return

			}
			fmt.Println("read from client, data:", string(buf[:n]))
			isActive <- true
		}
	}()
	for {
		select {
		case <-isActive:
		case <-time.After(time.Minute * 10):
			for thisServer.OmLock.TryLock() {
				thisServer.OnlineContext.RemoveClient(clientName)
				isActive <- false
			}
			thisServer.OmLock.Unlock()
		}
	}
}
func (thisServer *Server) broadCast(msg string) {
	thisServer.OmLock.Lock()
	defer thisServer.OmLock.Unlock()
	message := messages.NewMessage(nil, nil, nil, msg)
	thisServer.OnlineContext.ServerChan() <- message.JsonMessage()
}
func (thisServer *Server) processMessage() {
	//TODO 服务端路由消息
}
