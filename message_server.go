package main

import (
	"fmt"
	"sync"
)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgch  chan Message
	quitch chan struct{}
}

func (s *Server) StartAndListen(wg *sync.WaitGroup) {
free:
	for {
		select {
		case msg := <-s.msgch:
			fmt.Printf("Received message from: %s payload %s\n", msg.From, msg.Payload)
			wg.Done()
		case <-s.quitch:
			break free
		}
	}
}
func sendMessageToServer(msgch chan Message, from string, payload string) {
	msg := Message{
		From:    from,
		Payload: payload,
	}

	msgch <- msg
}

func shutDownServer(quitch chan struct{}) {
	close(quitch)
}
