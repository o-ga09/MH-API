package main

import (
	"context"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"mh-api/internal/service/monsters"
	"mh-api/internal/service/weapons"
	"mh-api/internal/service/items"
	"mh-api/internal/service/skills"
)

// Mock services for testing
type MockMonsterService struct {
	mock.Mock
}

func (m *MockMonsterService) FetchMonsterDetail(ctx context.Context, id string) ([]*monsters.FetchMonsterListDto, error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]*monsters.FetchMonsterListDto), args.Error(1)
}

type MockWeaponService struct {
	mock.Mock
}

func (m *MockWeaponService) SearchWeapons(ctx context.Context, params weapons.SearchWeaponsParams) (*weapons.ListWeaponsResponse, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(*weapons.ListWeaponsResponse), args.Error(1)
}

type MockItemService struct {
	mock.Mock
}

func (m *MockItemService) GetAllItems(ctx context.Context) (*items.ItemListResponseDTO, error) {
	args := m.Called(ctx)
	return args.Get(0).(*items.ItemListResponseDTO), args.Error(1)
}

func (m *MockItemService) GetItemByID(ctx context.Context, itemID string) (*items.ItemDTO, error) {
	args := m.Called(ctx, itemID)
	return args.Get(0).(*items.ItemDTO), args.Error(1)
}

func (m *MockItemService) GetItemByMonsterID(ctx context.Context, monsterID string) (*items.ItemByMonster, error) {
	args := m.Called(ctx, monsterID)
	return args.Get(0).(*items.ItemByMonster), args.Error(1)
}

type MockSkillService struct {
	mock.Mock
}

func (m *MockSkillService) GetAllSkills(ctx context.Context) (*skills.SkillListResponseDTO, error) {
	args := m.Called(ctx)
	return args.Get(0).(*skills.SkillListResponseDTO), args.Error(1)
}

func (m *MockSkillService) GetSkillByID(ctx context.Context, skillID string) (*skills.SkillDTO, error) {
	args := m.Called(ctx, skillID)
	return args.Get(0).(*skills.SkillDTO), args.Error(1)
}

func TestMCPServerInitialization(t *testing.T) {
	// Create mock services
	mockMonsterService := &MockMonsterService{}
	mockWeaponService := &MockWeaponService{}
	mockItemService := &MockItemService{}
	mockSkillService := &MockSkillService{}

	mcpServer := &MCPServer{
		monsterService: mockMonsterService,
		weaponService:  mockWeaponService,
		itemService:    mockItemService,
		skillService:   mockSkillService,
	}

	assert.NotNil(t, mcpServer)
	assert.NotNil(t, mcpServer.monsterService)
	assert.NotNil(t, mcpServer.weaponService)
	assert.NotNil(t, mcpServer.itemService)
	assert.NotNil(t, mcpServer.skillService)
}

func TestHandleToolsList(t *testing.T) {
	// Create mock services
	mockMonsterService := &MockMonsterService{}
	mockWeaponService := &MockWeaponService{}
	mockItemService := &MockItemService{}
	mockSkillService := &MockSkillService{}

	mcpServer := &MCPServer{
		monsterService: mockMonsterService,
		weaponService:  mockWeaponService,
		itemService:    mockItemService,
		skillService:   mockSkillService,
	}

	ctx := context.Background()
	request := mcp.Request{
		Method: "tools/list",
		Params: map[string]interface{}{},
	}

	result, err := mcpServer.handleToolsList(ctx, request)
	
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Check that the result contains the expected tools
	resultMap, ok := result.(map[string]interface{})
	assert.True(t, ok)
	
	tools, ok := resultMap["tools"].([]mcp.Tool)
	assert.True(t, ok)
	assert.Len(t, tools, 8) // We should have 8 tools

	// Check tool names
	expectedToolNames := []string{
		"get_monsters", "get_monster_by_id", "get_weapons", "get_items",
		"get_item_by_id", "get_items_by_monster", "get_skills", "get_skill_by_id",
	}

	toolNames := make([]string, len(tools))
	for i, tool := range tools {
		toolNames[i] = tool.Name
	}

	for _, expectedName := range expectedToolNames {
		assert.Contains(t, toolNames, expectedName)
	}
}

func TestGetMonsterByID(t *testing.T) {
	mockMonsterService := &MockMonsterService{}
	mockWeaponService := &MockWeaponService{}
	mockItemService := &MockItemService{}
	mockSkillService := &MockSkillService{}

	mcpServer := &MCPServer{
		monsterService: mockMonsterService,
		weaponService:  mockWeaponService,
		itemService:    mockItemService,
		skillService:   mockSkillService,
	}

	ctx := context.Background()
	monsterID := "1"

	// Mock the service call
	expectedMonsters := []*monsters.FetchMonsterListDto{
		{
			ID:   "1",
			Name: "Test Monster",
		},
	}
	mockMonsterService.On("FetchMonsterDetail", ctx, monsterID).Return(expectedMonsters, nil)

	args := map[string]interface{}{
		"monster_id": monsterID,
	}

	result, err := mcpServer.getMonsterByID(ctx, args)
	
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Verify the mock was called
	mockMonsterService.AssertExpectations(t)
}

func TestGetSkills(t *testing.T) {
	mockMonsterService := &MockMonsterService{}
	mockWeaponService := &MockWeaponService{}
	mockItemService := &MockItemService{}
	mockSkillService := &MockSkillService{}

	mcpServer := &MCPServer{
		monsterService: mockMonsterService,
		weaponService:  mockWeaponService,
		itemService:    mockItemService,
		skillService:   mockSkillService,
	}

	ctx := context.Background()

	// Mock the service call
	expectedSkills := &skills.SkillListResponseDTO{
		Skills: []skills.SkillDTO{
			{
				ID:          "1",
				Name:        "Test Skill",
				Description: "Test Description",
				Level: []skills.SkillLevelDTO{
					{
						Level:       1,
						Description: "Level 1 Description",
					},
				},
			},
		},
	}
	mockSkillService.On("GetAllSkills", ctx).Return(expectedSkills, nil)

	result, err := mcpServer.getSkills(ctx)
	
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Verify the mock was called
	mockSkillService.AssertExpectations(t)
}

func TestGetWeapons(t *testing.T) {
	mockMonsterService := &MockMonsterService{}
	mockWeaponService := &MockWeaponService{}
	mockItemService := &MockItemService{}
	mockSkillService := &MockSkillService{}

	mcpServer := &MCPServer{
		monsterService: mockMonsterService,
		weaponService:  mockWeaponService,
		itemService:    mockItemService,
		skillService:   mockSkillService,
	}

	ctx := context.Background()

	// Mock the service call
	expectedWeapons := &weapons.ListWeaponsResponse{
		Weapons: []weapons.WeaponData{
			{
				WeaponID:    "1",
				Name:        "Test Weapon",
				Description: "Test Description",
			},
		},
		TotalCount: 1,
		Offset:     0,
		Limit:      50,
	}

	// The service expects a SearchWeaponsParams struct
	expectedParams := weapons.SearchWeaponsParams{
		Limit:  nil,
		Offset: nil,
	}
	mockWeaponService.On("SearchWeapons", ctx, expectedParams).Return(expectedWeapons, nil)

	args := map[string]interface{}{}
	result, err := mcpServer.getWeapons(ctx, args)
	
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Verify the mock was called
	mockWeaponService.AssertExpectations(t)
}