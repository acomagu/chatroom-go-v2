Chatroom
========

__Create readable chatbot quickly with Go.__

## Description
A small library for chatbot for go.

First, to understand properly this library, watch this:

[![PPAP](http://img.youtube.com/vi/0E00Zuayv9Q/0.jpg)](http://www.youtube.com/watch?v=0E00Zuayv9Q)

You can write the awesome video as code like below:

![Source Code Sample](https://github.com/acomagu/chatroom-go/raw/master/img/Desktop.png)

## Ah! ...Do you need more details?

This library do only below:
- Call function(Topic)
- Pass messages to topic through pipes.

As the above image, this library will omit the state managements on actual code. It can also be said "wrapper of simple channel pipelines".

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

`Topic` is a function, includes a codes like above PPAP image.

You write the actual code talking with users, waiting user's reaction and replying to them.

```Go
func responseToNullpo(room chatrooms.Room) chatroom.DidTalk {
	a := room.WaitTextMsg()
	if a == "Nullpo" {
		postToSlack("Ga")
		return true
	}
	return false
}
```

The return value is whether talk with user or not. If it returns `false`, Chatroom will call other Topic, but if it's `true`, Chatroom stops the calling.

## How can I dance the PPAP on LINE or Facebook?

To pass messages from user to Chatroom library, use `Chatroom#Flush`.

Example on Slack bot:

```Go
	cr := chatroom.New(topics)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		if getSentUserName(body) == "slackbot" {
			return
		}

		// Pass the received message to Chatroom.
		cr.Flush(getReceivedMessage(body))

	})
```

(The whole code: [examples/nullpo/nullpo.go](https://github.com/acomagu/chatroom-go/blob/master/examples/nullpo/nullpo.go))


And you can use it to send to user. Call `Room#Send`, and you can receive it by `Chatroom#WaitSentMsg` or `Chatroom#WaitSentTextMsg`.

On the PPAP code:

```Go
func sender(userID string, cr chatroom.Chatroom) {
	for {
		text := cr.WaitSentTextMsg()
		bot.PushMessage(userID, linebot.NewTextMessage(text)).Do()
	}
}
```

(The whole code: [examples/ppap/ppap.go](https://github.com/acomagu/chatroom-go/blob/master/examples/ppap/ppap.go))

You can exclude UserID from Topic functions by using this feature.

## Are you interested in?

Read Reference and Examples!

[https://github.com/acomagu/chatroom-go/tree/master/examples](https://github.com/acomagu/chatroom-go/tree/master/examples)

## Requirement
- Golang

## Licence

[MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)

## Author

[acomagu](https://github.com/acomagu)
