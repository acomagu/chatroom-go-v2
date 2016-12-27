package chatroom

import ()

type roomInternal struct {
	in  chan interface{}
	out chan interface{}
}

// A Chatroom has all functions and channels to be exported from this package.
type Chatroom struct {
	// Entry RoomEntry
	topics []Topic
	entry roomInternal
}

// A Room has functions to wait or send messages with user. This is passed to Topic function as argument.
type Room interface {
	WaitMsg() interface{}
	WaitTextMsg() string
	Send(interface{})
}

// Flush inputs the value to be passed to Topic.
func (cr Chatroom) Flush(v interface{}) {
	cr.entry.in <- v
}

// WaitSentMsg waits and returns the value passed to all of Room#Send. Used to send messages to user (through chat service).
func (cr Chatroom) WaitSentMsg() interface{} {
	text := <-cr.entry.out
	return text
}

// WaitSentTextMsg waits and returns the TEXT value. It ignores the others.
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

// New creates and initialize a Chatroom. This also starts a go-routine to pass messages to Topics.
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
