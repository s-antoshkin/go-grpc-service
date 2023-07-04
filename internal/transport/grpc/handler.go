package grpc

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/s-antoshkin/go-grpc-service/internal/rocket"
	rkt "github.com/s-antoshkin/go-grpc-service/rocket-protos/rocket/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RocketService - define the interface that the concrete implementation
// has to adhere to
type RocketService interface {
	GetRocketByID(ctx context.Context, id string) (rocket.Rocket, error)
	InsertRocket(ctx context.Context, rkt rocket.Rocket) (rocket.Rocket, error)
	DeleteRocket(ctx context.Context, id string) error
}

// Handler - will handle incoming gRPC requests
type Handler struct {
	RocketService RocketService
	rkt.UnimplementedRocketServiceServer
}

// New - returns a new gRPC handler
func New(rktService RocketService) Handler {
	return Handler{
		RocketService: rktService,
	}
}

func (h *Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("could not listen on port 50051")
		return err
	}

	grpsServer := grpc.NewServer()
	rkt.RegisterRocketServiceServer(grpsServer, h)

	if err = grpsServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %s\n", err)
		return err
	}

	return nil
}

// GetRocket - retrieves a rocket by id and returns the response
func (h *Handler) GetRocket(ctx context.Context, r *rkt.GetRocketRequest) (*rkt.GetRocketResponse, error) {
	log.Println("Get Rocket gRPC endpoint hit")

	rocket, err := h.RocketService.GetRocketByID(ctx, r.Id)
	if err != nil {
		log.Printf("Failed to retrive rocket by ID")
		return &rkt.GetRocketResponse{}, err
	}

	return &rkt.GetRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   rocket.ID,
			Name: rocket.Name,
			Type: rocket.Type,
		},
	}, nil
}

// AddRocket - adds a rocket to the database
func (h *Handler) AddRocket(ctx context.Context, r *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	log.Println("Add Rocket gRPC endpoint hit")

	if _, err := uuid.Parse(r.Rocket.Id); err != nil {
		errorStatus := status.Error(codes.InvalidArgument, "uuid is not valid")
		log.Println("given uuid is not valid")
		return &rkt.AddRocketResponse{}, errorStatus
	}

	newRkt, err := h.RocketService.InsertRocket(ctx, rocket.Rocket{
		ID:   r.Rocket.Id,
		Name: r.Rocket.Name,
		Type: r.Rocket.Type,
	})
	if err != nil {
		log.Println("failed to insert rocket into database")
		return &rkt.AddRocketResponse{}, err
	}

	return &rkt.AddRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   newRkt.ID,
			Name: newRkt.Name,
			Type: newRkt.Type,
		},
	}, nil
}

// DeleteRocket - handler for deleting a rocket
func (h *Handler) DeleteRocket(ctx context.Context, r *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	log.Println("Delete Rocket gRPC endpoint hit")

	err := h.RocketService.DeleteRocket(context.Background(), r.Id)
	if err != nil {
		return &rkt.DeleteRocketResponse{}, err
	}

	return &rkt.DeleteRocketResponse{
		Status: "successfully deleted rocket",
	}, nil
}
