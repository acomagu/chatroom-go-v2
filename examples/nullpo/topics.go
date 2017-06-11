package main

import (
	"github.com/acomagu/chatroom-go-v2/chatroom"
)

var topics = []chatroom.Topic{responseToNullpo, responseToSegfo, responseToAny}

func responseToNullpo(room chatroom.Room) chatroom.DidTalk {
	if msg, ok := (<-room.In).(string); ok && msg == "Nullpo" {
		postToSlack("Ga")
		return true
	}
	return false
}

func responseToSegfo(room chatroom.Room) chatroom.DidTalk {
	if msg, ok := (<-room.In).(string); ok && msg == "Segfo" {
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
