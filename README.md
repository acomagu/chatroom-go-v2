Chatroom
========

Create readable chatbot quickly with Go.

## Description
A small library for chatbot for go.

First, to understand properly this library, watch this:

[![PPAP](http://img.youtube.com/vi/0E00Zuayv9Q/0.jpg)](http://www.youtube.com/watch?v=0E00Zuayv9Q)

You can write the awesome video as code like below:

![Source Code Sample](https://github.com/acomagu/chatroom-go/raw/master/img/Desktop.png)

## Oh, you're already full with PPAP?

OK, explain the details.

And do only this:
- Call function(Topic)
- Pass messages to topic through pipes.

As the above image, this library will omit the state managements on actual code. It can also be said "wrapper of simple channel pipelines".

But this library DON'T do below:
- Manage states of each users
- Communicate with chat service, Facebook, LINE and the like.

So, if you must keep datas over Topic or by users, you must write a bit more code. Like this:

```Go
cr, ok := crs[userID]

if !ok {
  cr = chatroom.New(topics)
  crs[userID] = cr

  ...

```

Create instance for each user.

## You're interested in?

Read Reference and Examples!

[https://github.com/acomagu/chatroom-go/tree/master/examples](https://github.com/acomagu/chatroom-go/tree/master/examples)

## Requirement
- Golang

## Licence

[MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)

## Author

[acomagu](https://github.com/acomagu)
