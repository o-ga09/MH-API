package handler

import (
	"log"
	"mh-api/api/entity"
	"mh-api/api/service"
	"net/http"
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

func (m *MonsterHandler) GetAll(c *gin.Context) {
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

func (m *MonsterHandler) Create(c *gin.Context) {
	name := c.PostForm("name")
	desc := c.PostForm("desc")
	Location := c.PostForm("location")
	specify := c.PostForm("specify")
	weak_a := c.PostForm("weakness_A")
	weak_e := c.PostForm("weakness_E")
	
	monsterJson := entity.MonsterJson{
		Name: entity.MonsterName{Value: name},
		Desc: entity.MonsterDesc{Value: desc},
		Location: entity.MonsterLocation{Value: Location},
		Specify: entity.MonsterSpecify{Value: specify},
		Weakness_attack: entity.MonsterWeakness_A{Value: weak_a},
		Weakness_element: entity.MonsterWeakness_E{Value: weak_e},
	}

	err := m.monsterService.Create(monsterJson)
	if err != nil {
		c.JSON(500,gin.H{
			"err": "can not create record",
		})
		log.Printf("err: %v",err)
		return
	}
	c.JSON(200,Messageresponse{Message: "success!"})
}

func (m MonsterHandler) Update(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	monsterId := entity.MonsterId{Value: i}

	name := c.PostForm("name")
	desc := c.PostForm("desc")
	Location := c.PostForm("location")
	specify := c.PostForm("specify")
	weak_a := c.PostForm("weakness_A")
	weak_e := c.PostForm("weakness_E")

	monsterJson := entity.MonsterJson{
		Name: entity.MonsterName{Value: name},
		Desc: entity.MonsterDesc{Value: desc},
		Location: entity.MonsterLocation{Value: Location},
		Specify: entity.MonsterSpecify{Value: specify},
		Weakness_attack: entity.MonsterWeakness_A{Value: weak_a},
		Weakness_element: entity.MonsterWeakness_E{Value: weak_e},
	}

	err := m.monsterService.Update(monsterId,monsterJson)
	if err != nil {
		c.JSON(500,gin.H{
			"err": "can not create record",
		})
		log.Printf("err: %v",err)
		return
	}
	c.JSON(200,Messageresponse{Message: "success!"})
}

func (m MonsterHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	monsterId := entity.MonsterId{Value: i}

	err := m.monsterService.Delete(monsterId)
	if err != nil {
		c.JSON(500,gin.H{
			"err": "can not delete record",
		})
		log.Printf("err: %v",err)
		return
	}
	c.JSON(200,Messageresponse{Message: "success!"})
}

func (m *MonsterHandler) CreateJson(c *gin.Context) {
	var data RequestJson
	if err := c.ShouldBindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	for _, record := range data.Req {

		monsterJson := entity.MonsterJson{
			Name: entity.MonsterName{Value: record.Name},
			Desc: entity.MonsterDesc{Value: record.Desc},
			Location: entity.MonsterLocation{Value: record.Location},
			Specify: entity.MonsterSpecify{Value: record.Specify},
			Weakness_attack: entity.MonsterWeakness_A{Value: record.Weakness_attack},
			Weakness_element: entity.MonsterWeakness_E{Value: record.Weakness_element},
		}
		err := m.monsterService.Create(monsterJson)
		if err != nil {
			c.JSON(500,gin.H{
				"err": "can not create record",
			})
			log.Printf("err: %v",err)
			return
		}
	}

	c.JSON(200,Messageresponse{Message: "success!"})
}


func ProvideMonsterHandler(monsterService service.MonsterService) MonsterHandler {
	return MonsterHandler{monsterService: monsterService}
}

type Messageresponse struct {
	Message string `json:"message"`
}

type RequestJson struct {
	Req []Json `json:"req"`
}

type Json struct {
  	Name             string       `json:"name"`
    Desc             string       `json:"desc"`
    Location         string   `json:"location"`
    Specify          string    `json:"specify"`
    Weakness_attack  string `json:"weakness_attack"`
    Weakness_element string `json:"weakness_element"`
}
