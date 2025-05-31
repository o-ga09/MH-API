package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	request "mh-api/internal/controller/monster"
	"mh-api/internal/database/mysql"
	"mh-api/internal/service/items"
	"mh-api/internal/service/monsters"
	"mh-api/internal/service/skills"
	"mh-api/internal/service/weapons"
)

type contextKey string

const paramKey contextKey = "param"

type MCPServer struct {
	monsterService monsters.IMonsterService
	weaponService  *weapons.WeaponService
	itemService    items.IitemService
	skillService   skills.ISkillService
}

func main() {
	monsterRepo := mysql.NewMonsterRepository()
	monsterQS := mysql.NewmonsterQueryService()
	monsterService := monsters.NewMonsterService(monsterRepo, monsterQS)

	weaponQS := mysql.NewWeaponQueryService()
	weaponService := weapons.NewWeaponService(weaponQS)

	itemRepo := mysql.NewItemQueryService()
	itemService := items.NewService(monsterQS, itemRepo)

	skillRepo := mysql.NewSkillQueryService()
	skillService := skills.NewService(skillRepo)

	mcpServer := &MCPServer{
		monsterService: monsterService,
		weaponService:  weaponService,
		itemService:    itemService,
		skillService:   skillService,
	}

	s := server.NewMCPServer(
		"Calculator Demo",
		"1.0.0",
		server.WithToolCapabilities(false),
		server.WithRecovery(),
	)

	tools := mcpServer.AddTools()
	s.AddTools(tools...)

	sse := server.NewSSEServer(s)
	log.Print("Starting SSE server on :8080")
	if err := sse.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (m *MCPServer) AddTools() []server.ServerTool {
	srvTools := []server.ServerTool{
		{
			Tool: mcp.Tool{
				Name:        "get_monsters",
				Description: "Get a list of all monsters with their details",
				InputSchema: mcp.ToolInputSchema{
					Type: "object",
					Properties: map[string]interface{}{
						"monster_ids": map[string]interface{}{
							"type":        "string",
							"description": "Filter by monster ID (optional, can be a comma-separated list of IDs)",
						},
						"name": map[string]interface{}{
							"type":        "string",
							"description": "Filter by monster name (optional, supports partial matches)",
						},
						"sort": map[string]interface{}{
							"type":        "string",
							"description": "Sort order for the results (optional, 'asc' for ascending, 'desc' for descending, default is 'asc')",
							"enum":        []string{"asc", "desc"},
						},
						"offset": map[string]interface{}{
							"type":        "integer",
							"description": "Offset for pagination (optional, default: 0)",
							"minimum":     0,
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
			Handler: m.getMonsters,
		},
		{
			Tool: mcp.Tool{
				Name:        "get_monster_by_id",
				Description: "Get detailed information about a specific monster by ID",
				InputSchema: mcp.ToolInputSchema{
					Type: "object",
					Properties: map[string]interface{}{
						"monster_id": map[string]interface{}{
							"type":        "string",
							"description": "The unique identifier of the monster",
						},
					},
					Required: []string{"monster_id"},
				},
			},
			Handler: m.getMonsterByID,
		},
		{
			Tool: mcp.Tool{
				Name:        "get_weapons",
				Description: "Search weapons with various filters",
				InputSchema: mcp.ToolInputSchema{
					Type: "object",
					Properties: map[string]interface{}{
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
			Handler: m.getWeapons,
		},
		{
			Tool: mcp.Tool{
				Name:        "get_items",
				Description: "Get a list of all items",
				InputSchema: mcp.ToolInputSchema{
					Type:       "object",
					Properties: map[string]interface{}{},
				},
			},
			Handler: m.getItems,
		},
		{
			Tool: mcp.Tool{

				Name:        "get_item_by_id",
				Description: "Get detailed information about a specific item by ID",
				InputSchema: mcp.ToolInputSchema{
					Type: "object",
					Properties: map[string]interface{}{
						"item_id": map[string]interface{}{
							"type":        "string",
							"description": "The unique identifier of the item",
						},
					},
					Required: []string{"item_id"},
				},
			},
			Handler: m.getItemByID,
		},
		{
			Tool: mcp.Tool{
				Name:        "get_items_by_monster",
				Description: "Get items that can be obtained from a specific monster",
				InputSchema: mcp.ToolInputSchema{
					Type: "object",
					Properties: map[string]interface{}{
						"monster_id": map[string]interface{}{
							"type":        "string",
							"description": "The unique identifier of the monster",
						},
					},
					Required: []string{"monster_id"},
				},
			},
			Handler: m.getItemsByMonster,
		},
		{
			Tool: mcp.Tool{
				Name:        "get_skills",
				Description: "Get a list of all skills with their level details",
				InputSchema: mcp.ToolInputSchema{
					Type:       "object",
					Properties: map[string]interface{}{},
				},
			},
			Handler: m.getSkills,
		},
		{
			Tool: mcp.Tool{
				Name:        "get_skill_by_id",
				Description: "Get detailed information about a specific skill by ID",
				InputSchema: mcp.ToolInputSchema{
					Type: "object",
					Properties: map[string]interface{}{
						"skill_id": map[string]interface{}{
							"type":        "string",
							"description": "The unique identifier of the skill",
						},
					},
					Required: []string{"skill_id"},
				},
			},
			Handler: m.getSkillByID,
		},
	}

	return srvTools
}

func (m *MCPServer) getMonsters(ctx context.Context, args mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	offset := args.GetInt("offset", 0)
	limit := args.GetInt("limit", 50)
	if offset < 0 {
		return mcp.NewToolResultText("Error: offset must be greater than or equal to 0"), nil
	}
	if limit < 1 || limit > 100 {
		return mcp.NewToolResultText("Error: limit must be between 1 and 100"), nil
	}
	monsterIDs := args.GetString("monster_ids", "")
	name := args.GetString("name", "")
	sort := args.GetString("sort", "asc")

	param := request.RequestParam{
		MonsterIds:  monsterIDs,
		MonsterName: name,
		Sort:        sort,
		Limit:       limit,
		Offset:      (offset - 1) * limit,
	}
	ctx = context.WithValue(ctx, paramKey, param)

	monsters, err := m.monsterService.FetchMonsterDetail(ctx, "")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error retrieving monsters: %v", err)), nil
	}
	content, err := json.MarshalIndent(monsters, "", "  ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error formatting monster data: %v", err)), nil
	}
	if len(monsters) == 0 {
		return mcp.NewToolResultText("No monsters found"), nil
	}

	return mcp.NewToolResultText(string(content)), nil
}

func (m *MCPServer) getMonsterByID(ctx context.Context, args mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	monsterID := args.GetString("monster_id", "")
	if monsterID == "" {
		return mcp.NewToolResultText("Error: monster_id is required and must be a string"), nil
	}

	monsters, err := m.monsterService.FetchMonsterDetail(ctx, monsterID)
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error retrieving monster with ID %s: %v", monsterID, err)), nil
	}

	content, err := json.MarshalIndent(monsters, "", "  ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error formatting monster data: %v", err)), nil
	}

	return mcp.NewToolResultText(string(content)), nil
}

func (m *MCPServer) getWeapons(ctx context.Context, args mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	wid := args.GetString("weapon_id", "")
	n := args.GetString("name", "")
	l := args.GetString("limit", "50")
	o := args.GetString("offset", "0")

	limit, err := strconv.Atoi(l)
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error parsing limit: %v", err)), nil
	}

	offset, err := strconv.Atoi(o)
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error parsing offset: %v", err)), nil
	}

	// Create search params
	params := weapons.SearchWeaponsParams{
		WeaponID: &wid,
		Name:     &n,
		Limit:    &limit,
		Offset:   &offset,
	}

	weapons, err := m.weaponService.SearchWeapons(ctx, params)
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error retrieving weapons: %v", err)), nil
	}

	content, err := json.MarshalIndent(weapons, "", "  ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error formatting weapons data: %v", err)), nil
	}

	return mcp.NewToolResultText(string(content)), nil
}

