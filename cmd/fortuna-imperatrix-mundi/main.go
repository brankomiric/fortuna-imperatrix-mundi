package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/brankomiric/fortuna-imperatrix-mundi/internal/server"
	"github.com/subosito/gotenv"
)

func init() {
	err := gotenv.Load()
	if err != nil {
		log.Printf("gotenv.Load() error: %s\n", err.Error())
	}
}

func main() {
	log.Println("Starting service...")

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("PORT is not set. Using default port 3000.")
		port = "3000"
	}

	app := server.SetupRouter()

	go func() {
		log.Fatal(app.Listen(port))
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT, os.Interrupt)

	<-stop

	stopCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	log.Println("Shutting down service")

	err := app.ShutdownWithContext(stopCtx)
	if err != nil {
		log.Printf("app.ShutdownWithContext() error: %s\n", err.Error())
	}
}
