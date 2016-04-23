package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ksss/ghost/idobata"
	"github.com/ksss/ghost/pusher"
)

func main() {
	seed, err := idobata.ShowSeed()
	if err != nil {
		panic(err)
	}

	conn, err := pusher.NewConn(idobata.PusherAppKey)
	if err != nil {
		panic(err)
	}

	var data *pusher.ConnectionEstablishedData
	// expect pusher:connection_established
	err = conn.Receive(&data)
	if err != nil {
		panic(err)
	}

	auth, err := idobata.CreateAuth(data.SocketID, seed.Records.Bot.ChannelName)
	if err != nil {
		panic(err)
	}

	err = conn.Subscribe(&pusher.SubscribeData{
		Channel:     seed.Records.Bot.ChannelName,
		Auth:        auth.Auth,
		ChannelData: auth.ChannelData,
	})
	if err != nil {
		panic(err)
	}

	// expect pusher_internal:subscription_succeeded
	if err := conn.Receive(nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("run loop")
	for {
		var post idobata.Post
		for i := 0; i < 2; i++ {
			if err := conn.Receive(&post); err != nil {
				panic(err)
			}
		}
		if !post.Message.IsMentionsInclude(seed.Records.Bot.ID) {
			continue
		}
		go func() {
			_, err = idobata.CreateMessage(strconv.Itoa(post.Message.RoomID), "せやな")
			if err != nil {
				panic(err)
			}
		}()
	}
}
