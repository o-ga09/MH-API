package mysql

import (
	"context"
	"mh-api/internal/domain/items"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemQueryService_Save(t *testing.T) {
	ctx := context.Background()
	service := NewItemQueryService()

	dummyItem := items.NewItem("dummyId", "dummyName", "dummyUrl")

	err := service.Save(ctx, *dummyItem)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Save method not implemented")
}

func TestItemQueryService_Remove(t *testing.T) {
	ctx := context.Background()
	service := NewItemQueryService()

	err := service.Remove(ctx, "dummyId")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Remove method not implemented")
}
