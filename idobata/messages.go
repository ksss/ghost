package idobata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Body      string `json:"body"`
	BodyPlain string `json:"body_plain"`
	CreatedAt string `json:"created_at"`
	Embeds    []struct {
		HTML     string `json:"html"`
		Provider string `json:"provider"`
	} `json:"embeds"`
	ID            int      `json:"id"`
	ImageUrls     []string `json:"image_urls"`
	Mentions      []int    `json:"mentions"`
	RoomID        int      `json:"room_id"`
	SenderIconURL string   `json:"sender_icon_url"`
	SenderID      int      `json:"sender_id"`
	SenderName    string   `json:"sender_name"`
	SenderType    string   `json:"sender_type"`
}

type Post struct {
	Message *Message `json:"message"`
}

func (msg *Message) IsMentionsInclude(id int) bool {
	for _, mid := range msg.Mentions {
		if mid == id {
			return true
		}
	}
	return false
}

func CreateMessage(roomID string, source string) (*Message, error) {
	req, err := http.NewRequest(
		"POST",
		"https://idobata.io/api/messages",
		bytes.NewBufferString(
			fmt.Sprintf(
				"message[room_id]=%s&message[source]=%s&message[format]=markdown",
				roomID,
				source,
			),
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
	var post Post
	err = json.NewDecoder(resp.Body).Decode(&post)
	if err != nil {
		return nil, err
	}
	return post.Message, nil
}
