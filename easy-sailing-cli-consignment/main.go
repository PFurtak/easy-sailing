package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"context"

	pb "github.com/PFurtak/easy-sailing/easy-sailing-service-consignment/proto/consignment"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pbConsignment, error) {
	var consignment *pb.Consignment data, err := ioutil.ReadFile(file)
	if err != nil {
		retun nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	// connect to server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pg.NewShippingServiceClient(conn)

	// contact server and print out response
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	}