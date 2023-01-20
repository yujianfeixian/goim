package main

import "encoding/json"

type Message struct {
	source Client
	data   string
}

func (thisMessage *Message) jsonMessage() string {
	value, err := json.Marshal(thisMessage)
	if err != nil {
		return ""
	}
	return string(value)
}
