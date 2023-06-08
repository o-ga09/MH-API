package gateway

import (
	"errors"
	"fmt"
	"mh-api/api/driver"
	"mh-api/api/entity"
	"mh-api/api/util"
)

type MonsterGateway struct {
	monsterDriver driver.MonsterDriver
}

func (g MonsterGateway) GetAll() (entity.Monsters,error) {
	res := g.monsterDriver.GetAll()

	var result entity.Monsters

	if len(res) == 0 {
		return entity.Monsters{}, errors.New("0件のレコードを取得しました！")
	}
	 
	for _,r := range res {

		weakness_a := util.Mapping(r.Weakness_attack)
		weakness_e := util.Mapping(r.Weakness_element)
		data := entity.Monster{
			Id: entity.MonsterId{Value: r.Id},
			Name: entity.MonsterName{Value: r.Name},
			Desc: entity.MonsterDesc{Value: r.Desc},
			Location: entity.MonsterLocation{Value: r.Location},
			Specify: entity.MonsterSpecify{Value: r.Specify},
			Weakness_attack: entity.MonsterWeakness_A{Value: weakness_a},
			Weakness_element: entity.MonsterWeakness_E{Value: weakness_e},
		}
		result.Values =  append(result.Values,data)
	}

	return result,nil
}

func (g MonsterGateway) GetById(id entity.MonsterId) (entity.Monster, error) {
	monsterId := id.Value
	res := g.monsterDriver.GetById(monsterId)
	if res.Id == 0 {
		return entity.Monster{}, fmt.Errorf("{%d} のレコードはありませんでした！",id)
	}

	weakness_a := util.Mapping(res.Weakness_attack)
	weakness_e := util.Mapping(res.Weakness_element)
	result := entity.Monster{
		Id: entity.MonsterId{Value: res.Id},
		Name: entity.MonsterName{Value: res.Name},
		Desc: entity.MonsterDesc{Value: res.Desc},
		Location: entity.MonsterLocation{Value: res.Location},
		Specify: entity.MonsterSpecify{Value: res.Specify},
		Weakness_attack: entity.MonsterWeakness_A{Value: weakness_a},
		Weakness_element: entity.MonsterWeakness_E{Value: weakness_e},
	}

	return result, nil
} 

func (g MonsterGateway) Create(monsterJson entity.MonsterJson) error {
	weakness_a := util.Strtomap(monsterJson.Weakness_attack.Value)
	weakness_e := util.Strtomap(monsterJson.Weakness_element.Value)
	driverJson := driver.MonsterJson{
		Name: monsterJson.Name.Value,
		Desc: monsterJson.Desc.Value,
		Location: monsterJson.Location.Value,
		Specify: monsterJson.Specify.Value,
		Weakness_attack: weakness_a,
		Weakness_element: weakness_e,
	}

	err := g.monsterDriver.Create(driverJson)
	return err
}

func (g MonsterGateway) Update(id entity.MonsterId,monsterJson entity.MonsterJson) error {
	monsterId := id.Value
	weakness_a := util.Strtomap(monsterJson.Weakness_attack.Value)
	weakness_e := util.Strtomap(monsterJson.Weakness_element.Value)
	driverJson := driver.MonsterJson{
		Name: monsterJson.Name.Value,
		Desc: monsterJson.Desc.Value,
		Location: monsterJson.Location.Value,
		Specify: monsterJson.Specify.Value,
		Weakness_attack: weakness_a,
		Weakness_element: weakness_e,
	}

	err := g.monsterDriver.Update(monsterId,driverJson)
	return err
}

func (g MonsterGateway) Delete(id entity.MonsterId) error {
	monsterId := id.Value

	err := g.monsterDriver.Delete(monsterId)
	return err
}

func ProvideMonsterDriver(monsterDriver driver.MonsterDriver) MonsterGateway {
	return MonsterGateway{monsterDriver: monsterDriver}
}