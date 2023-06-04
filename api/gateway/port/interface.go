package port

import "mh-api/api/entity"

type MonsterInterface interface {
	GetAll() (entity.Monsters, error)
	GetById(id entity.MonsterId) (entity.Monster, error)
	Create(monsterJson entity.MonsterJson) error
}