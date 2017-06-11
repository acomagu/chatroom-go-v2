package main

import (
	"github.com/acomagu/chatroom-go-v2/chatroom"
)

var topics = []chatroom.Topic{ppap}

func ppap(room chatroom.Room) chatroom.DidTalk {
	if msg, ok := (<-room.In); ok && msg != "PPAP" {
		return false
	}
	room.Out <- "I have a pen."
	room.Out <- "I have a/an ..."
	apple, ok := (<-room.In).(string)
	if !ok {
		return true
	}
	room.Out <- "Ah!"
	room.Out <- apple + "Pen!"

	room.Out <- "I have a pen."
	room.Out <- "I have a/an ..."
	pineapple, ok := (<-room.In).(string)
	if !ok {
		return true
	}
	room.Out <- "Ah!"
	room.Out <- pineapple + "Pen!"

	room.Out <- apple + "Pen,"
	room.Out <- pineapple + "Pen,"
	room.Out <- "Ah!"

	room.Out <- "Pen" + pineapple + apple + "Pen!"
	return true
}
