package onlineContext

import (
	"goim/client"
	"net"
	"sync"
	"time"
)

type OnlineContext struct {
	onlineMap  map[string]*client.Client
	serverChan chan string
	mapLock    sync.RWMutex
}

func (thisOnlineContext *OnlineContext) ServerChan() chan string {
	return thisOnlineContext.serverChan
}

func (thisOnlineContext *OnlineContext) SetServerChan(serverChan chan string) {
	thisOnlineContext.serverChan = serverChan
}

func NewOnlineContext() *OnlineContext {
	return &OnlineContext{
		onlineMap:  map[string]*client.Client{},
		serverChan: make(chan string),
	}
}
func (thisOnlineContext *OnlineContext) AddNewClient(name string, addr string, conn net.Conn) {
	newClient := client.NewClient(name, addr, conn)
	onlineTime := time.Now()
	key := newClient.Name()
	newClient.SetOnLineTime(onlineTime)
	newClient.SetLastActiveTime(onlineTime)
	defer thisOnlineContext.mapLock.Unlock()
	for thisOnlineContext.mapLock.TryLock() {
		thisOnlineContext.onlineMap[key] = newClient
	}
}
func (thisOnlineContext *OnlineContext) RemoveClient(clientName string) {
	defer thisOnlineContext.mapLock.Unlock()
	for thisOnlineContext.mapLock.TryLock() {
		c := thisOnlineContext.onlineMap[clientName]
		close(c.C())
		c.Conn().Close()
		delete(thisOnlineContext.onlineMap, clientName)
	}
}
func (thisOnlineContext *OnlineContext) clearAll() {
	defer thisOnlineContext.mapLock.Unlock()
	for thisOnlineContext.mapLock.TryLock() {
		thisOnlineContext.onlineMap = map[string]*client.Client{}
	}

}
