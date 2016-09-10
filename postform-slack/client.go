package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

type SlackClient struct {
	Token string
}

type slackResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

var avatarRegex = regexp.MustCompile("^:[^:]+:$")

const slackEndpoint = "https://slack.com/api/chat.postMessage"

func (slackClient SlackClient) getAvatarField(avatar string) string {
	if avatarRegex.MatchString(avatar) {
		return "icon_emoji"
	} else {
		return "icon_url"
	}
}

func (slackClient SlackClient) getPayload(channel, username, avatar, message string) string {
	avatarField := slackClient.getAvatarField(avatar)
	payload := url.Values{
		"token":     {slackClient.Token},
		"channel":   {channel},
		"username":  {username},
		"text":      {message},
		"as_user":   {"false"},
		"parse":     {"full"},
		avatarField: {avatar},
	}
	return payload.Encode()
}

func (slackClient SlackClient) PostMessage(channel, username, avatar, message string) error {
	payload := slackClient.getPayload(channel, username, avatar, message)
	data := bytes.NewBufferString(payload)
	resp, err := http.Post(slackEndpoint, "application/x-www-form-urlencoded", data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var response slackResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	if !response.Ok {
		return errors.New(response.Error)
	}
	return nil
}

func New(token string) *SlackClient {
	return &SlackClient{Token: token}
}
