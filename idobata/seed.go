package idobata

import (
	"encoding/json"
	"net/http"
)

type Seed struct {
	Records struct {
		Bot struct {
			ChannelName              string      `json:"channel_name"`
			CreatedAt                string      `json:"created_at"`
			Email                    interface{} `json:"email"`
			EnterBehaviorDesktop     string      `json:"enter_behavior_desktop"`
			EnterBehaviorMobile      string      `json:"enter_behavior_mobile"`
			IconURL                  string      `json:"icon_url"`
			ID                       int         `json:"id"`
			JoinIds                  []int       `json:"join_ids"`
			MembershipIds            []int       `json:"membership_ids"`
			MessageFoldable          bool        `json:"message_foldable"`
			Name                     string      `json:"name"`
			ReceiveBroadcastMentions bool        `json:"receive_broadcast_mentions"`
			Status                   string      `json:"status"`
			UseMarkdown              bool        `json:"use_markdown"`
		} `json:"bot"`
		BotJoins []struct {
			GuyID                    int           `json:"guy_id"`
			ID                       int           `json:"id"`
			ReceiveBroadcastMentions interface{}   `json:"receive_broadcast_mentions"`
			RoomID                   int           `json:"room_id"`
			Starred                  bool          `json:"starred"`
			UnreadMentionIds         []interface{} `json:"unread_mention_ids"`
			UnreadMessageIds         []interface{} `json:"unread_message_ids"`
		} `json:"bot_joins"`
		Joins []struct {
			GuyID  int `json:"guy_id"`
			ID     int `json:"id"`
			RoomID int `json:"room_id"`
		} `json:"joins"`
		Memberships []struct {
			GuyID          int    `json:"guy_id"`
			ID             int    `json:"id"`
			OrganizationID int    `json:"organization_id"`
			Role           string `json:"role"`
		} `json:"memberships"`
		Organizations []struct {
			ID    int `json:"id"`
			Links struct {
				Bots        string `json:"bots"`
				Invitations string `json:"invitations"`
				Rooms       string `json:"rooms"`
			} `json:"links"`
			MembershipIds []int  `json:"membership_ids"`
			Name          string `json:"name"`
			Slug          string `json:"slug"`
		} `json:"organizations"`
		Rooms []struct {
			BotJoinIds        []int       `json:"bot_join_ids"`
			Description       string      `json:"description"`
			DescriptionSource string      `json:"description_source"`
			ID                int         `json:"id"`
			InvitationEnabled bool        `json:"invitation_enabled"`
			InvitationToken   interface{} `json:"invitation_token"`
			JoinIds           []int       `json:"join_ids"`
			Links             struct {
				HookEndpoints string `json:"hook_endpoints"`
				Messages      string `json:"messages"`
			} `json:"links"`
			Name           string `json:"name"`
			OrganizationID int    `json:"organization_id"`
		} `json:"rooms"`
	} `json:"records"`
	Version int `json:"version"`
}

func ShowSeed() (s *Seed, err error) {
	req, err := http.NewRequest(
		"GET",
		"https://idobata.io/api/seed",
		nil,
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
	err = json.NewDecoder(resp.Body).Decode(&s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
