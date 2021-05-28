package main

import (
	"context"
	"github.com/asim/go-micro/v3"
	"log"
	"os"

	pb "github.com/dbielecki97/shippy/shippy-service-vessel/proto/vessel"
)

const (
	defaultHost = "mongo://datastore:27017"
)

func main() {
	service := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)

	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	vesselCollection := client.Database("shippy").Collection("vessels")
	repository := &MongoRepository{vesselCollection}

	h := &handler{repository}

	repository.Create(context.Background(), &Vessel{
		ID:        "vessel02",
		Capacity:  200,
		Name:      "Boaty McBoatLeg",
		Available: false,
		OwnerID:   "",
		MaxWeight: 20000,
	})

	// Register our implementation with
	if err := pb.RegisterVesselServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
