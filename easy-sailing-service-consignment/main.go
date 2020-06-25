package main

import (
	"context"
	"log"
	"net"
	"sync"

	// Import generated protobuf code
	pb "github.com/PFurtak/easy-sailing/tree/dev/easy-sailing-service-consignment/proto/consignment"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}

// Repository - dummy repo, simulates the use of a datastore. Will replace with a real implementation
type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

// Create a new consignment
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.mu.Unlock()
	return consignment, nil
}

//Service should implement all of the methods to satisfy the service defined in our protobuf definition.
//You can check the interface in the generated code for the exact signatures.

type service struct {
	repo repository
}

//Create consignment - we defined just one method on this service, which is the create method.
//This takes in a context and a request as an argument, and is handled by thr gRPC server.
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (pb.response, error) {
	//save consignment
	consignment, err := s.repo.Create(req)

	if err != nil {
		return nil, err
	}

	//Return matching the 'response' message we created in our protobuf definition.
	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func main() {
	repo := &Repository{}

	// Set up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Register our service with the gRPC server, this will tie the implementation into the auto-generated 
	// interface code for our protobuf definitions.
	pb.RegisterShippingServiceServer(s, %service{repo})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.PrintIn("Running on port", port)
	id err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
