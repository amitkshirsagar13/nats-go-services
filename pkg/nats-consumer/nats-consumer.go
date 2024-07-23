package natsconsumer

import (
	"context"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func GetConsumer(ctx context.Context, subject string) {

	nc, err := nats.Connect("nats://127.0.0.1:15000")
	if err != nil {
		log.Fatal("Failed to connect to NATS server:", err)
	}
	defer nc.Close()

	fmt.Println("Connected to NATS server on port 15000")

	messages := make(chan *nats.Msg, 1000)

	// we're subscribing to the subject
	// and assigning our channel as reference to receive messages there
	subscription, err := nc.ChanSubscribe(subject, messages)
	if err != nil {
		log.Fatal("Failed to subscribe to subject:", err)
	}

	defer func() {
		subscription.Unsubscribe()
		close(messages)
	}()

	log.Println("Subscribed to", subject)

	for {
		select {
		case <-ctx.Done():
			log.Println("exiting from consumer")
			return
		case msg := <-messages:
			log.Println("received", string(msg.Data))
		}
	}
}
