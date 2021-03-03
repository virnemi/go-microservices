package main

import (
	"context"
	"log"
	"net"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"../models"
	"./api"
)

// Starts the Client API and the HTTP server
func main() {
	params, er := models.NewEnvParams(false)
	if er != nil {
		log.Fatalf("Error: Environment parameters not set: %v\n", er)
		return
	}
	log.Printf("Client API running with parameters:\n %#v\n", params)
	router := api.NewPortRoute(params)

	ctx, cancel := context.WithCancel(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/ports/list", router.List) //.Methods("GET")
	mux.HandleFunc("/port", router.Get)        //.Methods("GET")
	mux.HandleFunc("/ports", router.Post)      //.Methods("POST")

	clientApi := &http.Server{
		Addr:        params.ApiAddr,
		Handler:     mux,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}

	go func() {
		if err := clientApi.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Client API ListenAndServe: %v", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	<-signalChan
	log.Print("Gracefully shutting down...\n")

	go func() {
		<-signalChan
		log.Fatal("Terminating...\n")
	}()

	ctx2, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := clientApi.Shutdown(ctx2); err != nil {
		log.Printf("shutdown error: %v\n", err)
		defer os.Exit(1)
		return
	} else {
		log.Printf("Gracefully stopped\n")
	}

	cancel()

	defer os.Exit(0)
	return
}
