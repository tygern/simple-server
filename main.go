package main

import (
	"fmt"
	"github.com/tygern/simple-server/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s := server.Create(":8080")
	go server.Start(s)
	defer server.Stop(s)

	waitForInterrupt()
}

func waitForInterrupt() os.Signal {
	interruptChan := make(chan os.Signal)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM)

	sig := <-interruptChan
	fmt.Printf("\n%s received\n", sig)
	return sig
}
