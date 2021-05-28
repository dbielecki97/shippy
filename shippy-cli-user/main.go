package main

import (
	"github.com/asim/go-micro/v3"
	"log"
	"os"

	pb "github.com/dbielecki97/shippy/shippy-service-user/proto/user"
	"golang.org/x/net/context"
)

func main() {

	service := micro.NewService(
		micro.Name("shippy.cli.user"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	service.Init()

	client := pb.NewUserService("shippy.service.user", service.Client())

	//name := "Dawid Bielecki"
	email := "dawid.bielecki97@gmail.com"
	password := "test123"
	//company := "Acaisoft"

	//r, err := client.Create(context.Background(), &pb.User{
	//	Name:     name,
	//	Email:    email,
	//	Password: password,
	//	Company:  company,
	//})
	//if err != nil {
	//	log.Fatalf("Could not create: %v", err)
	//}
	//log.Printf("Created: %s", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}

	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(context.Background(), &pb.User{
		Email:    email,
		Password: password,
	})

	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
	}

	log.Printf("Your access token is: %s \n", authResponse.Token)

	// let's just exit because
	os.Exit(0)
}
