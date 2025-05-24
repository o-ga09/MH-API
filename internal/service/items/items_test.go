package items

import (
	"context"
	"errors"
	"mh-api/internal/domain/items"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockItemRepository struct {
	FindAllFunc func(ctx context.Context) (items.Items, error)
	SaveFunc    func(ctx context.Context, m items.Item) error
	RemoveFunc  func(ctx context.Context, itemId string) error
}

func (m *MockItemRepository) FindAll(ctx context.Context) (items.Items, error) {
	if m.FindAllFunc != nil {
		return m.FindAllFunc(ctx)
	}
	return nil, errors.New("FindAllFunc not implemented in mock")
}

func (m *MockItemRepository) Save(ctx context.Context, i items.Item) error {
	if m.SaveFunc != nil {
		return m.SaveFunc(ctx, i)
	}
	return errors.New("SaveFunc not implemented in mock")
}

func (m *MockItemRepository) Remove(ctx context.Context, itemId string) error {
	if m.RemoveFunc != nil {
		return m.RemoveFunc(ctx, itemId)
	}
	return errors.New("RemoveFunc not implemented in mock")
}

func TestService_GetAllItems_Success(t *testing.T) {
	ctx := context.Background()
	mockRepo := &MockItemRepository{}
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
	mockRepo := &MockItemRepository{}
	itemService := NewService(mockRepo)

	mockRepo.FindAllFunc = func(ctx context.Context) (items.Items, error) {
		return items.Items{}, nil
	}

	expectedDTO := &ItemListResponseDTO{
		Items: []ItemDTO{},
	}

	actualDTO, err := itemService.GetAllItems(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedDTO, actualDTO)
}

func TestService_GetAllItems_RepositoryError(t *testing.T) {
	ctx := context.Background()
	mockRepo := &MockItemRepository{}
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
