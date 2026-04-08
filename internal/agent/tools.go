package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"

	"mh-api/internal/domain/items"
	"mh-api/internal/domain/monsters"
	"mh-api/internal/domain/skills"
	"mh-api/internal/domain/weapons"
)

// MonHunTools provides tools for accessing MonHun API data
type MonHunTools struct {
	monsterRepo monsters.Repository
	weaponRepo  weapons.Repository
	itemRepo    items.Repository
	skillRepo   skills.Repository
}

// NewMonHunTools creates a new MonHunTools instance
func NewMonHunTools(
	monsterRepo monsters.Repository,
	weaponRepo weapons.Repository,
	itemRepo items.Repository,
	skillRepo skills.Repository,
) *MonHunTools {
	return &MonHunTools{
		monsterRepo: monsterRepo,
		weaponRepo:  weaponRepo,
		itemRepo:    itemRepo,
		skillRepo:   skillRepo,
	}
}

// GetTools returns all available tools for the agent
func (m *MonHunTools) GetTools() ([]tool.Tool, error) {
	tools := []tool.Tool{}

	getMonstersT, err := functiontool.New(functiontool.Config{
		Name:        "get_monsters",
		Description: "モンスターの一覧を取得します。名前・ID・属性（usage_element）・弱点属性（weakness_element）でフィルタリングでき、ページネーションもサポートしています。属性値の例: 火, 水, 雷, 氷, 龍",
	}, func(ctx tool.Context, args GetMonstersInput) (string, error) {
		return m.getMonsters(ctx, args)
	})
	if err != nil {
		return nil, err
	}
	tools = append(tools, getMonstersT)

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

// Input types for tools
type GetMonstersInput struct {
	MonsterIDs      string `json:"monster_ids,omitempty"`
	Name            string `json:"name,omitempty"`
	UsageElement    string `json:"usage_element,omitempty"`
	WeaknessElement string `json:"weakness_element,omitempty"`
	Sort            string `json:"sort,omitempty"`
	Offset          int    `json:"offset,omitempty"`
	Limit           int    `json:"limit,omitempty"`
}

type GetMonsterByIDInput struct {
	MonsterID string `json:"monster_id"`
}

type GetWeaponsInput struct {
	WeaponID string `json:"weapon_id,omitempty"`
	Name     string `json:"name,omitempty"`
	Limit    int    `json:"limit,omitempty"`
	Offset   int    `json:"offset,omitempty"`
}

type GetItemByIDInput struct {
	ItemID string `json:"item_id"`
}

type GetItemsByMonsterInput struct {
	MonsterID string `json:"monster_id"`
}

type GetSkillByIDInput struct {
	SkillID string `json:"skill_id"`
}

type EmptyInput struct{}

func (m *MonHunTools) getMonsters(ctx context.Context, args GetMonstersInput) (string, error) {
	limit := args.Limit
	if limit == 0 {
		limit = 50
	}
	if args.Offset < 0 {
		return "", fmt.Errorf("offset must be greater than or equal to 0")
	}
	if limit < 1 || limit > 100 {
		return "", fmt.Errorf("limit must be between 1 and 100")
	}

	sort := args.Sort
	if sort == "" {
		sort = "asc"
	}

	params := monsters.SearchParams{
		MonsterIds:      args.MonsterIDs,
		MonsterName:     args.Name,
		UsageElement:    args.UsageElement,
		WeaknessElement: args.WeaknessElement,
		Sort:            sort,
		Limit:           limit,
		Offset:          args.Offset,
	}

	result, err := m.monsterRepo.FindAll(ctx, params)
	if err != nil {
		return "", fmt.Errorf("error retrieving monsters: %w", err)
	}

	if len(result.Monsters) == 0 {
		return "No monsters found", nil
	}

	content, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error formatting monster data: %w", err)
	}
	return string(content), nil
}

func (m *MonHunTools) getMonsterByID(ctx context.Context, args GetMonsterByIDInput) (string, error) {
	if args.MonsterID == "" {
		return "", fmt.Errorf("monster_id is required and must be a string")
	}

	monster, err := m.monsterRepo.FindById(ctx, args.MonsterID)
	if err != nil {
		return "", fmt.Errorf("error retrieving monster with ID %s: %w", args.MonsterID, err)
	}

	content, err := json.MarshalIndent(monster, "", "  ")
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

	params := weapons.SearchParams{
		WeaponID: &args.WeaponID,
		Name:     &args.Name,
		Limit:    &limit,
		Offset:   &offset,
	}

	result, err := m.weaponRepo.Find(ctx, params)
	if err != nil {
		return "", fmt.Errorf("error retrieving weapons: %w", err)
	}

	content, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error formatting weapons data: %w", err)
	}
	return string(content), nil
}

func (m *MonHunTools) getItems(ctx context.Context) (string, error) {
	itemList, err := m.itemRepo.FindAll(ctx)
	if err != nil {
		return "", fmt.Errorf("error retrieving items: %w", err)
	}

	content, err := json.MarshalIndent(itemList, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error formatting items data: %w", err)
	}
	return string(content), nil
}

func (m *MonHunTools) getItemByID(ctx context.Context, args GetItemByIDInput) (string, error) {
	if args.ItemID == "" {
		return "", fmt.Errorf("item_id is required and must be a string")
	}

	_, err := strconv.Atoi(args.ItemID)
	if err != nil {
		return "", fmt.Errorf("item_id must be a valid integer: %w", err)
	}

	item, err := m.itemRepo.FindByID(ctx, args.ItemID)
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

	_, err := strconv.Atoi(args.MonsterID)
	if err != nil {
		return "", fmt.Errorf("monster_id must be a valid integer: %w", err)
	}

	itemList, err := m.itemRepo.FindByMonsterID(ctx, args.MonsterID)
	if err != nil {
		return "", fmt.Errorf("error retrieving items for monster ID %s: %w", args.MonsterID, err)
	}

	content, err := json.MarshalIndent(itemList, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error formatting items data: %w", err)
	}
	return string(content), nil
}

func (m *MonHunTools) getSkills(ctx context.Context) (string, error) {
	skillList, err := m.skillRepo.FindAll(ctx)
	if err != nil {
		return "", fmt.Errorf("error retrieving skills: %w", err)
	}

	content, err := json.MarshalIndent(skillList, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error formatting skills data: %w", err)
	}
	return string(content), nil
}

func (m *MonHunTools) getSkillByID(ctx context.Context, args GetSkillByIDInput) (string, error) {
	if args.SkillID == "" {
		return "", fmt.Errorf("skill_id is required and must be a string")
	}

	skill, err := m.skillRepo.FindById(ctx, args.SkillID)
	if err != nil {
		return "", fmt.Errorf("error retrieving skill with ID %s: %w", args.SkillID, err)
	}

	content, err := json.MarshalIndent(skill, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error formatting skill data: %w", err)
	}
	return string(content), nil
}
