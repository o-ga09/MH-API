package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"mh-api/internal/database/mysql"
	"mh-api/internal/service/monsters"
	"mh-api/internal/service/weapons"
	"mh-api/internal/service/items"
	"mh-api/internal/service/skills"
)

type MCPServer struct {
	monsterService monsters.IMonsterService
	weaponService  *weapons.WeaponService
	itemService    items.IitemService
	skillService   skills.ISkillService
}

func main() {
	ctx := context.Background()

	// Initialize repositories and services directly
	monsterRepo := mysql.NewMonsterRepository()
	monsterQS := mysql.NewmonsterQueryService()
	monsterService := monsters.NewMonsterService(monsterRepo, monsterQS)

	weaponQS := mysql.NewWeaponQueryService()
	weaponService := weapons.NewWeaponService(weaponQS)

	itemRepo := mysql.NewItemQueryService()
	itemService := items.NewService(itemRepo)

	skillRepo := mysql.NewSkillQueryService()
	skillService := skills.NewService(skillRepo)

	mcpServer := &MCPServer{
		monsterService: monsterService,
		weaponService:  weaponService,
		itemService:    itemService,
		skillService:   skillService,
	}

	s := server.NewStdioServer()

	err := s.SetRequestHandler("tools/list", mcpServer.handleToolsList)
	if err != nil {
		log.Fatal("Failed to set tools/list handler:", err)
	}

	err = s.SetRequestHandler("tools/call", mcpServer.handleToolsCall)
	if err != nil {
		log.Fatal("Failed to set tools/call handler:", err)
	}

	if err := s.Serve(); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}

func (m *MCPServer) handleToolsList(ctx context.Context, request mcp.Request) (interface{}, error) {
	tools := []mcp.Tool{
		{
			Name:        "get_monsters",
			Description: "Get a list of all monsters with their details",
			InputSchema: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"page": map[string]interface{}{
						"type":        "integer",
						"description": "Page number (optional, default: 1)",
						"minimum":     1,
					},
					"limit": map[string]interface{}{
						"type":        "integer", 
						"description": "Number of items per page (optional, default: 50)",
						"minimum":     1,
						"maximum":     100,
					},
				},
			},
		},
		{
			Name:        "get_monster_by_id",
			Description: "Get detailed information about a specific monster by ID",
			InputSchema: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"monster_id": map[string]interface{}{
						"type":        "string",
						"description": "The unique identifier of the monster",
					},
				},
				"required": []string{"monster_id"},
			},
		},
		{
			Name:        "get_weapons",
			Description: "Search weapons with various filters",
			InputSchema: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"weapon_id": map[string]interface{}{
						"type":        "string",
						"description": "Filter by weapon ID (optional)",
					},
					"name": map[string]interface{}{
						"type":        "string",
						"description": "Filter by weapon name (optional)",
					},
					"limit": map[string]interface{}{
						"type":        "integer",
						"description": "Number of items to return (optional, default: 50)",
						"minimum":     1,
						"maximum":     100,
					},
					"offset": map[string]interface{}{
						"type":        "integer",
						"description": "Number of items to skip (optional, default: 0)",
						"minimum":     0,
					},
				},
			},
		},
		{
			Name:        "get_items",
			Description: "Get a list of all items",
			InputSchema: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{},
			},
		},
		{
			Name:        "get_item_by_id",
			Description: "Get detailed information about a specific item by ID",
			InputSchema: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"item_id": map[string]interface{}{
						"type":        "string",
						"description": "The unique identifier of the item",
					},
				},
				"required": []string{"item_id"},
			},
		},
		{
			Name:        "get_items_by_monster",
			Description: "Get items that can be obtained from a specific monster",
			InputSchema: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"monster_id": map[string]interface{}{
						"type":        "string",
						"description": "The unique identifier of the monster",
					},
				},
				"required": []string{"monster_id"},
			},
		},
		{
			Name:        "get_skills",
			Description: "Get a list of all skills with their level details",
			InputSchema: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{},
			},
		},
		{
			Name:        "get_skill_by_id",
			Description: "Get detailed information about a specific skill by ID",
			InputSchema: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"skill_id": map[string]interface{}{
						"type":        "string",
						"description": "The unique identifier of the skill",
					},
				},
				"required": []string{"skill_id"},
			},
		},
	}

	return map[string]interface{}{
		"tools": tools,
	}, nil
}

func (m *MCPServer) handleToolsCall(ctx context.Context, request mcp.Request) (interface{}, error) {
	params := request.Params.(map[string]interface{})
	toolName := params["name"].(string)
	arguments := params["arguments"].(map[string]interface{})

	switch toolName {
	case "get_monsters":
		return m.getMonsters(ctx, arguments)
	case "get_monster_by_id":
		return m.getMonsterByID(ctx, arguments)
	case "get_weapons":
		return m.getWeapons(ctx, arguments)
	case "get_items":
		return m.getItems(ctx)
	case "get_item_by_id":
		return m.getItemByID(ctx, arguments)
	case "get_items_by_monster":
		return m.getItemsByMonster(ctx, arguments)
	case "get_skills":
		return m.getSkills(ctx)
	case "get_skill_by_id":
		return m.getSkillByID(ctx, arguments)
	default:
		return nil, fmt.Errorf("unknown tool: %s", toolName)
	}
}

