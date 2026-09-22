package main

import (
	"fmt"
	"time"
)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgch  chan Message
	quitch chan struct{}
}

func (s *Server) StartAndListen() {
free:
	for {
		select {
		// block here until someone is sending a message to the server
		case msg := <-s.msgch:
			fmt.Printf("received a message from: %s, payload: %s\n", msg.From, msg.Payload)
		case <-s.quitch:
			fmt.Println("the server is doing a gracefull shutdown")
			break free
		default:
		}
	}

	fmt.Println("The server is shutdown")
}

func sendMessageToServer(msgch chan Message, payload string) {
	msg := Message{
		From:    "Anthony",
		Payload: payload,
	}

	msgch <- msg
}

func gracefullQuitServer(quitch chan struct{}) {
	close(quitch)
}

func main() {
	s := &Server{
		msgch:  make(chan Message),
		quitch: make(chan struct{}),
	}

	go s.StartAndListen()

	go func() {
		time.Sleep(time.Second * 2)
		sendMessageToServer(s.msgch, "How's it going?")
	}()

	go func() {
		time.Sleep(time.Second * 4)
		gracefullQuitServer(s.quitch)
	}()

	select {}
}
