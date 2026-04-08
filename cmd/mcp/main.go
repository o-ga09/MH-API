package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"mh-api/pkg/config"
	"mh-api/pkg/profiler"

	request "mh-api/internal/controller/monster"
	"mh-api/internal/database/mysql"
	"mh-api/internal/service/items"
	"mh-api/internal/service/monsters"
	"mh-api/internal/service/skills"
	"mh-api/internal/service/weapons"
)

type MCPServer struct {
	monsterService monsters.IMonsterService
	weaponService  *weapons.WeaponService
	itemService    items.IitemService
	skillService   skills.ISkillService
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	stopProfiler := profiler.StartPyroscope(cfg, "mh-mcp")
	defer stopProfiler()

	monsterRepo := mysql.NewMonsterRepository()
	monsterQS := mysql.NewmonsterQueryService()
	monsterService := monsters.NewMonsterService(monsterRepo, monsterQS)

	weaponQS := mysql.NewWeaponQueryService()
	weaponService := weapons.NewWeaponService(weaponQS)

	itemRepo := mysql.NewItemQueryService()
	itemService := items.NewService(monsterQS, itemRepo)

	skillRepo := mysql.NewSkillQueryService()
	skillService := skills.NewService(skillRepo)

	m := &MCPServer{
		monsterService: monsterService,
		weaponService:  weaponService,
		itemService:    itemService,
		skillService:   skillService,
	}

	s := mcp.NewServer(&mcp.Implementation{Name: "MH-API MCP Server", Version: "1.0.0"}, nil)
	m.registerTools(s)

	handler := mcp.NewSSEHandler(func(r *http.Request) *mcp.Server { return s }, nil)
	http.Handle("/sse", handler)
	log.Print("Starting SSE server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// getMonstersInput はモンスター一覧取得ツールの入力パラメータです。
type getMonstersInput struct {
	MonsterIDs      string `json:"monster_ids"      jsonschema:"Filter by monster ID (optional, can be a comma-separated list of IDs)"`
	Name            string `json:"name"             jsonschema:"Filter by monster name (optional, supports partial matches)"`
	UsageElement    string `json:"usage_element"    jsonschema:"Filter by monster's usage element (optional, e.g. Fire/Water/Lightning/Ice/Dragon)"`
	WeaknessElement string `json:"weakness_element" jsonschema:"Filter by monster's weakness element (optional, e.g. Fire/Water/Lightning/Ice/Dragon)"`
	Sort            string `json:"sort"             jsonschema:"Sort order: asc or desc (default: asc)"`
	Offset          int    `json:"offset"           jsonschema:"Offset for pagination (optional, default: 0)"`
	Limit           int    `json:"limit"            jsonschema:"Number of items per page (optional, default: 50, max: 100)"`
}

type getMonsterByIDInput struct {
	MonsterID string `json:"monster_id" jsonschema:"The unique identifier of the monster"`
}

type getWeaponsInput struct {
	WeaponID string `json:"weapon_id" jsonschema:"Filter by weapon ID (optional)"`
	Name     string `json:"name"      jsonschema:"Filter by weapon name (optional)"`
	Limit    int    `json:"limit"     jsonschema:"Number of items to return (optional, default: 50, max: 100)"`
	Offset   int    `json:"offset"    jsonschema:"Number of items to skip (optional, default: 0)"`
}

type getItemByIDInput struct {
	ItemID string `json:"item_id" jsonschema:"The unique identifier of the item"`
}

type getItemsByMonsterInput struct {
	MonsterID string `json:"monster_id" jsonschema:"The unique identifier of the monster"`
}

type getSkillByIDInput struct {
	SkillID string `json:"skill_id" jsonschema:"The unique identifier of the skill"`
}

func (m *MCPServer) registerTools(s *mcp.Server) {
	mcp.AddTool(s, &mcp.Tool{Name: "get_monsters", Description: "Get a list of all monsters with their details"}, m.getMonsters)
	mcp.AddTool(s, &mcp.Tool{Name: "get_monster_by_id", Description: "Get detailed information about a specific monster by ID"}, m.getMonsterByID)
	mcp.AddTool(s, &mcp.Tool{Name: "get_weapons", Description: "Search weapons with various filters"}, m.getWeapons)
	mcp.AddTool(s, &mcp.Tool{Name: "get_items", Description: "Get a list of all items"}, m.getItems)
	mcp.AddTool(s, &mcp.Tool{Name: "get_item_by_id", Description: "Get detailed information about a specific item by ID"}, m.getItemByID)
	mcp.AddTool(s, &mcp.Tool{Name: "get_items_by_monster", Description: "Get items that can be obtained from a specific monster"}, m.getItemsByMonster)
	mcp.AddTool(s, &mcp.Tool{Name: "get_skills", Description: "Get a list of all skills with their level details"}, m.getSkills)
	mcp.AddTool(s, &mcp.Tool{Name: "get_skill_by_id", Description: "Get detailed information about a specific skill by ID"}, m.getSkillByID)
}

func textResult(text string) (*mcp.CallToolResult, any, error) {
	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: text}},
	}, nil, nil
}

