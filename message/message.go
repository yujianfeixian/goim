package message

type Message struct {
	source    string
	target    string
	msgHeader map[string]string
	msgBody   interface{}
}

func (m *Message) Source() string {
	return m.source
}

func (m *Message) SetSource(source string) {
	m.source = source
}

func (m *Message) Target() string {
	return m.target
}

func (m *Message) SetTarget(target string) {
	m.target = target
}

func (m *Message) MsgHeader() map[string]string {
	return m.msgHeader
}

func (m *Message) SetMsgHeader(msgHeader map[string]string) {
	m.msgHeader = msgHeader
}

func (m *Message) MsgBody() interface{} {
	return m.msgBody
}

func (m *Message) SetMsgBody(msgBody interface{}) {
	m.msgBody = msgBody
}

func NewMessage(source string, target string, msgHeader map[string]string, msgBody interface{}) *Message {
	return &Message{source: source, target: target, msgHeader: msgHeader, msgBody: msgBody}
}
