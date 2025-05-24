package weapon

type ListWeaponsResponse struct {
	TotalCount int                    `json:"total_count,omitempty"`
	Limit      int                    `json:"limit,omitempty"`
	Offset     int                    `json:"offset,omitempty"`
	Weapons    []WeaponDetailResponse `json:"weapons,omitempty"`
}

type WeaponDetailResponse struct {
	WeaponID      string `json:"weapon_id,omitempty"`
	Name          string `json:"name,omitempty"`
	ImageURL      string `json:"image_url,omitempty"`
	Rare          string `json:"rare,omitempty"`
	Attack        string `json:"attack,omitempty"`
	ElementAttack string `json:"element_attack,omitempty"`
	Sharpness     string `json:"sharpness,omitempty"`
	Critical      string `json:"critical,omitempty"`
	Description   string `json:"description,omitempty"`
}
