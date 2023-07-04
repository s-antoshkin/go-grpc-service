package db

import (
	"context"
	"errors"
	"log"

	"github.com/s-antoshkin/go-grpc-service/internal/rocket"
	uuid "github.com/satori/go.uuid"
)

// GetRocketByID - retrieves a rocket from the database by id
func (s *Store) GetRocketByID(ctx context.Context, id string) (rocket.Rocket, error) {
	var rkt rocket.Rocket
	row := s.db.QueryRowContext(
		ctx,
		`SELECT id, name, type FROM rockets WHERE id=$1;`,
		id,
	)
	err := row.Scan(&rkt.ID, &rkt.Name, &rkt.Type)
	if err != nil {
		log.Println(err.Error())
		return rocket.Rocket{}, err
	}

	return rkt, nil
}

// InsertRocket - inserts a rocket into the rockets table
func (s *Store) InsertRocket(ctx context.Context, rkt rocket.Rocket) (rocket.Rocket, error) {
	_, err := s.db.NamedQueryContext(
		ctx,
		`INSERT INTO rockets (id, name, type)
		VALUES (:id, :name, :type)`,
		rkt,
	)
	if err != nil {
		log.Println(err.Error())
		return rocket.Rocket{}, errors.New("failed to insert into database")
	}

	return rocket.Rocket{
		ID:   rkt.ID,
		Name: rkt.Name,
		Type: rkt.Type,
	}, nil
}

func (s *Store) DeleteRocket(ctx context.Context, id string) error {
	uid, err := uuid.FromString(id)
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(
		ctx,
		`DELETE FROM rockets WHERE id = $1`,
		uid,
	)
	if err != nil {
		return err
	}

	return nil
}
