package main

import (
	"log"

	"github.com/vietbm-hcm/go-grpc-service-course/internal/db"
	"github.com/vietbm-hcm/go-grpc-service-course/internal/rocket"
	"github.com/vietbm-hcm/go-grpc-service-course/internal/transport/grpc"
)

func Run() error {
	
	// responsible for initializing and starting
	// out gRPC server
	rocketStore, err := db.New()
	if err != nil {
		return err
	}

	err = rocketStore.Migrate()
	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}
	
	rktService := rocket.New(rocketStore)
	rktHandler := grpc.New(rktService)

	if err := rktHandler.Serve(); err != nil {
		return err
	}
 
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}