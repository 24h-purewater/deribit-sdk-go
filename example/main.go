package main

import (
	"github.com/24h-purewater/deribit-sdk-go"
	"log"
)

const (
	clientID = "[your clientID]"
	secret   = "[your secret]"
)

func main() {
	client := deribit.New(deribit.EndpointProd, clientID, secret)
	orderBook, err := client.GetOrderBook("BTC-12NOV21-45000-P", 3)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(orderBook)
}
