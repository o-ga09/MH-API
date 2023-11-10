package service

import (
	"mh-api/api/entity"
	"mh-api/api/service/port"
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
	return res, nil
}

func (s MonsterService) GetById(id entity.MonsterId) (entity.Monster, error) {
	res, err := s.monsterInterface.GetById(id)
	if err != nil {
		return entity.Monster{}, err
	}
	return res, nil
}

func (s MonsterService) Create(monsterJson entity.MonsterJson) error {
	err := s.monsterInterface.Create(monsterJson)
	return err
}

func (s MonsterService) Update(id entity.MonsterId, monsterJson entity.MonsterJson) error {
	err := s.monsterInterface.Update(id, monsterJson)
	return err
}

func (s MonsterService) Delete(id entity.MonsterId) error {
	err := s.monsterInterface.Delete(id)
	return err
}
