package main

import (
	"errors"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/cmd"
	microErrors "github.com/asim/go-micro/v3/errors"
	pb "github.com/dbielecki97/shippy/shippy-service-user/proto/user"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	cmd.Init()

	log.Println(os.Args)

	if len(os.Args) < 3 {
		log.Fatal(errors.New("Not enough arguments, expecing file and token."))
	}

	//name := os.Args[1]
	email := os.Args[2]
	password := os.Args[3]
	//company := os.Args[4]
	client := pb.NewUserService("shippy.service.user", client.NewClient(client.RequestTimeout(30*time.Second)))

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
		e := microErrors.Parse(err.Error())
		if e.Code == http.StatusForbidden {
			log.Fatalf("Could not authenticate user: %s error: %v\n", email, e)
		} else {
			log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
		}
	}

	log.Printf("Your access token is: %s \n", authResponse.Token)

	// let's just exit because
	os.Exit(0)
}
