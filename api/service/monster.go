package service

import (
	"mh-api/api/entity"
	"mh-api/api/gateway/port"
)

type MonsterService struct {
	monsterInterface port.MonsterInterface
}

func ProvideMonsterDriver(monsterGateway port.MonsterInterface) MonsterService {
	return MonsterService{monsterInterface: monsterGateway}
}

func (s MonsterService) GetAll() (entity.Monsters, error) {
	res, err := s.monsterInterface.GetAll()
	if err != nil {
		return entity.Monsters{}, err
	}
	return res,nil
}