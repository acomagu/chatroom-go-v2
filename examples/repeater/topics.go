package main

import (
	"github.com/acomagu/chatroom-go/chatroom"
)

func responseToNullpo(room chatrooms.Room) bool {
	a := room.WaitTextMsg()
	if a == "Nullpo" {
		postToSlack("Ga")
		return true
	}
	return false
}

func responseToSegfo(room chatrooms.Room) bool {
	a := room.WaitTextMsg()
	if a == "Segfo" {
		postToSlack("Na")
		return true
	}
	return false
}

func responseToAny(room chatrooms.Room) bool {
	_ = room.WaitTextMsg()
	postToSlack("None")
	return true
}
