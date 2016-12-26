// Demo for Multiple Topics

package main

import (
	"io/ioutil"
	"net/url"

	"encoding/json"
	"net/http"
	"github.com/acomagu/chatroom-go/chatroom"
)

const slackIncomingWebhookURL = "<Slack Incoming Webhook URL>"

// Slack struct is used for request and response to Slack service.
type Slack struct {
	Text string `json:"text"`
}

func main() {
	cr := chatroom.New(topics)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		if getSentUserName(body) == "slackbot" {
			return
		}
		// Pass the received message to Chatroom.
		cr.Flush(getReceivedMessage(body))
	})
	http.ListenAndServe(":8000", nil)
}

func postToSlack(text string) {
	jsonStr, _ := json.Marshal(Slack{Text: text})
	http.PostForm(slackIncomingWebhookURL, url.Values{"payload": {string(jsonStr)}})
}

func getReceivedMessage(body []byte) string {
	parsed, _ := url.ParseQuery(string(body))
	return parsed["text"][0]
}

func getSentUserName(body []byte) string {
	parsed, _ := url.ParseQuery(string(body))
	return parsed["user_name"][0]
}
