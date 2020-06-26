module github.com/PFurtak/easy-sailing/easy-sailing-service-consignment

go 1.14

replace github.com/PFurtak/easy-sailing/easy-sailing-service-consignment => ../easy-sailing-service-consignment

require (
	github.com/golang/protobuf v1.4.2
	google.golang.org/grpc v1.30.0
)
