package handler

import (
	"log/slog"

	"mh-api/app/internal/middleware"
	"mh-api/app/internal/service/monsters"

	"net/http"

	"github.com/gin-gonic/gin"
)

type MonsterHandler struct {
	monsterService monsters.MonsterService
}

func NewMonsterHandler(s monsters.MonsterService) *MonsterHandler {
	return &MonsterHandler{
		monsterService: s,
	}
}

func (m *MonsterHandler) GetAll(c *gin.Context) {
	id, ook := c.Params.Get("id")
	if ook {
		id = ""
	}
	res, err := m.monsterService.GetMonster(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "can not get records",
		})
		slog.Log(c, middleware.SeverityError, "err", err)
		return
	}

	monsters := []ResponseJson{}
	for _, r := range res {
		monsters = append(monsters, ResponseJson{
			Id:               r.ID,
			Name:             r.Name,
			Desc:             r.Description,
			Location:         "dummy location",
			Category:         "dummy category",
			Title:            "dummy title",
			Weakness_attack:  "dummy weakness attack",
			Weakness_element: "dummy weakness element",
		})
	}
	response := Monsters{
		Total:    len(res),
		Monsters: monsters,
	}
	c.JSON(http.StatusOK, response)
}

func (m *MonsterHandler) GetById(c *gin.Context) {
	id, ook := c.Params.Get("id")
	if ook {
		id = ""
	}
	res, err := m.monsterService.GetMonster(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "can not get records",
		})
		slog.Log(c, middleware.SeverityError, "err", err)
		return
	}

	response := Monster{
		Monster: ResponseJson{
			Id:               res[0].ID,
			Name:             res[0].Name,
			Desc:             res[0].Description,
			Location:         "dummy location",
			Category:         "dummy category",
			Title:            "dummy title",
			Weakness_attack:  "dummy weakness attack",
			Weakness_element: "dummy weakness element",
		},
	}
	c.JSON(http.StatusOK, response)
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
	MonsterId        string `json:"monster_id,omitempty"`
	Name             string `json:"name,omitempty"`
	Desc             string `json:"desc,omitempty"`
	Location         string `json:"location,omitempty"`
	Category         string `json:"category,omitempty"`
	Title            string `json:"title,omitempty"`
	Weakness_attack  string `json:"weakness_attack,omitempty"`
	Weakness_element string `json:"weakness_element,omitempty"`
}

type ResponseJson struct {
	Id               string `json:"monster_id,omitempty"`
	Name             string `json:"name,omitempty"`
	Desc             string `json:"desc,omitempty"`
	Location         string `json:"location,omitempty"`
	Category         string `json:"category,omitempty"`
	Title            string `json:"title,omitempty"`
	Weakness_attack  string `json:"weakness_attack,omitempty"`
	Weakness_element string `json:"weakness_element,omitempty"`
}
