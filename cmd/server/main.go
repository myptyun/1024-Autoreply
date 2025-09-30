package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"modbus-rtu-server/internal/api"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server := api.NewServer()

	go func() {
		if err := server.Start(); err != nil {
			log.Fatalf("server start failed: %v", err)
		}
	}()

	<-ctx.Done()
	stop()

	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("server shutdown error: %v", err)
	}
}