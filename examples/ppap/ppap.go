package main

import (
	"net/http"

	"github.com/acomagu/chatroom-go/chatroom"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	bot, _ = linebot.New("<channel secret>", "<channel access token>")
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8000", nil)
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	events, _ := bot.ParseRequest(req)
	for _, event := range events {
		if event != nil && event.Type == linebot.EventTypeMessage {
			if msg, ok := event.Message.(*linebot.TextMessage); ok {
				handleTextMsg(msg.Text, event.Source.UserID)
			}
		}
	}
}

var crs = make(map[string]chatroom.Chatroom)

func handleTextMsg(text string, userID string) {
	cr, ok := crs[userID]

	if !ok {
		cr = chatroom.New(topics)
		crs[userID] = cr
		go sender(userID, cr)
	}
	cr.Flush(text)
}

func sender(userID string, cr chatroom.Chatroom) {
	for {
		text := cr.WaitSentTextMsg()
		bot.PushMessage(userID, linebot.NewTextMessage(text)).Do()
	}
}
