package chatroom

// A Chatroom has all functions and channels to be exported from this package.
type Chatroom struct {
	topics []Topic
	Room
}

// A Room has functions to wait or send messages with user. This is passed to Topic function as argument.
type Room struct {
	In  chan interface{}
	Out chan interface{}
}

// Flush inputs the value to be passed to Topic.
func (cr Chatroom) Flush(v interface{}) {
	cr.In <- v
}

// WaitSentMsg waits and returns the value passed to all of Room#Send. Used to send messages to user (through chat service).
func (cr Chatroom) WaitSentMsg() interface{} {
	text := <-cr.Out
	return text
}

// WaitSentTextMsg waits and returns the string value. It ignores the others.
func (cr Chatroom) WaitSentTextMsg() string {
	for {
		if str, ok := cr.WaitSentMsg().(string); ok {
			return str
		}
	}
}

func (room Room) WaitMsg() interface{} {
	return <-room.In
}

func (room Room) WaitTextMsg() string {
	for {
		if str, ok := (<-room.In).(string); ok {
			return str
		}
	}
}

func (room Room) Send(v interface{}) {
	room.Out <- v
}

// New creates and initialize a Chatroom. This also starts a go-routine to pass messages to Topics.
func New(topics []Topic) Chatroom {
	room := Room{
		In:  make(chan interface{}),
		Out: make(chan interface{}),
	}
	chatroom := Chatroom{
		topics: topics,
		Room:   room,
	}
	go chatroom.talk(room)
	return chatroom
}
