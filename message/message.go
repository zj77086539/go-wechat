package message

import (
	wechat "github.com/zj77086539/go-wechat"
)

// Message 消息
type Message struct {
	client wechat.Clienter
}

// NewMessage ...
func NewMessage(c wechat.Clienter) *Message {
	return &Message{
		client: c,
	}
}
