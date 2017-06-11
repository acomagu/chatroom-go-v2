Chatroom [![godoc](https://img.shields.io/badge/reference-godoc-blue.svg)](https://godoc.org/github.com/acomagu/chatroom-go/chatroom)
========


__Create readable chatbot quickly with Go.__

## Description
A small library for chatbot for go.

This library do only below:
- Call function(Topic)
- Pass messages to topic through pipes.

This library will omit the state managements on actual code. It can also be said "wrapper of simple channel pipelines".

But this library DON'T do below:
- Manage states of each users
- Communicate with chat service, Facebook, LINE and the like.

So, if you must keep data over Topic or for each users, you must write a bit more code. Generaly it will be creating instances of Chatrom for each user and use closure for `Topic` funcs, or use global variables.

```Go
cr, ok := crs[userID]

if !ok {
  cr = chatroom.New(topics)
  crs[userID] = cr

  ...

```

## What is `Topic`?

`Topic` is a function, logic of a unit of conversation with user.

You write the actual code talking with users, waiting user's reaction and replying to them.

```Go
func responseToNullpo(room chatrooms.Room) chatroom.DidTalk {
	a := <-room.In
	if a == "Nullpo" {
		postToSlack("Ga")
		return true
	}
	return false
}
```

The return value is whether talk with user or not. If it returns `false`, Chatroom will call other Topic, but if it's `true`, Chatroom stops the calling.

## How to connect LINE or Facebook Messenger with this library

To pass the messages from user to Chatroom library, use `Chatroom#In`.

Example on Slack bot:

```Go
	cr := chatroom.New(topics)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		if getSentUserName(body) == "slackbot" {
			return
		}

		// Pass the received message to Chatroom.
		cr.In <- getReceivedMessage(body)

	})
```

(The whole code: [examples/nullpo/nullpo.go](https://github.com/acomagu/chatroom-go/blob/master/examples/nullpo/nullpo.go))


And you can use it to send to user. Call `Room#Out`, and you can receive it by `Chatroom#Out`.

For instance,

```Go
func sender(userID string, cr chatroom.Chatroom) {
	for {
		text := <-cr.Out
		bot.PushMessage(userID, linebot.NewTextMessage(text)).Do()
	}
}
```

(The whole code: [examples/ppap/ppap.go](https://github.com/acomagu/chatroom-go/blob/master/examples/ppap/ppap.go))

You can exclude UserID from Topic functions by using this feature.

## Are you interested in?

Read Reference and Examples!

[chatroom - GoDoc](https://godoc.org/github.com/acomagu/chatroom-go-v2/chatroom)

[chatroom-go/examples at master · acomagu/chatroom-go](https://github.com/acomagu/chatroom-go-v2/tree/master/examples)

## Requirement
- Golang

## Licence

[MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)

## Author

[acomagu](https://github.com/acomagu)
