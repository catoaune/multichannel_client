// Package slack implements functions for sending messages to a Slack channel
// You need a Slack incoming webhook url to be able to messages to a Slack channel
package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

//Config for Slack
type Config struct {
	ConfigType string
	URL        string
}

//RequestBody struct for data being sent to Slack
type RequestBody struct {
	Text string `json:"text"`
}

type RequestBodyFormatted struct {
	Blocks []Blocks `json:"blocks"`
}
type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
type Blocks struct {
	Type string `json:"type"`
	Text Text   `json:"text"`
}

//NewConfig returns a new Config
func NewConfig(URL string) Config {
	newConfig := Config{ConfigType: "Slack", URL: URL}
	return newConfig
}

// SendNotification will post to an 'Incoming Webook' url setup in Slack Apps. It accepts
// some text and the slack channel is saved within Slack.
func (c Config) SendNotification(msg string) error {

	slackBody, _ := json.Marshal(RequestBody{Text: msg})
	req, err := http.NewRequest(http.MethodPost, c.URL, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}

// SendFormattedNotification will post to an 'Incoming Webook' url setup in Slack Apps. It accepts
// markdown formatted text and the slack channel is saved within Slack.
func (c Config) SendFormattedNotification(msg string) error {
	requestBodyFormatted := new(RequestBodyFormatted)
    blocks := new(Blocks)
    text := new(Text)


	text.Type = "mrkdwn"
	text.Text = msg
	blocks.Type = "section"
	blocks.Text = *text

	var block = []Blocks{}
	block = append(block, *blocks)

	requestBodyFormatted.Blocks = block
	slackBody, _ := json.Marshal(requestBodyFormatted)
	req, err := http.NewRequest(http.MethodPost, c.URL, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}