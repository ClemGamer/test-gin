package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ClemGamer/test-gin/database"
	"github.com/ClemGamer/test-gin/entities"
	"github.com/gin-gonic/gin"
)

type User struct{}

func (c User) All(ctx *gin.Context) {
	users := []entities.User{}
	database.Instance.Find(&users)
	bs, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(http.StatusOK, string(bs))
}
