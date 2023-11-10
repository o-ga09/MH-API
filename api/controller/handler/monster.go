package handler

import (
	"fmt"
	"log/slog"
	"mh-api/api/entity"
	"mh-api/api/middleware"
	"mh-api/api/service"

	"net/http"

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
		weak_a := Weakness_attack{
			FrontLegs: AttackCatetgory(r.Weakness_attack.Value.FrontLegs),
			HindLegs:  AttackCatetgory(r.Weakness_attack.Value.HindLegs),
			Head:      AttackCatetgory(r.Weakness_attack.Value.Head),
			Body:      AttackCatetgory(r.Weakness_attack.Value.Body),
			Tail:      AttackCatetgory(r.Weakness_attack.Value.Tail),
		}
		weak_e := Weakness_element{
			FrontLegs: Elements(r.Weakness_element.Value.FrontLegs),
			HindLegs:  Elements(r.Weakness_element.Value.HindLegs),
			Head:      Elements(r.Weakness_element.Value.Head),
			Body:      Elements(r.Weakness_element.Value.Body),
			Tail:      Elements(r.Weakness_element.Value.Tail),
		}
		new := ResponseJson{r.Id.Value, r.Name.Value, r.Desc.Value, r.Location.Value, r.Category.Value, r.Title.Value, weak_a, weak_e}
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
	monsterId := entity.MonsterId{Value: id}

	res, err := m.monsterService.GetById(monsterId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: fmt.Sprintf("can not get id : %s", id)})
		slog.Log(c, middleware.SeverityError, "error", "error", err.Error())
	}
	weak_a := Weakness_attack{
		FrontLegs: AttackCatetgory(res.Weakness_attack.Value.FrontLegs),
		HindLegs:  AttackCatetgory(res.Weakness_attack.Value.HindLegs),
		Head:      AttackCatetgory(res.Weakness_attack.Value.Head),
		Body:      AttackCatetgory(res.Weakness_attack.Value.Body),
		Tail:      AttackCatetgory(res.Weakness_attack.Value.Tail),
	}
	weak_e := Weakness_element{
		FrontLegs: Elements(res.Weakness_element.Value.FrontLegs),
		HindLegs:  Elements(res.Weakness_element.Value.HindLegs),
		Head:      Elements(res.Weakness_element.Value.Head),
		Body:      Elements(res.Weakness_element.Value.Body),
		Tail:      Elements(res.Weakness_element.Value.Tail),
	}
	monster := ResponseJson{res.Id.Value, res.Name.Value, res.Desc.Value, res.Location.Value, res.Category.Value, res.Title.Value, weak_a, weak_e}
	response := Monster{
		Monster: monster,
	}

	c.JSON(http.StatusOK, response)
}

