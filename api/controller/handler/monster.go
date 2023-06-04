package handler

import (
	"log"
	"mh-api/api/entity"
	"mh-api/api/service"
	"strconv"

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

func (m *MonsterHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	monsterId := entity.MonsterId{Value: i}

	res, err := m.monsterService.GetById(monsterId)
	if err != nil {
		c.JSON(500,gin.H{
			"err": "can not get record",
		})
	}

	c.JSON(200,res)
}


func ProvideMonsterHandler(monsterService service.MonsterService) MonsterHandler {
	return MonsterHandler{monsterService: monsterService}
}