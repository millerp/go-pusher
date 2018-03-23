package main

// Display on console live Trades and Orders book from Bitstamp
// run with
// 		go run client.go

import (
	"github.com/toorop/go-pusher"
	"log"
)

const (
	APP_KEY = "0e7966c095f399721b75" // bitstamp
)

func main() {

	log.Println("init...")
	//pusherClient, err := pusher.NewClient(APP_KEY)
	// if you need to connect to custom endpoint
	pusherClient, err := pusher.NewCustomClient(APP_KEY, "ws-us2.pusher.com:443", "wss")
	if err != nil {
		log.Fatalln(err)
	}
	// Subscribe
	err = pusherClient.Subscribe("live_trades")
	if err != nil {
		log.Println("Subscription error : ", err)
	}

	log.Println("first subscribe done")

	// Bind events
	dataChannelTrade, err := pusherClient.Bind("data")
	if err != nil {
		log.Println("Bind error: ", err)
	}
	log.Println("Binded to 'data' event")

	// Test bind err
	errChannel, err := pusherClient.Bind(pusher.ErrEvent)
	if err != nil {
		log.Println("Bind error: ", err)
	}
	log.Println("Binded to 'ErrEvent' event")

	log.Println("init done")

	for {
		select {
		case dataEvt := <-dataChannelTrade:
			log.Println("ORDER BOOK: " + dataEvt.Data)
		case errEvt := <-errChannel:
			log.Println("ErrEvent: " + errEvt.Data)
			pusherClient.Close()
		}
	}
}
