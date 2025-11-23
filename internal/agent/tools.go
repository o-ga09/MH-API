package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"

	request "mh-api/internal/controller/monster"
	"mh-api/internal/service/items"
	"mh-api/internal/service/monsters"
	"mh-api/internal/service/skills"
	"mh-api/internal/service/weapons"
)

// MonHunTools provides tools for accessing MonHun API data
type MonHunTools struct {
	monsterService monsters.IMonsterService
	weaponService  *weapons.WeaponService
	itemService    items.IitemService
	skillService   skills.ISkillService
}

// NewMonHunTools creates a new MonHunTools instance
func NewMonHunTools(
	monsterService monsters.IMonsterService,
	weaponService *weapons.WeaponService,
	itemService items.IitemService,
	skillService skills.ISkillService,
) *MonHunTools {
	return &MonHunTools{
		monsterService: monsterService,
		weaponService:  weaponService,
		itemService:    itemService,
		skillService:   skillService,
	}
}

// GetTools returns all available tools for the agent
func (m *MonHunTools) GetTools() ([]tool.Tool, error) {
	tools := []tool.Tool{}

	// get_monsters tool
	getMonstersT, err := functiontool.New(functiontool.Config{
		Name:        "get_monsters",
		Description: "モンスターの一覧を取得します。名前やIDでフィルタリングでき、ページネーションもサポートしています。",
	}, func(ctx tool.Context, args GetMonstersInput) (string, error) {
		return m.getMonsters(ctx, args)
	})
	if err != nil {
		return nil, err
	}
	tools = append(tools, getMonstersT)

	// get_monster_by_id tool
	getMonsterByIDT, err := functiontool.New(functiontool.Config{
		Name:        "get_monster_by_id",
		Description: "指定されたIDのモンスターの詳細情報を取得します。",
	}, func(ctx tool.Context, args GetMonsterByIDInput) (string, error) {
		return m.getMonsterByID(ctx, args)
	})
	if err != nil {
		return nil, err
	}
	tools = append(tools, getMonsterByIDT)

	// get_weapons tool
	getWeaponsT, err := functiontool.New(functiontool.Config{
		Name:        "get_weapons",
		Description: "武器を検索します。様々なフィルタでの絞り込みが可能です。",
	}, func(ctx tool.Context, args GetWeaponsInput) (string, error) {
		return m.getWeapons(ctx, args)
	})
	if err != nil {
		return nil, err
	}
	tools = append(tools, getWeaponsT)

	// get_items tool
	getItemsT, err := functiontool.New(functiontool.Config{
		Name:        "get_items",
		Description: "全てのアイテムの一覧を取得します。",
	}, func(ctx tool.Context, args EmptyInput) (string, error) {
		return m.getItems(ctx)
	})
	if err != nil {
		return nil, err
	}
	tools = append(tools, getItemsT)

	// get_item_by_id tool
	getItemByIDT, err := functiontool.New(functiontool.Config{
		Name:        "get_item_by_id",
		Description: "指定されたIDのアイテムの詳細情報を取得します。",
	}, func(ctx tool.Context, args GetItemByIDInput) (string, error) {
		return m.getItemByID(ctx, args)
	})
	if err != nil {
		return nil, err
	}
	tools = append(tools, getItemByIDT)

	// get_items_by_monster tool
	getItemsByMonsterT, err := functiontool.New(functiontool.Config{
		Name:        "get_items_by_monster",
		Description: "指定されたモンスターから入手できるアイテムの一覧を取得します。",
	}, func(ctx tool.Context, args GetItemsByMonsterInput) (string, error) {
		return m.getItemsByMonster(ctx, args)
	})
	if err != nil {
		return nil, err
	}
	tools = append(tools, getItemsByMonsterT)

	// get_skills tool
	getSkillsT, err := functiontool.New(functiontool.Config{
		Name:        "get_skills",
		Description: "全てのスキルとそのレベル詳細の一覧を取得します。",
	}, func(ctx tool.Context, args EmptyInput) (string, error) {
		return m.getSkills(ctx)
	})
	if err != nil {
		return nil, err
	}
	tools = append(tools, getSkillsT)

	// get_skill_by_id tool
	getSkillByIDT, err := functiontool.New(functiontool.Config{
		Name:        "get_skill_by_id",
		Description: "指定されたIDのスキルの詳細情報を取得します。",
	}, func(ctx tool.Context, args GetSkillByIDInput) (string, error) {
		return m.getSkillByID(ctx, args)
	})
	if err != nil {
		return nil, err
	}
	tools = append(tools, getSkillByIDT)

	return tools, nil
}

