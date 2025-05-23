package weapons

import (
	"mh-api/internal/domain/weapons"
)

type WeaponData struct {
	Attack        string `json:"attack"`
	Critical      string `json:"critical"`
	Description   string `json:"description"`
	ElementAttack string `json:"element_attack"`
	ImageURL      string `json:"image_url"`
	MonsterID     string `json:"monster_id"`
	Name          string `json:"name"`
	Rare          string `json:"rare"`
	Sharpness     string `json:"sharpness"`
}

type ListWeaponsResponse struct {
	Weapons    []WeaponData `json:"weapons"`
	TotalCount int          `json:"total_count"`
	Offset     int          `json:"offset"`
	Limit      int          `json:"limit"`
}
type SearchWeaponsParams struct {
	Limit     *int
	Offset    *int
	Sort      *string
	Order     *int
	MonsterID *string
	Name      *string
	NameKana  *string
}

func ToWeaponData(weapon *weapons.Weapon) WeaponData {
	return WeaponData{
		MonsterID:     weapon.GetID(),
		Name:          weapon.GetName(),
		ImageURL:      weapon.GetURL(),
		Rare:          weapon.GetRERATY(),
		Attack:        weapon.GetAttack(),
		ElementAttack: weapon.GetElementAttack(),
		Sharpness:     weapon.GetSharpness(),
		Critical:      weapon.GetCritical(),
		Description:   weapon.GetDescription(),
	}
}

func ToWeaponDataList(domainWeapons []*weapons.Weapon) []WeaponData {
	weaponDataList := make([]WeaponData, len(domainWeapons))
	for i, dw := range domainWeapons {
		weaponDataList[i] = ToWeaponData(dw)
	}
	return weaponDataList
}
