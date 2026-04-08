package di

import (
"mh-api/internal/controller/armor"
"mh-api/internal/database/mysql"
)

func InitArmorHandler() *armor.ArmorHandler {
repo := mysql.NewArmorRepository()
return armor.NewArmorHandler(repo)
}