// Input/Output types for tools
type GetMonstersInput struct {
	MonsterIDs string `json:"monster_ids,omitempty" jsonschema:"description=モンスターIDでフィルタリング（オプション、カンマ区切りで複数指定可能）"`
	Name       string `json:"name,omitempty" jsonschema:"description=モンスター名でフィルタリング（オプション、部分一致）"`
	Sort       string `json:"sort,omitempty" jsonschema:"description=ソート順（オプション、'asc' または 'desc'、デフォルトは 'asc'）,enum=asc,enum=desc"`
	Offset     int    `json:"offset,omitempty" jsonschema:"description=ページネーションのオフセット（オプション、デフォルト: 0）,minimum=0"`
	Limit      int    `json:"limit,omitempty" jsonschema:"description=1ページあたりの件数（オプション、デフォルト: 50、最大: 100）,minimum=1,maximum=100"`
}

type GetMonsterByIDInput struct {
	MonsterID string `json:"monster_id" jsonschema:"description=モンスターの一意な識別子,required"`
}

type GetWeaponsInput struct {
	WeaponID string `json:"weapon_id,omitempty" jsonschema:"description=武器IDでフィルタリング（オプション）"`
	Name     string `json:"name,omitempty" jsonschema:"description=武器名でフィルタリング（オプション）"`
	Limit    int    `json:"limit,omitempty" jsonschema:"description=取得する件数（オプション、デフォルト: 50）,minimum=1,maximum=100"`
	Offset   int    `json:"offset,omitempty" jsonschema:"description=スキップする件数（オプション、デフォルト: 0）,minimum=0"`
}

type GetItemByIDInput struct {
	ItemID string `json:"item_id" jsonschema:"description=アイテムの一意な識別子,required"`
}

type GetItemsByMonsterInput struct {
	MonsterID string `json:"monster_id" jsonschema:"description=モンスターの一意な識別子,required"`
}

type GetSkillByIDInput struct {
	SkillID string `json:"skill_id" jsonschema:"description=スキルの一意な識別子,required"`
}

type EmptyInput struct{}

// Tool handler functions
func (m *MonHunTools) getMonsters(ctx context.Context, args GetMonstersInput) (string, error) {
	offset := args.Offset
	limit := args.Limit
	if limit == 0 {
		limit = 50
	}
	if offset < 0 {
		return "", fmt.Errorf("offset must be greater than or equal to 0")
	}
	if limit < 1 || limit > 100 {
		return "", fmt.Errorf("limit must be between 1 and 100")
	}

	sort := args.Sort
	if sort == "" {
		sort = "asc"
	}

	param := request.RequestParam{
		MonsterIds:  args.MonsterIDs,
		MonsterName: args.Name,
		Sort:        sort,
		Limit:       limit,
		Offset:      (offset - 1) * limit,
	}
	ctx = context.WithValue(ctx, "param", param)

	monsters, err := m.monsterService.FetchMonsterDetail(ctx, "")
	if err != nil {
		return "", fmt.Errorf("error retrieving monsters: %w", err)
	}

	content, err := json.MarshalIndent(monsters, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error formatting monster data: %w", err)
	}

	if len(monsters.Monsters) == 0 {
		return "No monsters found", nil
	}

	return string(content), nil
}

