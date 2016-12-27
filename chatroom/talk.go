package chatroom

import (
	"fmt"
)

// TopicChan includes chatroom channel, the channel pass the returned value from topic.
type TopicChan struct {
	Chatroom roomInternal
	Return   chan bool
}

// A DidTalk means whether the Topic talks with user. A Topic function must return this type value.
type DidTalk bool

// Topic type is the function express a bunch of flow in chattting. Pass slice of this to New(), and the function called them in order. If one of them returns true, the loop breaks.
type Topic func(Room) DidTalk

func (cr Chatroom) talk(chatroom roomInternal) {
	topicChans := []TopicChan{}
	for _, topic := range cr.topics {
		topicChan := loopTopic(topic, chatroom)
		topicChans = append(topicChans, topicChan)
	}
	middleChatroom, clearPool, broadcastPool := poolMessages(chatroom)
	changeDestTopicTo := distributeMessage(middleChatroom)
	go controller(topicChans, changeDestTopicTo, broadcastPool, clearPool)
}

func controller(topicChans []TopicChan, changeDestTopicTo chan roomInternal, broadcastPool chan bool, clearPool chan bool) {
	for {
		for i, topicChan := range topicChans {
			changeDestTopicTo <- topicChan.Chatroom
			if i > 0 { // for the start time.
				broadcastPool <- true
			}
			didTalk := <-topicChan.Return
			if didTalk {
				clearPool <- true
				break
			}
		}
		clearPool <- true
	}
}

// This pipe stores messages from user with flowing next Chatroom(middleChatroom). And this provides functions, clearPool and broadcastPool. This is used in controller().
func poolMessages(chatroom roomInternal) (roomInternal, chan bool, chan bool) {
	middleChatroom := roomInternal{
		in:  make(chan interface{}),
		out: chatroom.out,
	}
	clearPool := make(chan bool)
	broadcastPool := make(chan bool)

	go func(chatroom roomInternal, middleChatroom roomInternal, clearPool <-chan bool, broadcastPool <-chan bool) {
		var pool []interface{}
		for {
			select {
			case message := <-chatroom.in:
				pool = append(pool, message)
				middleChatroom.in <- message

			case <-clearPool:
				pool = pool[:0]

			case <-broadcastPool:
				for _, message := range pool {
					middleChatroom.in <- message
				}
			}
		}
	}(chatroom, middleChatroom, clearPool, broadcastPool)

	return middleChatroom, clearPool, broadcastPool
}

// distributeMessage pass message from chatroom to chatroom. The chatroom of destination will change as needed, changed by value of channel, changeDestTopicTo.
func distributeMessage(middleChatroom roomInternal) chan roomInternal {
	changeDestTopicTo := make(chan roomInternal)

	go func(middleChatroom roomInternal, changeDestTopicTo <-chan roomInternal) {
		var dest roomInternal
		dest = <-changeDestTopicTo
		for {
			select {
			case message := <-middleChatroom.in:
				if dest == (roomInternal{}) {
					fmt.Println("Error: the destination chatroom is not set.")
					break
				}
				dest.in <- message

			case _dest := <-changeDestTopicTo:
				dest = _dest
			}
		}
	}(middleChatroom, changeDestTopicTo)

	return changeDestTopicTo
}

// loopTopic just loops topic.
func loopTopic(topic Topic, chatroom roomInternal) TopicChan {
	topicChan := TopicChan{
		Chatroom: roomInternal{
			in:  make(chan interface{}),
			out: chatroom.out,
		},
		Return: make(chan bool),
	}

	go func(topic Topic, topicChan TopicChan) {
		for {
			topicChan.Return <- bool(topic(topicChan.Chatroom))
		}
	}(topic, topicChan)

	return topicChan
}
