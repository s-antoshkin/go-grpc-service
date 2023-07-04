package rocket

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRocketService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("tests get rocket by id", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.
			EXPECT().
			GetRocketByID(context.Background(), id).
			Return(Rocket{
				ID: id,
			}, nil)

		rocketService := New(rocketStoreMock)
		rkt, err := rocketService.GetRocketByID(
			context.Background(),
			id,
		)

		assert.NoError(t, err)
		assert.Equal(t, "UUID-1", rkt.ID)
	})

	t.Run("tests insert rocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.
			EXPECT().
			InsertRocket(context.Background(), Rocket{
				ID: id,
			}).
			Return(Rocket{
				ID: id,
			}, nil)

		rocketService := New(rocketStoreMock)
		rkt, err := rocketService.InsertRocket(
			context.Background(),
			Rocket{
				ID: id,
			},
		)

		assert.NoError(t, err)
		assert.Equal(t, "UUID-1", rkt.ID)
	})

	t.Run("tests delete rocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.
			EXPECT().
			DeleteRocket(context.Background(), id).
			Return(nil)
		rocketService := New(rocketStoreMock)
		err := rocketService.DeleteRocket(context.Background(), "UUID-1")
		assert.NoError(t, err)

	})
}
