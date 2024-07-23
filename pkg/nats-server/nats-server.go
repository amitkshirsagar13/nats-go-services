package natsServer

import (
	"log"

	natsServer "github.com/nats-io/nats-server/v2/server"
)

func CreateNATServer() {

	// options to create nats server
	opts := &natsServer.Options{
		ServerName:     "local_nats_server",
		Host:           "localhost",
		Port:           15000,
		NoLog:          false,
		NoSigs:         false,
		MaxControlLine: 4096,
		MaxPayload:     65536,
	}

	// initialize the server struct
	server, err := natsServer.NewServer(opts)
	if err != nil {
		log.Fatal(err)
	}

	// run the nats server based on the options
	err = natsServer.Run(server)
	if err != nil {
		log.Fatal("Failed to start NATS server:", err)
	}

	log.Println("NATS server started")
}
