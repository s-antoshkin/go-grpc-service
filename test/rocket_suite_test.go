// +acceptance

package test

import (
	"context"
	"testing"

	rocket "github.com/s-antoshkin/go-grpc-service/rocket-protos/rocket/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RocketTestSuite struct {
	suite.Suite
}

func (s *RocketTestSuite) TestAddRocket() {
	s.T().Run("adds a new rocket soccessfully", func(t *testing.T) {
		client := GetClient()
		resp, err := client.AddRocket(
			context.Background(),
			&rocket.AddRocketRequest{
				Rocket: &rocket.Rocket{
					Id:   "b4fa20be-e8a9-4d00-84e6-54bc69ad4ba7",
					Name: "V1",
					Type: "Falcon",
				},
			},
		)
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), "b4fa20be-e8a9-4d00-84e6-54bc69ad4ba7", resp.Rocket.Id)
	})

	s.T().Run("validates the uuid in the new rocket is a uuid", func(t *testing.T) {
		client := GetClient()
		_, err := client.AddRocket(
			context.Background(),
			&rocket.AddRocketRequest{
				Rocket: &rocket.Rocket{
					Id:   "not-a-valid-uuid",
					Name: "V1",
					Type: "Falcon",
				},
			},
		)
		assert.Error(s.T(), err)
		st := status.Convert(err)
		assert.Equal(s.T(), codes.InvalidArgument, st.Code())
	})
}

func TestRocketService(t *testing.T) {
	suite.Run(t, new(RocketTestSuite))
}
