package main

import (
	"fmt"
	"log"
	"main/env"
	"main/messaging"
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
		log.Printf("error with messaging: %s\n", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		orders, err := msg.Consume()
		if err != nil {
			log.Printf("error with consuming: %s\n", err)
		} else {
			for order := range orders {
				if order.Ack(false) != nil {
					log.Printf("error with acking: %s\n", err)
					continue
				}
				log.Printf("Received a message: %s", order.Body)
			}
		}
	}()

	fmt.Println("Waiting for orders...")
	<-c
}
