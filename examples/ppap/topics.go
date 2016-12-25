package main

import (
	"github.com/acomagu/chatroom-go/chatroom"
)

var topics = []chatroom.Topic{ppap}

func ppap(room chatroom.Room) bool {
	if room.WaitTextMsg() != "PPAP" {
		return false
	}
	room.Send("I have a pen.")
	room.Send("I have a/an ...")
	apple := room.WaitTextMsg()
	room.Send("Ah!")
	room.Send(apple + "Pen!")

	room.Send("I have a pen.")
	room.Send("I have a/an ...")
	pineapple := room.WaitTextMsg()
	room.Send("Ah!")
	room.Send(pineapple + "Pen!")

	room.Send(apple + "Pen,")
	room.Send(pineapple + "Pen,")
	room.Send("Ah!")

	room.Send("Pen" + pineapple + apple + "Pen!")
	return true
}