func (m *MCPServer) getMonsters(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Note: Current monster service only supports fetching by ID
	// For listing all monsters, we'll return an informative message
	return map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": "Monster service currently supports fetching specific monsters by ID only. Use get_monster_by_id tool with a specific monster ID.",
			},
		},
	}, nil
}

func (m *MCPServer) getMonsterByID(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	monsterID, ok := args["monster_id"].(string)
	if !ok {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": "Error: monster_id is required and must be a string",
				},
			},
		}, nil
	}

	monsters, err := m.monsterService.FetchMonsterDetail(ctx, monsterID)
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Error retrieving monster with ID %s: %v", monsterID, err),
				},
			},
		}, nil
	}

	content, err := json.MarshalIndent(monsters, "", "  ")
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Error formatting monster data: %v", err),
				},
			},
		}, nil
	}

	return map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": string(content),
			},
		},
	}, nil
}

func (m *MCPServer) getWeapons(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	var weaponID, name *string
	var limit, offset *int

	if wid, ok := args["weapon_id"].(string); ok && wid != "" {
		weaponID = &wid
	}
	if n, ok := args["name"].(string); ok && n != "" {
		name = &n
	}
	if l, ok := args["limit"]; ok {
		if limitFloat, ok := l.(float64); ok {
			limitInt := int(limitFloat)
			limit = &limitInt
		}
	}
	if o, ok := args["offset"]; ok {
		if offsetFloat, ok := o.(float64); ok {
			offsetInt := int(offsetFloat)
			offset = &offsetInt
		}
	}

	// Create search params
	params := weapons.SearchWeaponsParams{
		WeaponID: weaponID,
		Name:     name,
		Limit:    limit,
		Offset:   offset,
	}

	weapons, err := m.weaponService.SearchWeapons(ctx, params)
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Error retrieving weapons: %v", err),
				},
			},
		}, nil
	}

	content, err := json.MarshalIndent(weapons, "", "  ")
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Error formatting weapons data: %v", err),
				},
			},
		}, nil
	}

	return map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": string(content),
			},
		},
	}, nil
}

func (m *MCPServer) getItems(ctx context.Context) (interface{}, error) {
	items, err := m.itemService.GetAllItems(ctx)
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Error retrieving items: %v", err),
				},
			},
		}, nil
	}

	content, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Error formatting items data: %v", err),
				},
			},
		}, nil
	}

	return map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": string(content),
			},
		},
	}, nil
}

func (m *MCPServer) getItemByID(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	itemID, ok := args["item_id"].(string)
	if !ok {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": "Error: item_id is required and must be a string",
				},
			},
		}, nil
	}

	// Convert string to int as the service expects an int
	itemIDInt, err := strconv.Atoi(itemID)
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Error: item_id must be a valid integer: %v", err),
				},
			},
		}, nil
	}

	item, err := m.itemService.GetItemByID(ctx, strconv.Itoa(itemIDInt))
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Error retrieving item with ID %s: %v", itemID, err),
				},
			},
		}, nil
	}

	content, err := json.MarshalIndent(item, "", "  ")
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Error formatting item data: %v", err),
				},
			},
		}, nil
	}

	return map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": string(content),
			},
		},
	}, nil
}

func (m *MCPServer) getItemsByMonster(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	monsterID, ok := args["monster_id"].(string)
	if !ok {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": "Error: monster_id is required and must be a string",
				},
			},
		}, nil
	}

	// Convert string to int as the service expects an int
	monsterIDInt, err := strconv.Atoi(monsterID)
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Error: monster_id must be a valid integer: %v", err),
				},
			},
		}, nil
	}

	items, err := m.itemService.GetItemByMonsterID(ctx, strconv.Itoa(monsterIDInt))
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Error retrieving items for monster ID %s: %v", monsterID, err),
				},
			},
		}, nil
	}

	content, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Error formatting items data: %v", err),
				},
			},
		}, nil
	}

	return map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": string(content),
			},
		},
	}, nil
}

func (m *MCPServer) getSkills(ctx context.Context) (interface{}, error) {
	skills, err := m.skillService.GetAllSkills(ctx)
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Error retrieving skills: %v", err),
				},
			},
		}, nil
	}

	content, err := json.MarshalIndent(skills, "", "  ")
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Error formatting skills data: %v", err),
				},
			},
		}, nil
	}

	return map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": string(content),
			},
		},
	}, nil
}

func (m *MCPServer) getSkillByID(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	skillID, ok := args["skill_id"].(string)
	if !ok {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": "Error: skill_id is required and must be a string",
				},
			},
		}, nil
	}

	skill, err := m.skillService.GetSkillByID(ctx, skillID)
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Error retrieving skill with ID %s: %v", skillID, err),
				},
			},
		}, nil
	}

	content, err := json.MarshalIndent(skill, "", "  ")
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Error formatting skill data: %v", err),
				},
			},
		}, nil
	}

	return map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": string(content),
			},
		},
	}, nil
}