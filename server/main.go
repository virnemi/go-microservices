package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"../gorpc"
	"../models"
	"./service"
	"google.golang.org/grpc"
)

// Starts the PortDomainService and the gRPC server
func main() {
	log.Print("Starting PortDomainService...\n")
	params, er := models.NewEnvParams(true)
	if er != nil {
		log.Fatalf("Error: Environment parameters not set: %v\n", er)
		return
	}
	grpcAddr := params.GrpcAddr
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Error listening to port %s: %v", grpcAddr, err)
	}
	defer lis.Close()
	_, cancel := context.WithCancel(context.Background())

	grpcServer := grpc.NewServer()
	portDomainService := service.NewPortDomainService()

	go func() {
		gorpc.RegisterPortDomainServer(grpcServer, portDomainService)
		grpcServer.Serve(lis)
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

	grpcServer.GracefulStop()
	cancel()

	defer os.Exit(0)
	return
}
