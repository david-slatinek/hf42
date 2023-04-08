package main

import (
	"encoding/json"
	"fmt"
	"log"
	"main/env"
	"main/messaging"
	"main/model"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := env.Load("env/.env")
	if err != nil {
		log.Fatalf("failed to load env variables: %s", err)
	}

	msg, err := messaging.NewMessaging()
	if err != nil {
		log.Fatalf("error with messaging: %s\n", err)
	}
	defer msg.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		orders, err := msg.Consume()
		if err != nil {
			log.Printf("error with consuming: %s\n", err)
		} else {
			for order := range orders {

				var ord model.Order
				err := json.Unmarshal(order.Body, &ord)
				if err != nil {
					log.Printf("error with unmarshalling: %s\n", err)
					continue
				}

				log.Printf("Received order: %v", ord)
			}
		}
	}()

	fmt.Println("Waiting for orders...")
	<-c
}
