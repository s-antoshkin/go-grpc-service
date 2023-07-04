//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/s-antoshkin/go-grpc-service/internal/rocket Store
package rocket

import "context"

// Rocket - contain the defenition of rocket
type Rocket struct {
	ID   string
	Name string
	Type string
}

// Store - defines the interface we expect database implementation to follow
type Store interface {
	GetRocketByID(ctx context.Context, id string) (Rocket, error)
	InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error)
	DeleteRocket(ctx context.Context, id string) error
}

// Service - rocket service, responsible for updating rocket inventory
type Service struct {
	Store Store
}

// New - return a new instance of rocket service
func New(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// GetRocketByID - retrieves a rocket based on the ID from the Store
func (s *Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(ctx, id)
	if err != nil {
		return Rocket{}, err
	}

	return rkt, nil
}

// InsertRocket - insert a new rocket into the store
func (s *Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(ctx, rkt)
	if err != nil {
		return Rocket{}, err
	}

	return rkt, nil
}

// DeleteRocket - delete a rocket from store
func (s *Service) DeleteRocket(ctx context.Context, id string) error {
	err := s.Store.DeleteRocket(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
