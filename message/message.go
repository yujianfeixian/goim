package message

import (
	"encoding/json"
	"goim/client"
)

type Message struct {
	source    *client.Client
	target    *client.Client
	msgHeader map[string]string
	msgBody   interface{}
}

func (thisMessage *Message) Source() *client.Client {
	return thisMessage.source
}

func (thisMessage *Message) SetSource(source *client.Client) {
	thisMessage.source = source
}

func (thisMessage *Message) Target() *client.Client {
	return thisMessage.target
}

func (thisMessage *Message) SetTarget(target *client.Client) {
	thisMessage.target = target
}

func (thisMessage *Message) MsgHeader() map[string]string {
	return thisMessage.msgHeader
}

func (thisMessage *Message) SetMsgHeader(msgHeader map[string]string) {
	thisMessage.msgHeader = msgHeader
}

func (thisMessage *Message) MsgBody() interface{} {
	return thisMessage.msgBody
}

func (thisMessage *Message) SetMsgBody(msgBody interface{}) {
	thisMessage.msgBody = msgBody
}

func NewMessage(source *client.Client, target *client.Client, msgHeader map[string]string, msgBody interface{}) *Message {
	return &Message{source: source, target: target, msgHeader: msgHeader, msgBody: msgBody}
}

func (thisMessage *Message) JsonMessage() string {
	value, err := json.Marshal(thisMessage)
	if err != nil {
		return ""
	}
	return string(value)
}
