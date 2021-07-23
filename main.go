package main

import (
	"context"
)

// WIP just hacking a bit

func main() {

}

// MessageChain is a rolling chain of messages
// Ifm.maxMessages is reached it will roll over when adding new elements
type MessageChain struct {
	maxMessages uint64
	msgCount    uint64
	chain       []Message
}

func New(n uint64) *MessageChain {
	m := MessageChain{}
	m.chain = make([]Message, n)
	m.maxMessages = n
	return &m
}

func (m *MessageChain) Add(msg Message) (index uint64) {
	index = m.msgCount
	if index > m.maxMessages {
		index %= m.maxMessages
	}
	m.chain[index] = msg
	return index
}

// Address is a public key that can be used to encrypt messages
// and serves as the recipients address at the same time
type Address []byte

type Message struct {
	Text     string
	Sender   Address
	Receiver Address
}

func write(ctx context.Context, msg Message) {

}
