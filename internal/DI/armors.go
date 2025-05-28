package DI

import (
	"mh-api/internal/controller/armor"
	"mh-api/internal/database/mysql"
	"mh-api/internal/service/armors"
)

func InitArmorHandler() *armor.ArmorHandler {
	armorQueryService := mysql.NewArmorQueryService()
	armorService := armors.NewArmorService(armorQueryService)
	armorHandler := armor.NewArmorHandler(armorService)
	return armorHandler
}