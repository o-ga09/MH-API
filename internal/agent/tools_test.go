package agent

import (
"testing"

"github.com/stretchr/testify/assert"
"github.com/stretchr/testify/require"

"mh-api/internal/database/mysql"
)

func TestNewMonHunTools(t *testing.T) {
// Initialize repositories
monsterRepo := mysql.NewMonsterRepository()
weaponRepo := mysql.NewWeaponRepository()
itemRepo := mysql.NewItemQueryService()
skillRepo := mysql.NewSkillQueryService()

// Create tools
monHunTools := NewMonHunTools(monsterRepo, weaponRepo, itemRepo, skillRepo)
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
cfg := &Config{
GeminiAPIKey: "test-api-key",
GeminiModel:  "gemini-2.0-flash-exp",
Port:         "8081",
}

t.Run("Config validation", func(t *testing.T) {
assert.NotEmpty(t, cfg.GeminiAPIKey)
assert.NotEmpty(t, cfg.GeminiModel)
assert.NotEmpty(t, cfg.Port)
})
}
