package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var lametricPushUrl string
var lametricToken string

type JiraWebhookPayload struct {
	Issue struct {
		Key    string
		Fields struct {
			Summary  string
			Assignee struct {
				DisplayName string
			}
		}
	}
}

type NameIndicatorFrame struct {
	Index int    `json:"index"`
	Text  string `json:"text"`
	Icon  string `json:"icon"`
}

type LametricPush struct {
	Frames []NameIndicatorFrame `json:"frames"`
}

func parseJiraWebhook(body io.ReadCloser) (*JiraWebhookPayload, error) {
	decoder := json.NewDecoder(body)
	message := &JiraWebhookPayload{}

	err := decoder.Decode(&message)
	if err != nil {
		return nil, err
	} else if &message.Issue.Key == nil {
		return nil, errors.New("parsing failed")
	}

	log.Println("received payload...", message)

	return message, nil
}

func pushToLametric(text string, icon string) {
	data := &LametricPush{
		Frames: []NameIndicatorFrame{
			{
				Index: 0,
				Text:  text,
				Icon:  icon,
			},
		},
	}
	json, err := json.Marshal(data)
	if err != nil {
		log.Println("failed marshalling message", err)
	} else {
		jsonString := bytes.NewBuffer(json)
		req, _ := http.NewRequest("POST", lametricPushUrl, jsonString)
		req.Header.Set("Accept", "application/json")
		req.Header.Set("X-Access-Token", lametricToken)

		resp, err := http.DefaultClient.Do(req)
		defer resp.Body.Close()
		if err != nil {
			fmt.Println("failed to push", err)
		} else {
			log.Println(lametricPushUrl, resp.Status)
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	data, err := parseJiraWebhook(r.Body)
	if err != nil {
		http.Error(w, "invalid payload ("+err.Error()+")", 500)
	} else {
		text := data.Issue.Key + ": " + data.Issue.Fields.Summary
		go pushToLametric(text, "i142")
		fmt.Fprintf(w, "%s", "OK! Received and redirected...")
	}
}

func main() {
	port := os.Getenv("PORT")
	lametricPushUrl = os.Getenv("LAMETRIC_PUSH_URL")
	lametricToken = os.Getenv("LAMETRIC_TOKEN")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	if lametricPushUrl == "" || lametricToken == "" {
		log.Fatal("$LAMETRIC_PUSH_URL and $LAMETRIC_TOKEN must be set")
	}

	log.Println("start listening on port", port+"...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
