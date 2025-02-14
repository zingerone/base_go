package main

import (
	"context"
	"github.com/jessevdk/go-flags"
	"github.com/zingerone/base_go/cmd/api"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var opts struct {
		Command string `short:"c" long:"command" description:"Command to run (e.g., run-api, run-socket, run-event)" required:"true"`
	}
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		cancel()
	}()
	switch opts.Command {
	case "api":
		api.RunAPI(ctx)
	case "socket":
		// Add your RunSocket function call here
		log.Println("Running Socket server...")
	case "event":
		// Add your RunEvent function call here
		log.Println("Running Event server...")
	default:
		log.Fatalf("Unknown command: %s", opts.Command)
	}

}
