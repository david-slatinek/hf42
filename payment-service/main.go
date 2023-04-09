package main

import (
	"encoding/json"
	"fmt"
	"log"
	"main/client"
	"main/db"
	"main/email"
	"main/env"
	"main/file"
	"main/messaging"
	"main/model"
	"main/pdf"
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

	cli, err := client.NewClient()
	if err != nil {
		log.Fatalf("error with client: %s\n", err)
	}
	defer func(cli client.Client) {
		if err := cli.Close(); err != nil {
			log.Printf("error with closing client: %s\n", err)
		}
	}(cli)

	paymentDB, err := db.NewPaymentDB()
	if err != nil {
		log.Fatalf("error with payment db: %s\n", err)
	}
	defer func(paymentDB db.PaymentDB) {
		if err := paymentDB.Close(); err != nil {
			log.Printf("error with closing payment db: %s\n", err)
		}
	}(paymentDB)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		orders, err := msg.Consume()
		if err != nil {
			log.Printf("error with consuming: %s\n", err)
		} else {
			log.Printf("orders length: %d\n", len(orders))

			for order := range orders {
				var ord model.Order
				err := json.Unmarshal(order.Body, &ord)
				if err != nil {
					log.Printf("error with unmarshalling: %s\n", err)
					continue
				}

				err = cli.ValidateBooks(ord)
				if err != nil {
					log.Printf("error with validating books: %s\n", err)
					continue
				}
				log.Printf("books validated for order: %s\n", ord.OrderID)

				err = paymentDB.CreateOrder(ord)
				if err != nil {
					log.Printf("error with creating order: %s\n", err)
					continue
				}
				log.Printf("order created: %s\n", ord.OrderID)

				err = pdf.CreatePDF(ord)
				if err != nil {
					log.Printf("error with pdf: %s\n", err)
					continue
				}
				log.Printf("pdf created for order: %s\n", ord.OrderID)

				location, err := file.UploadFile(ord.OrderID)
				if err != nil {
					log.Printf("error with uploading file: %s\n", err)
					continue
				}
				log.Printf("file uploaded for order: %s\n", ord.OrderID)

				err = email.SendEmail(location)
				if err != nil {
					log.Printf("error with sending email: %s\n", err)
					continue
				}
				log.Printf("email sent for order: %s\n", ord.OrderID)
			}
		}
	}()

	fmt.Println("waiting for orders...")
	<-c
	fmt.Println("exiting...")
}
