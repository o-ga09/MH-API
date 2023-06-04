package gateway

import (
	"errors"
	"mh-api/api/driver"
	"mh-api/api/entity"
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

		data := entity.Monster{
			Id: entity.MonsterId{Value: r.Id},
			Name: entity.MonsterName{Value: r.Name},
			Desc: entity.MonsterDesc{Value: r.Desc},
			Location: entity.MonsterLocation{Value: r.Location},
			Specify: entity.MonsterSpecify{Value: r.Specify},
			Weakness_attack: entity.MonsterWeakness_A{Value: r.Weakness_attack},
			Weakness_element: entity.MonsterWeakness_E{Value: r.Weakness_element},
		}
		result.Values =  append(result.Values,data)
	}

	return result,nil
}

func ProvideMonsterDriver(monsterDriver driver.MonsterDriver) MonsterGateway {
	return MonsterGateway{monsterDriver: monsterDriver}
}