package onlineContext

import (
	"goim/client"
	"goim/message"
	"net"
	"sync"
	"time"
)

type OnlineContext struct {
	onlineMap  map[string]*client.Client
	serverChan chan string
	mapLock    sync.RWMutex
}

func NewOnlineContext() *OnlineContext {
	return &OnlineContext{
		onlineMap:  map[string]*client.Client{},
		serverChan: make(chan string),
	}
}
func (thisOnlineContext *OnlineContext) BroadCast(s string) {
	for _, v := range thisOnlineContext.onlineMap {
		v.C() <- s
	}
}
func (thisOnlineContext *OnlineContext) RouteMessage(message *message.Message) {
	source := message.Source()
	target := message.Target()
	if target == "all" {
		thisOnlineContext.BroadCast(source + " say: " + message.MsgBody().(string))
	} else {
		thisOnlineContext.onlineMap[target].C() <- source + " say: " + message.MsgBody().(string)
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
		err := c.Conn().Close()
		if err != nil {
			return
		}
		delete(thisOnlineContext.onlineMap, clientName)
	}
}
func (thisOnlineContext *OnlineContext) clearAll() {
	defer thisOnlineContext.mapLock.Unlock()
	for thisOnlineContext.mapLock.TryLock() {
		thisOnlineContext.onlineMap = map[string]*client.Client{}
	}

}

func (thisOnlineContext *OnlineContext) RenameClient(name string, newName string) {
	defer thisOnlineContext.mapLock.Unlock()
	for thisOnlineContext.mapLock.TryLock() {
		thisOnlineContext.onlineMap[name].SetName(newName)
		thisOnlineContext.onlineMap[newName] = thisOnlineContext.onlineMap[name]

		delete(thisOnlineContext.onlineMap, name)
	}
}
