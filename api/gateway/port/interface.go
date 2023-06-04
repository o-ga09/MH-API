package port

import "mh-api/api/entity"

type MonsterInterface interface {
	GetAll() (entity.Monsters, error)
}