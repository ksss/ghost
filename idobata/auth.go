package idobata

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

type Auth struct {
	Auth        string `json:"auth"`
	ChannelData string `json:"channel_data"`
}

func CreateAuth(socketID string, channelName string) (auth *Auth, err error) {
	req, err := http.NewRequest(
		"POST",
		"https://idobata.io/pusher/auth",
		bytes.NewBufferString(
			url.Values{
				"socket_id":    {socketID},
				"channel_name": {channelName},
			}.Encode(),
		),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-API-Token", APIToken)
	req.Header.Add("User-Agent", "ghost")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&auth)
	if err != nil {
		return nil, err
	}
	return auth, nil
}