func (m *MonHunTools) getMonsterByID(ctx context.Context, args GetMonsterByIDInput) (string, error) {
	if args.MonsterID == "" {
		return "", fmt.Errorf("monster_id is required and must be a string")
	}

	monsters, err := m.monsterService.FetchMonsterDetail(ctx, args.MonsterID)
	if err != nil {
		return "", fmt.Errorf("error retrieving monster with ID %s: %w", args.MonsterID, err)
	}

	content, err := json.MarshalIndent(monsters, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error formatting monster data: %w", err)
	}

	return string(content), nil
}

func (m *MonHunTools) getWeapons(ctx context.Context, args GetWeaponsInput) (string, error) {
	limit := args.Limit
	if limit == 0 {
		limit = 50
	}
	offset := args.Offset

	// Create search params
	params := weapons.SearchWeaponsParams{
		WeaponID: &args.WeaponID,
		Name:     &args.Name,
		Limit:    &limit,
		Offset:   &offset,
	}

	weapons, err := m.weaponService.SearchWeapons(ctx, params)
	if err != nil {
		return "", fmt.Errorf("error retrieving weapons: %w", err)
	}

	content, err := json.MarshalIndent(weapons, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error formatting weapons data: %w", err)
	}

	return string(content), nil
}

func (m *MonHunTools) getItems(ctx context.Context) (string, error) {
	items, err := m.itemService.GetAllItems(ctx)
	if err != nil {
		return "", fmt.Errorf("error retrieving items: %w", err)
	}

	content, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error formatting items data: %w", err)
	}

	return string(content), nil
}

func (m *MonHunTools) getItemByID(ctx context.Context, args GetItemByIDInput) (string, error) {
	if args.ItemID == "" {
		return "", fmt.Errorf("item_id is required and must be a string")
	}

	// Convert string to int as the service expects an int
	itemIDInt, err := strconv.Atoi(args.ItemID)
	if err != nil {
		return "", fmt.Errorf("item_id must be a valid integer: %w", err)
	}

	item, err := m.itemService.GetItemByID(ctx, strconv.Itoa(itemIDInt))
	if err != nil {
		return "", fmt.Errorf("error retrieving item with ID %s: %w", args.ItemID, err)
	}

	content, err := json.MarshalIndent(item, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error formatting item data: %w", err)
	}

	return string(content), nil
}

func (m *MonHunTools) getItemsByMonster(ctx context.Context, args GetItemsByMonsterInput) (string, error) {
	if args.MonsterID == "" {
		return "", fmt.Errorf("monster_id is required and must be a string")
	}

	// Convert string to int as the service expects an int
	monsterIDInt, err := strconv.Atoi(args.MonsterID)
	if err != nil {
		return "", fmt.Errorf("monster_id must be a valid integer: %w", err)
	}

	items, err := m.itemService.GetItemByMonsterID(ctx, strconv.Itoa(monsterIDInt))
	if err != nil {
		return "", fmt.Errorf("error retrieving items for monster ID %s: %w", args.MonsterID, err)
	}

	content, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error formatting items data: %w", err)
	}

	return string(content), nil
}

func (m *MonHunTools) getSkills(ctx context.Context) (string, error) {
	skills, err := m.skillService.GetAllSkills(ctx)
	if err != nil {
		return "", fmt.Errorf("error retrieving skills: %w", err)
	}

	content, err := json.MarshalIndent(skills, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error formatting skills data: %w", err)
	}

	return string(content), nil
}

func (m *MonHunTools) getSkillByID(ctx context.Context, args GetSkillByIDInput) (string, error) {
	if args.SkillID == "" {
		return "", fmt.Errorf("skill_id is required and must be a string")
	}

	skill, err := m.skillService.GetSkillByID(ctx, args.SkillID)
	if err != nil {
		return "", fmt.Errorf("error retrieving skill with ID %s: %w", args.SkillID, err)
	}

	content, err := json.MarshalIndent(skill, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error formatting skill data: %w", err)
	}

	return string(content), nil
}