func (m *MCPServer) getMonsters(ctx context.Context, _ *mcp.CallToolRequest, in getMonstersInput) (*mcp.CallToolResult, any, error) {
	ctx = mysql.New(ctx)

	offset := in.Offset
	limit := in.Limit
	if limit == 0 {
		limit = 50
	}
	if offset < 0 {
		return textResult("Error: offset must be greater than or equal to 0")
	}
	if limit < 1 || limit > 100 {
		return textResult("Error: limit must be between 1 and 100")
	}
	sort := in.Sort
	if sort == "" {
		sort = "asc"
	}

	param := request.RequestParam{
		MonsterIds:      in.MonsterIDs,
		MonsterName:     in.Name,
		UsageElement:    in.UsageElement,
		WeaknessElement: in.WeaknessElement,
		Sort:            sort,
		Limit:           limit,
		Offset:          (offset - 1) * limit,
	}
	ctx = context.WithValue(ctx, request.CtxParamKey, param)

	result, err := m.monsterService.FetchMonsterDetail(ctx, "")
	if err != nil {
		return textResult(fmt.Sprintf("Error retrieving monsters: %v", err))
	}
	if len(result.Monsters) == 0 {
		return textResult("No monsters found")
	}
	content, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return textResult(fmt.Sprintf("Error formatting monster data: %v", err))
	}
	return textResult(string(content))
}

func (m *MCPServer) getMonsterByID(ctx context.Context, _ *mcp.CallToolRequest, in getMonsterByIDInput) (*mcp.CallToolResult, any, error) {
	ctx = mysql.New(ctx)

	if in.MonsterID == "" {
		return textResult("Error: monster_id is required and must be a string")
	}

	result, err := m.monsterService.FetchMonsterDetail(ctx, in.MonsterID)
	if err != nil {
		return textResult(fmt.Sprintf("Error retrieving monster with ID %s: %v", in.MonsterID, err))
	}
	content, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return textResult(fmt.Sprintf("Error formatting monster data: %v", err))
	}
	return textResult(string(content))
}

func (m *MCPServer) getWeapons(ctx context.Context, _ *mcp.CallToolRequest, in getWeaponsInput) (*mcp.CallToolResult, any, error) {
	ctx = mysql.New(ctx)

	limit := in.Limit
	if limit == 0 {
		limit = 50
	}
	offset := in.Offset

	params := weapons.SearchWeaponsParams{
		WeaponID: &in.WeaponID,
		Name:     &in.Name,
		Limit:    &limit,
		Offset:   &offset,
	}

	result, err := m.weaponService.SearchWeapons(ctx, params)
	if err != nil {
		return textResult(fmt.Sprintf("Error retrieving weapons: %v", err))
	}
	content, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return textResult(fmt.Sprintf("Error formatting weapons data: %v", err))
	}
	return textResult(string(content))
}

func (m *MCPServer) getItems(ctx context.Context, _ *mcp.CallToolRequest, _ any) (*mcp.CallToolResult, any, error) {
	ctx = mysql.New(ctx)

	result, err := m.itemService.GetAllItems(ctx)
	if err != nil {
		return textResult(fmt.Sprintf("Error retrieving items: %v", err))
	}
	content, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return textResult(fmt.Sprintf("Error formatting items data: %v", err))
	}
	return textResult(string(content))
}

func (m *MCPServer) getItemByID(ctx context.Context, _ *mcp.CallToolRequest, in getItemByIDInput) (*mcp.CallToolResult, any, error) {
	ctx = mysql.New(ctx)

	if in.ItemID == "" {
		return textResult("Error: item_id is required and must be a string")
	}
	itemIDInt, err := strconv.Atoi(in.ItemID)
	if err != nil {
		return textResult(fmt.Sprintf("Error: item_id must be a valid integer: %v", err))
	}
	result, err := m.itemService.GetItemByID(ctx, strconv.Itoa(itemIDInt))
	if err != nil {
		return textResult(fmt.Sprintf("Error retrieving item with ID %s: %v", in.ItemID, err))
	}
	content, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return textResult(fmt.Sprintf("Error formatting item data: %v", err))
	}
	return textResult(string(content))
}

func (m *MCPServer) getItemsByMonster(ctx context.Context, _ *mcp.CallToolRequest, in getItemsByMonsterInput) (*mcp.CallToolResult, any, error) {
	ctx = mysql.New(ctx)

	if in.MonsterID == "" {
		return textResult("Error: monster_id is required and must be a string")
	}
	monsterIDInt, err := strconv.Atoi(in.MonsterID)
	if err != nil {
		return textResult(fmt.Sprintf("Error: monster_id must be a valid integer: %v", err))
	}
	result, err := m.itemService.GetItemByMonsterID(ctx, strconv.Itoa(monsterIDInt))
	if err != nil {
		return textResult(fmt.Sprintf("Error retrieving items for monster ID %s: %v", in.MonsterID, err))
	}
	content, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return textResult(fmt.Sprintf("Error formatting items data: %v", err))
	}
	return textResult(string(content))
}

func (m *MCPServer) getSkills(ctx context.Context, _ *mcp.CallToolRequest, _ any) (*mcp.CallToolResult, any, error) {
	ctx = mysql.New(ctx)

	result, err := m.skillService.GetAllSkills(ctx)
	if err != nil {
		return textResult(fmt.Sprintf("Error retrieving skills: %v", err))
	}
	content, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return textResult(fmt.Sprintf("Error formatting skills data: %v", err))
	}
	return textResult(string(content))
}

func (m *MCPServer) getSkillByID(ctx context.Context, _ *mcp.CallToolRequest, in getSkillByIDInput) (*mcp.CallToolResult, any, error) {
	ctx = mysql.New(ctx)

	if in.SkillID == "" {
		return textResult("Error: skill_id is required and must be a string")
	}
	result, err := m.skillService.GetSkillByID(ctx, in.SkillID)
	if err != nil {
		return textResult(fmt.Sprintf("Error retrieving skill with ID %s: %v", in.SkillID, err))
	}
	content, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return textResult(fmt.Sprintf("Error formatting skill data: %v", err))
	}
	return textResult(string(content))
}
