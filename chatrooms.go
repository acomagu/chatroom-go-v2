package chatrooms

import ()

type Message string

// Room type keeps the connections to user.
type Room struct {
	In  chan Message
	Out chan []Message
}

// Chatroom has all functions and channels to be exported from this package.
type Chatrooms struct {
	Entry Room
	topics []Topic
}

func New(topics []Topic) Chatrooms {
	chatrooms := Chatrooms{
		Entry: Room{
			In: make(chan Message),
			Out: make(chan []Message),
		},
		topics: topics,
	}
	go chatrooms.talk(chatrooms.Entry)
	return chatrooms
}
