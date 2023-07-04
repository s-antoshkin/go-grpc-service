// +acceptance

package test

import (
	"log"

	rocket "github.com/s-antoshkin/go-grpc-service/rocket-protos/rocket/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetClient() rocket.RocketServiceClient {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	rocketClient := rocket.NewRocketServiceClient(conn)

	return rocketClient
}
