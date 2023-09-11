package controllers

import (
	"github.com/Theusant/crud-go-api/src/models/controllers/test/config/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestErr("Do you get routes bad request.")
	c.JSON(err.Code, err)
}
