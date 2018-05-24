package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/go-chi/chi"
	"github.com/nlopes/slack"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

var (
	InvalidToken          = errors.New("Invalid Verification Token")
	InvalidArgumentNumber = errors.New("Invalid Argument Number")
	InvalidPayload        = errors.New("Invalid payload")
)

func ResponseError(ctx context.Context, w http.ResponseWriter, err error) {
	log.Errorf(ctx, err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func interactiveHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseError(ctx, w, InvalidPayload)
		return
	}

	jsonStr, err := url.QueryUnescape(string(buf)[8:])
	if err != nil {
		ResponseError(ctx, w, InvalidPayload)
		return
	}
	log.Infof(ctx, jsonStr)

	var message slack.AttachmentActionCallback
	if err := json.Unmarshal([]byte(jsonStr), &message); err != nil {
		ResponseError(ctx, w, InvalidPayload)
		return
	}

	if message.Token != os.Getenv("SLACK_VERIFICATION_TOKEN") {
		ResponseError(ctx, w, InvalidToken)
		return
	}

	var msg slack.Msg

	switch message.Actions[0].Value {
	case "deploy":
		msg = slack.Msg{
			Attachments: []slack.Attachment{
				{
					Text:       "Deploy started! :tada:",
					Color:      "good",
					CallbackID: "deployment",
				},
			},
		}
	case "cancel":
		msg = slack.Msg{
			Attachments: []slack.Attachment{
				{
					Text:       "Deploy cancelled",
					Color:      "danger",
					CallbackID: "deployment",
				},
			},
		}
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(&msg)
}

func slashHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	r.ParseForm()
	if r.PostForm.Get("token") != os.Getenv("SLACK_VERIFICATION_TOKEN") {
		ResponseError(ctx, w, InvalidToken)
		return
	}

	message := slack.Msg{
		Attachments: []slack.Attachment{
			{
				Text:       "Do you want to deploy to production server?",
				Color:      "warning",
				CallbackID: "deployment",
				Actions: []slack.AttachmentAction{
					{
						Name:  "deployOrNot",
						Type:  "button",
						Text:  "Deploy!",
						Value: "deploy",
					},
					{
						Name:  "deployOrNot",
						Type:  "button",
						Text:  "Cancel",
						Value: "cancel",
						Style: "danger",
					},
				},
			},
		},
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(&message)
}

func main() {
	r := chi.NewRouter()
	r.Post("/slash", slashHandler)
	r.Post("/interactive", interactiveHandler)
	http.Handle("/", r)

	appengine.Main()
}
