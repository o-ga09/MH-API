package controller

import (
	"context"
	"mh-api/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type monsterHandler struct {
	Service service.MonsterService
}

func NewMonsterHandler(s service.MonsterService) *monsterHandler {
	return &monsterHandler{
		Service: s,
	}
}

func (m *monsterHandler) GetMonsterById(c *gin.Context) {
	param := c.Params
	id_str, ok := param.Get("id")
	if !ok {
		c.JSON(500,gin.H{
			"err": "id is not understanded",
		})
	}
	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.JSON(500,gin.H{
			"err": "id is not converted",
		})
	}

	res, err := m.Service.FindMonsterById(context.Background(),id)
	if err != nil {
		c.JSON(500,gin.H{
			"err": "can not get record",
		})
	}

	c.JSON(200,gin.H{
		"res": res,
	})
}

func (m *monsterHandler) GetMonsterAll(c *gin.Context) {
	res, err := m.Service.FindAllMonsters(context.Background())
	if err != nil {
		c.JSON(500,gin.H{
			"err": "can not get records",
		})
	}
	c.JSON(200,gin.H{
		"res": res,
	})
}