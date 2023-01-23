package main

import (
	"sync"
	"time"
)

type OnlineContext struct {
	OnlineMap map[string]*Client
	mapLock   sync.RWMutex
}

func (thisOnlineContext *OnlineContext) addNewClient(client *Client) {
	remoteAddr := client.conn.RemoteAddr()
	onlineTime := time.Now()
	key := remoteAddr.String() + onlineTime.String()
	client.onLineTime = onlineTime
	client.lastActiveTime = onlineTime
	defer thisOnlineContext.mapLock.Unlock()
	for thisOnlineContext.mapLock.TryLock() {
		thisOnlineContext.OnlineMap[key] = client
	}
}
