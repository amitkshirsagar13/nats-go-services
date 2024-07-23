package main

import (
	"cmd/main/pkg/exit"
	natsconsumer "cmd/main/pkg/nats-consumer"
	natsproducer "cmd/main/pkg/nats-producer"
	natsServer "cmd/main/pkg/nats-server"
	"context"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go exit.CallExit(cancel)

	// create the local server
	natsServer.CreateNATServer()

	subject := "test"

	// register the consumer
	go natsconsumer.GetConsumer(ctx, subject)

	// register the producer
	go natsproducer.SetProducer(ctx, subject)

	<-ctx.Done()

	log.Println("server shutdown completed")
	log.Println("exiting gracefully")
}
