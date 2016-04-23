package pusher

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"

	"golang.org/x/net/websocket"
)

const (
	Host                = "ws.pusherapp.com"
	WebSocketSecurePort = "443"
)

type Conn struct {
	ws *websocket.Conn
}

type ConnectionEstablishedData struct {
	SocketID string `json:"socket_id"`
}

type SubscribeData struct {
	Channel     string `json:"channel"`
	Auth        string `json:"auth"`
	ChannelData string `json:"channel_data"`
}

type SubscribeEvent struct {
	Event string         `json:"event"`
	Data  *SubscribeData `json:"data"`
}

type AnyEvent map[string]interface{}

type ErrorEvent struct {
	Data struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"data"`
	Event string `json:"event"`
}

func NewConn(appKey string) (*Conn, error) {
	wURL := fmt.Sprintf(
		"wss://%s:%s/app/%s?client=%s&version=%s&protocol=%s",
		Host,
		WebSocketSecurePort,
		appKey,
		"ghost",
		"0.0.1",
		"6",
	)
	ws, err := websocket.Dial(wURL, "", "http://localhost/")
	if err != nil {
		return nil, err
	}
	conn := &Conn{
		ws: ws,
	}
	runtime.SetFinalizer(conn, connFinalizer)
	return conn, nil
}

func connFinalizer(conn *Conn) {
	conn.ws.Close()
	return
}

func (conn *Conn) Receive(expect interface{}) error {
	var event AnyEvent
	if err := json.NewDecoder(conn.ws).Decode(&event); err != nil {
		return err
	}

	if expect != nil {
		data, ok := event["data"].(string)
		if !ok {
			msg := fmt.Sprintf("Receive unexpected data. event=%s, data=%v", event["event"], event["data"])
			return errors.New(msg)
		}
		if err := json.Unmarshal([]byte(data), expect); err != nil {
			return err
		}
	}
	return nil
}

func (conn *Conn) Subscribe(data *SubscribeData) error {
	return websocket.JSON.Send(conn.ws, &SubscribeEvent{
		Event: "pusher:subscribe",
		Data:  data,
	})
}
