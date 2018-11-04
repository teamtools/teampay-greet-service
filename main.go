package main

import (
	"fmt"
	"log"

	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/mdns"
	"github.com/teamtools/teampay-greet-service/proto/greetpb"
	"golang.org/x/net/context"
)

type Repository interface {
	Greet(*greetpb.GreetRequest) (*greetpb.GreetResponse, error)
}
type service struct {
	repo Repository
}

func main() {
	fmt.Println("Starting up...")

	s := micro.NewService(
		micro.Name("teampay.greet"),
		// micro.Version("latest"),
	)

	s.Init()
	greetpb.RegisterGreetServiceHandler(s.Server(), &service{})

	if err := s.Run(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// HANDLERS
func (s *service) Greet(ctx context.Context, req *greetpb.GreetRequest, res *greetpb.GreetResponse) error {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstname()
	lastName := req.GetGreeting().GetLastname()

	greeting := "Hello " + firstName + " " + lastName
	fmt.Printf("Responding %v\n", greeting)
	res.Response = greeting
	fmt.Printf("Responded: %v", res)
	return nil
}
