package gateway

import (
	"errors"
	"fmt"
	"mh-api/api/entity"
	"mh-api/api/gateway/repository"
)

type MonsterGateway struct {
	monsterDriver repository.MonsterDriver
}

func (g MonsterGateway) GetAll() (entity.Monsters, error) {
	res := g.monsterDriver.GetAll()

	var result entity.Monsters

	if len(res) == 0 {
		return entity.Monsters{}, errors.New("NOT FOUND")
	}

	for _, r := range res {
		weak_a := entity.Weakness_attack{
			FrontLegs: entity.AttackCatetgory(r.Weakness_attack.FrontLegs),
			HindLegs:  entity.AttackCatetgory(r.Weakness_attack.HindLegs),
			Head:      entity.AttackCatetgory(r.Weakness_attack.Head),
			Body:      entity.AttackCatetgory(r.Weakness_attack.Body),
			Tail:      entity.AttackCatetgory(r.Weakness_attack.Tail),
		}
		weak_e := entity.Weakness_element{
			FrontLegs: entity.Elements(r.Weakness_element.FrontLegs),
			HindLegs:  entity.Elements(r.Weakness_element.HindLegs),
			Head:      entity.Elements(r.Weakness_element.Head),
			Body:      entity.Elements(r.Weakness_element.Body),
			Tail:      entity.Elements(r.Weakness_element.Tail),
		}
		data := entity.Monster{
			Id:               entity.MonsterId{Value: r.MonsterId},
			Name:             entity.MonsterName{Value: r.Name},
			Desc:             entity.MonsterDesc{Value: r.Desc},
			Location:         entity.MonsterLocation{Value: r.Location},
			Category:         entity.MonsterCategory{Value: r.Category},
			Title:            entity.GameTitle{Value: r.Title},
			Weakness_attack:  entity.MonsterWeakness_A{Value: weak_a},
			Weakness_element: entity.MonsterWeakness_E{Value: weak_e},
		}
		result.Values = append(result.Values, data)
	}

	return result, nil
}

func (g MonsterGateway) GetById(id entity.MonsterId) (entity.Monster, error) {
	res := g.monsterDriver.GetById(id.Value)
	if res.ID == 0 {
		return entity.Monster{}, fmt.Errorf("NOT FOUND : id = %s", id)
	}

	weak_a := entity.Weakness_attack{
		FrontLegs: entity.AttackCatetgory(res.Weakness_attack.FrontLegs),
		HindLegs:  entity.AttackCatetgory(res.Weakness_attack.HindLegs),
		Head:      entity.AttackCatetgory(res.Weakness_attack.Head),
		Body:      entity.AttackCatetgory(res.Weakness_attack.Body),
		Tail:      entity.AttackCatetgory(res.Weakness_attack.Tail),
	}
	weak_e := entity.Weakness_element{
		FrontLegs: entity.Elements(res.Weakness_element.FrontLegs),
		HindLegs:  entity.Elements(res.Weakness_element.HindLegs),
		Head:      entity.Elements(res.Weakness_element.Head),
		Body:      entity.Elements(res.Weakness_element.Body),
		Tail:      entity.Elements(res.Weakness_element.Tail),
	}
	result := entity.Monster{
		Id:               entity.MonsterId{Value: res.MonsterId},
		Name:             entity.MonsterName{Value: res.Name},
		Desc:             entity.MonsterDesc{Value: res.Desc},
		Location:         entity.MonsterLocation{Value: res.Location},
		Category:         entity.MonsterCategory{Value: res.Category},
		Title:            entity.GameTitle{Value: res.Title},
		Weakness_attack:  entity.MonsterWeakness_A{Value: weak_a},
		Weakness_element: entity.MonsterWeakness_E{Value: weak_e},
	}

	return result, nil
}

func (g MonsterGateway) Create(monsterJson entity.MonsterJson) error {
	weak_a := repository.Weakness_attack{
		FrontLegs: repository.AttackCatetgory(monsterJson.Weakness_attack.Value.FrontLegs),
		HindLegs:  repository.AttackCatetgory(monsterJson.Weakness_attack.Value.HindLegs),
		Head:      repository.AttackCatetgory(monsterJson.Weakness_attack.Value.Head),
		Body:      repository.AttackCatetgory(monsterJson.Weakness_attack.Value.Body),
		Tail:      repository.AttackCatetgory(monsterJson.Weakness_attack.Value.Tail),
	}
	weak_e := repository.Weakness_element{
		FrontLegs: repository.Elements(monsterJson.Weakness_element.Value.FrontLegs),
		HindLegs:  repository.Elements(monsterJson.Weakness_element.Value.HindLegs),
		Head:      repository.Elements(monsterJson.Weakness_element.Value.Head),
		Body:      repository.Elements(monsterJson.Weakness_element.Value.Body),
		Tail:      repository.Elements(monsterJson.Weakness_element.Value.Tail),
	}

	driverJson := repository.MonsterJson{
		Id:               monsterJson.Id.Value,
		Name:             monsterJson.Name.Value,
		Desc:             monsterJson.Desc.Value,
		Location:         monsterJson.Location.Value,
		Category:         monsterJson.Category.Value,
		Title:            monsterJson.Title.Value,
		Weakness_attack:  weak_a,
		Weakness_element: weak_e,
	}

	err := g.monsterDriver.Create(driverJson)
	return err
}

func (g MonsterGateway) Update(id entity.MonsterId, monsterJson entity.MonsterJson) error {
	weak_a := repository.Weakness_attack{
		FrontLegs: repository.AttackCatetgory(monsterJson.Weakness_attack.Value.FrontLegs),
		HindLegs:  repository.AttackCatetgory(monsterJson.Weakness_attack.Value.HindLegs),
		Head:      repository.AttackCatetgory(monsterJson.Weakness_attack.Value.Head),
		Body:      repository.AttackCatetgory(monsterJson.Weakness_attack.Value.Body),
		Tail:      repository.AttackCatetgory(monsterJson.Weakness_attack.Value.Tail),
	}
	weak_e := repository.Weakness_element{
		FrontLegs: repository.Elements(monsterJson.Weakness_element.Value.FrontLegs),
		HindLegs:  repository.Elements(monsterJson.Weakness_element.Value.HindLegs),
		Head:      repository.Elements(monsterJson.Weakness_element.Value.Head),
		Body:      repository.Elements(monsterJson.Weakness_element.Value.Body),
		Tail:      repository.Elements(monsterJson.Weakness_element.Value.Tail),
	}
	driverJson := repository.MonsterJson{
		Name:             monsterJson.Name.Value,
		Desc:             monsterJson.Desc.Value,
		Location:         monsterJson.Location.Value,
		Category:         monsterJson.Category.Value,
		Title:            monsterJson.Title.Value,
		Weakness_attack:  weak_a,
		Weakness_element: weak_e,
	}

	err := g.monsterDriver.Update(id.Value, driverJson)
	return err
}

func (g MonsterGateway) Delete(id entity.MonsterId) error {
	err := g.monsterDriver.Delete(id.Value)
	return err
}

func ProvideMonsterDriver(monsterDriver repository.MonsterDriver) MonsterGateway {
	return MonsterGateway{monsterDriver: monsterDriver}
}
