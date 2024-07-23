package exit

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func CallExit(cancel context.CancelFunc) {
	// listen for interrupts to exit gracefully
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)
	<-sigChannel
	close(sigChannel)
	cancel()
}
