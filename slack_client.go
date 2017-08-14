package bashir

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type SlackClient struct {
	WebhookUrl string
	Channel    string
	Icon       string
	BotName    string
}

func NewSlackClient(config *Config) *SlackClient {
	return &SlackClient{
		WebhookUrl: config.Slack.WebhookUrl,
		Channel:    config.Slack.Channel,
		Icon:       config.Slack.Icon,
		BotName:    config.Slack.BotName,
	}
}

func (c *SlackClient) SendMessage(message string) error {
	data := c.GetData(message)

	req, _ := http.NewRequest("POST", c.WebhookUrl, bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	client := &http.Client{}
	resp, _ := client.Do(req)

	_, _ = ioutil.ReadAll(resp.Body)

	return nil
}

func (c *SlackClient) GetData(message string) url.Values {
	data := url.Values{}
	jsonPayload := `
			{
				"channel": "%s",
				"username": "$%s",
				"text": "%s",
				"icon_emoji": "%s"
			}
		`

	jsonMessage := fmt.Sprintf(jsonPayload, c.Channel, c.BotName, message, c.Icon)
	data.Set("payload", jsonMessage)
	return data
}
