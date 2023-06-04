package handler

import (
	"log"
	"mh-api/api/service"

	"github.com/gin-gonic/gin"
)

type MonsterHandler struct {
	monsterService service.MonsterService
}

func NewMonsterHandler(s service.MonsterService) *MonsterHandler {
	return &MonsterHandler{
		monsterService: s,
	}
}

func (m *MonsterHandler) GetrAll(c *gin.Context) {
	res, err := m.monsterService.GetAll()
	if err != nil {
		c.JSON(500,gin.H{
			"err": "can not get records",
		})
		log.Printf("err: %v",err)
		return
	}
	c.JSON(200,res)
}

// func (m *MonsterHandler) GetMonsterById(c *gin.Context) {
// 	param := c.Params
// 	id_str, ok := param.Get("id")
// 	if !ok {
// 		c.JSON(500,gin.H{
// 			"err": "id is not understanded",
// 		})
// 	}
// 	id, err := strconv.Atoi(id_str)
// 	if err != nil {
// 		c.JSON(500,gin.H{
// 			"err": "id is not converted",
// 		})
// 	}

// 	res, err := m.Service.FindMonsterById(context.Background(),id)
// 	if err != nil {
// 		c.JSON(500,gin.H{
// 			"err": "can not get record",
// 		})
// 	}

// 	c.JSON(200,res)
// }


func ProvideMonsterHandler(monsterService service.MonsterService) MonsterHandler {
	return MonsterHandler{monsterService: monsterService}
}