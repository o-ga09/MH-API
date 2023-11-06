package handler

import (
	"fmt"
	"log/slog"
	"mh-api/api/entity"
	"mh-api/api/middleware"
	"mh-api/api/service"

	"mh-api/api/util"
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "can not get records",
		})
		slog.Log(c, middleware.SeverityError, "err", err)
		return
	}

	monsters := []ResponseJson{}
	for _, r := range res.Values {
		new := ResponseJson{r.Id.Value, r.Name.Value, r.Desc.Value, r.Location.Value, r.Specify.Value, r.Weakness_attack.Value, r.Weakness_element.Value}
		monsters = append(monsters, new)
	}

	response := Monsters{
		Total:    len(res.Values),
		Monsters: monsters,
	}
	c.JSON(http.StatusOK, response)
}

func (m *MonsterHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	monsterId := entity.MonsterId{Value: i}

	res, err := m.monsterService.GetById(monsterId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: fmt.Sprintf("can not get id : %d", i)})
		slog.Log(c, middleware.SeverityError, "error", "error", err.Error())
	}
	monster := ResponseJson{res.Id.Value, res.Name.Value, res.Desc.Value, res.Location.Value, res.Specify.Value, res.Weakness_attack.Value, res.Weakness_element.Value}
	response := Monster{
		Monster: monster,
	}
	c.JSON(http.StatusOK, response)
}

func (m *MonsterHandler) Create(c *gin.Context) {
	var requestBody map[string]interface{}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: err.Error()})
		return
	}

	name := requestBody["name"].(string)
	desc := requestBody["desc"].(string)
	Location := requestBody["location"].(string)
	specify := requestBody["specify"].(string)
	weak_a := requestBody["weakness_A"].(string)
	weak_e := requestBody["weakness_E"].(string)

	weak_map_a := util.Mapping(weak_a)
	weak_map_e := util.Mapping(weak_e)

	monsterJson := entity.MonsterJson{
		Name:             entity.MonsterName{Value: name},
		Desc:             entity.MonsterDesc{Value: desc},
		Location:         entity.MonsterLocation{Value: Location},
		Specify:          entity.MonsterSpecify{Value: specify},
		Weakness_attack:  entity.MonsterWeakness_A{Value: weak_map_a},
		Weakness_element: entity.MonsterWeakness_E{Value: weak_map_e},
	}

	err := m.monsterService.Create(monsterJson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: "can not create"})
		slog.Log(c, middleware.SeverityError, "err", "Error", err)
		return
	}
	slog.Log(c, middleware.SeverityInfo, "success", "info", MessageResponse{Message: "Record Create"})
	c.JSON(http.StatusOK, MessageResponse{Message: "success!"})
}

func (m MonsterHandler) Update(c *gin.Context) {
	var requestBody map[string]interface{}

	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	monsterId := entity.MonsterId{Value: i}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: err.Error()})
		return
	}

	name := requestBody["name"].(string)
	desc := requestBody["desc"].(string)
	Location := requestBody["location"].(string)
	specify := requestBody["specify"].(string)
	weak_a := requestBody["weakness_A"].(string)
	weak_e := requestBody["weakness_E"].(string)

	weak_map_a := util.Mapping(weak_a)
	weak_map_e := util.Mapping(weak_e)

	monsterJson := entity.MonsterJson{
		Name:     entity.MonsterName{Value: name},
		Desc:     entity.MonsterDesc{Value: desc},
		Location: entity.MonsterLocation{Value: Location},
		Specify:  entity.MonsterSpecify{Value: specify},

		Weakness_attack:  entity.MonsterWeakness_A{Value: weak_map_a},
		Weakness_element: entity.MonsterWeakness_E{Value: weak_map_e},
	}

	err := m.monsterService.Update(monsterId, monsterJson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: "can not update"})
		slog.Log(c, middleware.SeverityError, "err", err)
		return
	}
	slog.Log(c, middleware.SeverityInfo, "success", "info", MessageResponse{Message: "Record Update"})
	c.JSON(http.StatusOK, MessageResponse{Message: "success!"})
}

func (m MonsterHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	monsterId := entity.MonsterId{Value: i}

	err := m.monsterService.Delete(monsterId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: "can not delete"})
		slog.Log(c, middleware.SeverityError, "err", err)
		return
	}
	slog.Log(c, middleware.SeverityInfo, "success", "info", MessageResponse{Message: "Record Delete"})
	c.JSON(http.StatusOK, MessageResponse{Message: "success!"})
}

func (m *MonsterHandler) CreateJson(c *gin.Context) {
	var data RequestJson
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not create"})
		return
	}

	for _, record := range data.Req {
		weak_map_a := util.Mapping(record.Weakness_attack)
		weak_map_e := util.Mapping(record.Weakness_element)

		monsterJson := entity.MonsterJson{
			Name:             entity.MonsterName{Value: record.Name},
			Desc:             entity.MonsterDesc{Value: record.Desc},
			Location:         entity.MonsterLocation{Value: record.Location},
			Specify:          entity.MonsterSpecify{Value: record.Specify},
			Weakness_attack:  entity.MonsterWeakness_A{Value: weak_map_a},
			Weakness_element: entity.MonsterWeakness_E{Value: weak_map_e},
		}
		err := m.monsterService.Create(monsterJson)
		if err != nil {
			c.JSON(http.StatusInternalServerError, MessageResponse{Message: err.Error()})
			slog.Log(c, middleware.SeverityError, "err", err)
			return
		}
	}
	slog.Log(c, middleware.SeverityInfo, "success", "info", MessageResponse{Message: "Record Create"})
	c.JSON(http.StatusOK, MessageResponse{Message: "success!"})
}

func ProvideMonsterHandler(monsterService service.MonsterService) MonsterHandler {
	return MonsterHandler{monsterService: monsterService}
}

type Monsters struct {
	Total    int            `json:"total"`
	Monsters []ResponseJson `json:"monsters"`
}

type Monster struct {
	Monster ResponseJson `json:"monster"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type MessageRequest struct {
	Message string `json:"message"`
}

type RequestJson struct {
	Req []Json `json:"req"`
}

type Json struct {
	Name             string `json:"name"`
	Desc             string `json:"desc"`
	Location         string `json:"location"`
	Specify          string `json:"specify"`
	Weakness_attack  string `json:"weakness_attack"`
	Weakness_element string `json:"weakness_element"`
}

type ResponseJson struct {
	Id               int               `json:"id,omitempty"`
	Name             string            `json:"name,omitempty"`
	Desc             string            `json:"desc,omitempty"`
	Location         string            `json:"location,omitempty"`
	Specify          string            `json:"specify,omitempty"`
	Weakness_attack  map[string]string `json:"weakness___attack,omitempty"`
	Weakness_element map[string]string `json:"weakness___element,omitempty"`
}
