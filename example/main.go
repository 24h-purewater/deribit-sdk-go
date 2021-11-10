package main

import (
	"deribit"
	"log"
)

func main() {
	client := deribit.New(deribit.EndpointProd, "your clientID", "your secret")
	orderBook, err := client.GetOrderBook("BTC-12NOV21-45000-P", 3)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(orderBook)
}
