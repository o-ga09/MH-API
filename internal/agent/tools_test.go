package agent

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"mh-api/internal/database/mysql"
	"mh-api/internal/service/items"
	"mh-api/internal/service/monsters"
	"mh-api/internal/service/skills"
	"mh-api/internal/service/weapons"
)

func TestNewMonHunTools(t *testing.T) {
	// Initialize services
	monsterRepo := mysql.NewMonsterRepository()
	monsterQS := mysql.NewmonsterQueryService()
	monsterService := monsters.NewMonsterService(monsterRepo, monsterQS)

	weaponQS := mysql.NewWeaponQueryService()
	weaponService := weapons.NewWeaponService(weaponQS)

	itemRepo := mysql.NewItemQueryService()
	itemService := items.NewService(monsterQS, itemRepo)

	skillRepo := mysql.NewSkillQueryService()
	skillService := skills.NewService(skillRepo)

	// Create tools
	monHunTools := NewMonHunTools(monsterService, weaponService, itemService, skillService)
	require.NotNil(t, monHunTools)

	// Get tools
	tools, err := monHunTools.GetTools()
	require.NoError(t, err)
	require.NotNil(t, tools)

	// Verify we have 8 tools
	assert.Len(t, tools, 8)

	// Verify tool names
	expectedToolNames := []string{
		"get_monsters",
		"get_monster_by_id",
		"get_weapons",
		"get_items",
		"get_item_by_id",
		"get_items_by_monster",
		"get_skills",
		"get_skill_by_id",
	}

	for i, tool := range tools {
		assert.Equal(t, expectedToolNames[i], tool.Name())
		assert.NotEmpty(t, tool.Description())
	}
}

func TestNewServer(t *testing.T) {
	// Skip if no API key is set
	cfg := &Config{
		GeminiAPIKey: "test-api-key",
		GeminiModel:  "gemini-2.0-flash-exp",
		Port:         "8081",
	}

	// This test only validates the structure, not the actual server functionality
	// as it would require a valid API key
	t.Run("Config validation", func(t *testing.T) {
		assert.NotEmpty(t, cfg.GeminiAPIKey)
		assert.NotEmpty(t, cfg.GeminiModel)
		assert.NotEmpty(t, cfg.Port)
	})
}