func (m *MonsterHandler) Create(c *gin.Context) {
	var requestBody Json

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: err.Error()})
		return
	}
	weak_a := entity.Weakness_attack{
		FrontLegs: entity.AttackCatetgory(requestBody.Weakness_attack.FrontLegs),
		HindLegs:  entity.AttackCatetgory(requestBody.Weakness_attack.HindLegs),
		Head:      entity.AttackCatetgory(requestBody.Weakness_attack.Head),
		Body:      entity.AttackCatetgory(requestBody.Weakness_attack.Body),
		Tail:      entity.AttackCatetgory(requestBody.Weakness_attack.Tail),
	}
	weak_e := entity.Weakness_element{
		FrontLegs: entity.Elements(requestBody.Weakness_element.FrontLegs),
		HindLegs:  entity.Elements(requestBody.Weakness_element.HindLegs),
		Head:      entity.Elements(requestBody.Weakness_element.Head),
		Body:      entity.Elements(requestBody.Weakness_element.Body),
		Tail:      entity.Elements(requestBody.Weakness_element.Tail),
	}
	monsterJson := entity.MonsterJson{
		Id:               entity.MonsterId{Value: requestBody.MonsterId},
		Name:             entity.MonsterName{Value: requestBody.Name},
		Desc:             entity.MonsterDesc{Value: requestBody.Desc},
		Location:         entity.MonsterLocation{Value: requestBody.Location},
		Category:         entity.MonsterCategory{Value: requestBody.Category},
		Title:            entity.GameTitle{Value: requestBody.Title},
		Weakness_attack:  entity.MonsterWeakness_A{Value: weak_a},
		Weakness_element: entity.MonsterWeakness_E{Value: weak_e},
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
	var requestBody Json

	id := c.Param("id")
	monsterId := entity.MonsterId{Value: id}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: err.Error()})
		return
	}

	weak_a := entity.Weakness_attack{
		FrontLegs: entity.AttackCatetgory(requestBody.Weakness_attack.FrontLegs),
		HindLegs:  entity.AttackCatetgory(requestBody.Weakness_attack.HindLegs),
		Head:      entity.AttackCatetgory(requestBody.Weakness_attack.Head),
		Body:      entity.AttackCatetgory(requestBody.Weakness_attack.Body),
		Tail:      entity.AttackCatetgory(requestBody.Weakness_attack.Tail),
	}
	weak_e := entity.Weakness_element{
		FrontLegs: entity.Elements(requestBody.Weakness_element.FrontLegs),
		HindLegs:  entity.Elements(requestBody.Weakness_element.HindLegs),
		Head:      entity.Elements(requestBody.Weakness_element.Head),
		Body:      entity.Elements(requestBody.Weakness_element.Body),
		Tail:      entity.Elements(requestBody.Weakness_element.Tail),
	}
	monsterJson := entity.MonsterJson{
		Name:             entity.MonsterName{Value: requestBody.Name},
		Desc:             entity.MonsterDesc{Value: requestBody.Desc},
		Location:         entity.MonsterLocation{Value: requestBody.Location},
		Category:         entity.MonsterCategory{Value: requestBody.Category},
		Title:            entity.GameTitle{Value: requestBody.Title},
		Weakness_attack:  entity.MonsterWeakness_A{Value: weak_a},
		Weakness_element: entity.MonsterWeakness_E{Value: weak_e},
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
	monsterId := entity.MonsterId{Value: id}

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
		weak_a := entity.Weakness_attack{
			FrontLegs: entity.AttackCatetgory(record.Weakness_attack.FrontLegs),
			HindLegs:  entity.AttackCatetgory(record.Weakness_attack.HindLegs),
			Head:      entity.AttackCatetgory(record.Weakness_attack.Head),
			Body:      entity.AttackCatetgory(record.Weakness_attack.Body),
			Tail:      entity.AttackCatetgory(record.Weakness_attack.Tail),
		}
		weak_e := entity.Weakness_element{
			FrontLegs: entity.Elements(record.Weakness_element.FrontLegs),
			HindLegs:  entity.Elements(record.Weakness_element.HindLegs),
			Head:      entity.Elements(record.Weakness_element.Head),
			Body:      entity.Elements(record.Weakness_element.Body),
			Tail:      entity.Elements(record.Weakness_element.Tail),
		}
		monsterJson := entity.MonsterJson{
			Name:             entity.MonsterName{Value: record.Name},
			Desc:             entity.MonsterDesc{Value: record.Desc},
			Location:         entity.MonsterLocation{Value: record.Location},
			Category:         entity.MonsterCategory{Value: record.Category},
			Title:            entity.GameTitle{Value: record.Title},
			Weakness_attack:  entity.MonsterWeakness_A{Value: weak_a},
			Weakness_element: entity.MonsterWeakness_E{Value: weak_e},
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
	MonsterId        string           `json:"monster_id,omitempty"`
	Name             string           `json:"name,omitempty"`
	Desc             string           `json:"desc,omitempty"`
	Location         string           `json:"location,omitempty"`
	Category         string           `json:"category,omitempty"`
	Title            string           `json:"title,omitempty"`
	Weakness_attack  Weakness_attack  `json:"weakness_attack,omitempty"`
	Weakness_element Weakness_element `json:"weakness_element,omitempty"`
}

type ResponseJson struct {
	Id               string           `json:"monster_id,omitempty"`
	Name             string           `json:"name,omitempty"`
	Desc             string           `json:"desc,omitempty"`
	Location         string           `json:"location,omitempty"`
	Category         string           `json:"category,omitempty"`
	Title            string           `json:"title,omitempty"`
	Weakness_attack  Weakness_attack  `json:"weakness_attack,omitempty"`
	Weakness_element Weakness_element `json:"weakness_element,omitempty"`
}

type Weakness_attack struct {
	FrontLegs AttackCatetgory `json:"front_legs,omitempty"`
	Tail      AttackCatetgory `json:"tail,omitempty"`
	HindLegs  AttackCatetgory `json:"hind_legs,omitempty"`
	Body      AttackCatetgory `json:"body,omitempty"`
	Head      AttackCatetgory `json:"head,omitempty"`
}

type Weakness_element struct {
	FrontLegs Elements `json:"front_legs,omitempty"`
	Tail      Elements `json:"tail,omitempty"`
	HindLegs  Elements `json:"hind_legs,omitempty"`
	Body      Elements `json:"body,omitempty"`
	Head      Elements `json:"head,omitempty"`
}

type AttackCatetgory struct {
	Slashing string `json:"slashing,omitempty"`
	Blow     string `json:"blow,omitempty"`
	Bullet   string `json:"bullet,omitempty"`
}

type Elements struct {
	Fire      string `json:"fire,omitempty"`
	Water     string `json:"water,omitempty"`
	Lightning string `json:"lightning,omitempty"`
	Ice       string `json:"ice,omitempty"`
	Dragon    string `json:"dragon,omitempty"`
}