func (m *MCPServer) getItems(ctx context.Context, args mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	items, err := m.itemService.GetAllItems(ctx)
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error retrieving items: %v", err)), nil
	}

	content, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error formatting items data: %v", err)), nil
	}

	return mcp.NewToolResultText(string(content)), nil
}

func (m *MCPServer) getItemByID(ctx context.Context, args mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	itemID := args.GetString("item_id", "")
	if itemID == "" {
		return mcp.NewToolResultText("Error: item_id is required and must be a string"), nil
	}

	// Convert string to int as the service expects an int
	itemIDInt, err := strconv.Atoi(itemID)
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error: item_id must be a valid integer: %v", err)), nil
	}

	item, err := m.itemService.GetItemByID(ctx, strconv.Itoa(itemIDInt))
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error retrieving item with ID %s: %v", itemID, err)), nil
	}

	content, err := json.MarshalIndent(item, "", "  ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error formatting item data: %v", err)), nil
	}

	return mcp.NewToolResultText(string(content)), nil
}

func (m *MCPServer) getItemsByMonster(ctx context.Context, args mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	monsterID := args.GetString("monster_id", "")
	if monsterID == "" {
		return mcp.NewToolResultText("Error: monster_id is required and must be a string"), nil
	}

	// Convert string to int as the service expects an int
	monsterIDInt, err := strconv.Atoi(monsterID)
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error: monster_id must be a valid integer: %v", err)), nil
	}

	items, err := m.itemService.GetItemByMonsterID(ctx, strconv.Itoa(monsterIDInt))
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error retrieving items for monster ID %s: %v", monsterID, err)), nil
	}

	content, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error formatting items data: %v", err)), nil
	}

	return mcp.NewToolResultText(string(content)), nil
}

func (m *MCPServer) getSkills(ctx context.Context, args mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	skills, err := m.skillService.GetAllSkills(ctx)
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error retrieving skills: %v", err)), nil
	}

	content, err := json.MarshalIndent(skills, "", "  ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error formatting skills data: %v", err)), nil
	}

	return mcp.NewToolResultText(string(content)), nil
}

func (m *MCPServer) getSkillByID(ctx context.Context, args mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	skillID := args.GetString("skill_id", "")
	if skillID == "" {
		return mcp.NewToolResultText("Error: skill_id is required and must be a string"), nil
	}

	skill, err := m.skillService.GetSkillByID(ctx, skillID)
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error retrieving skill with ID %s: %v", skillID, err)), nil
	}

	content, err := json.MarshalIndent(skill, "", "  ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error formatting skill data: %v", err)), nil
	}

	return mcp.NewToolResultText(string(content)), nil
}
