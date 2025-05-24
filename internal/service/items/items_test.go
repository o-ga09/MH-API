package items

import (
	"context"
	"errors"
	"mh-api/internal/domain/items"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_GetAllItems_Success(t *testing.T) {
	ctx := context.Background()
	mockRepo := &items.RepositoryMock{}
	itemService := NewService(mockRepo)

	dummyDomainItems := items.Items{
		*items.NewItem("1", "Item One", "url1"),
		*items.NewItem("2", "Item Two", "url2"),
	}

	mockRepo.FindAllFunc = func(ctx context.Context) (items.Items, error) {
		return dummyDomainItems, nil
	}

	expectedDTO := &ItemListResponseDTO{
		Items: []ItemDTO{
			{ItemID: "1", ItemName: "Item One"},
			{ItemID: "2", ItemName: "Item Two"},
		},
	}

	actualDTO, err := itemService.GetAllItems(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedDTO, actualDTO)
}

func TestService_GetAllItems_Empty(t *testing.T) {
	ctx := context.Background()
	mockRepo := &items.RepositoryMock{}
	itemService := NewService(mockRepo)

	mockRepo.FindAllFunc = func(ctx context.Context) (items.Items, error) {
		return items.Items{}, nil
	}

	actualDTO, err := itemService.GetAllItems(ctx)

	assert.NoError(t, err)
	assert.Len(t, actualDTO.Items, 0)
}

func TestService_GetAllItems_RepositoryError(t *testing.T) {
	ctx := context.Background()
	mockRepo := &items.RepositoryMock{}
	itemService := NewService(mockRepo)

	expectedError := errors.New("repository error")
	mockRepo.FindAllFunc = func(ctx context.Context) (items.Items, error) {
		return nil, expectedError
	}

	actualDTO, err := itemService.GetAllItems(ctx)

	assert.Error(t, err)
	assert.Nil(t, actualDTO)
	assert.Equal(t, expectedError, err)
}
