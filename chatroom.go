package chatroom

import ()

type Message string

type roomInternal struct {
	in  chan interface{}
	out chan interface{}
}

// Chatroom has all functions and channels to be exported from this package.
type Chatroom struct {
	// Entry RoomEntry
	topics []Topic
	entry roomInternal
}

type Room interface {
	WaitMsg() interface{}
	WaitTextMsg() string
	Send(interface{})
}

func (cr Chatroom) Flush(text string) {
	cr.entry.in <- text
}

func (cr Chatroom) WaitSentMsg() interface{} {
	text := <-cr.entry.out
	return text
}

func (cr Chatroom) WaitSentTextMsg() string {
	for {
		if str, ok := cr.WaitSentMsg().(string); ok {
			return str
		}
	}
}

func (room roomInternal) WaitMsg() interface{} {
	return <-room.in
}

func (room roomInternal) WaitTextMsg() string {
	for {
		if str, ok := (<-room.in).(string); ok {
			return str
		}
	}
}

func (room roomInternal) Send(v interface{}) {
	room.out <- v
}

func New(topics []Topic) Chatroom {
	_room := roomInternal{
		in: make(chan interface{}),
		out: make(chan interface{}),
	}
	chatroom := Chatroom{
		// Entry: newRoomEntry(_room),
		topics: topics,
		entry: _room,
	}
	go chatroom.talk(_room)
	return chatroom
}
