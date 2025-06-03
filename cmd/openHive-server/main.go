package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"OpenHive/internal/server"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	// create context, litening fort interruption signal from OS
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	//waiting interrupt signal
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")
	stop()

	// ctx is used to inform the serber it has 5sec to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	//notify main routine that shutdown is complete
	done <- true
}

func main() {
	server := server.NewServer()

	//create a done channel to signal when shutdown is complete
	done := make(chan bool, 1)

	go gracefulShutdown(server, done)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	//wait for the gracefull shutdown to complete
	<-done
	log.Printf("graceful shutdown complete")
}
