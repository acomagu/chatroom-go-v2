package main

import (
	"github.com/acomagu/chatroom-go/chatroom"
)

var topics = []chatroom.Topic{responseToNullpo, responseToSegfo, responseToAny}

func responseToNullpo(room chatroom.Room) chatroom.DidTalk {
	a := room.WaitTextMsg()
	if a == "Nullpo" {
		postToSlack("Ga")
		return true
	}
	return false
}

func responseToSegfo(room chatroom.Room) chatroom.DidTalk {
	a := room.WaitTextMsg()
	if a == "Segfo" {
		postToSlack("Na")
		return true
	}
	return false
}

func responseToAny(room chatroom.Room) chatroom.DidTalk {
	_ = room.WaitTextMsg()
	postToSlack("None")
	return true
}
