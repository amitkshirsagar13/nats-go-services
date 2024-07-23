package natsproducer

import (
	"context"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func SetProducer(ctx context.Context, subject string) {
	nc, err := nats.Connect("nats://127.0.0.1:15000")
	if err != nil {
		log.Fatal("Failed to connect to NATS server:", err)
	}
	// close the connection on function exit
	defer nc.Close()

	i := 0

	for {
		select {
		case <-ctx.Done():
			log.Println("exiting from producer")
			return
		default:
			i += 1
			message := fmt.Sprintf("message %v", i)

			// Publish the message to the nats server
			err = nc.Publish(subject, []byte(message))
			if err != nil {
				log.Println("Failed to publish message:", err)
			} else {
				log.Println(message)
			}
		}
	}
}
