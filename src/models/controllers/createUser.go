package controllers

import (
	"fmt"

	"github.com/Theusant/crud-go-api/src/models/controllers/models/requests"
	"github.com/Theusant/crud-go-api/src/models/controllers/test/config/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var UserResquest requests.UserResquest

	if err := c.ShouldBindJSON(&UserResquest); err != nil {
		restErr := rest_err.NewBadRequestErr(
			fmt.Sprintf("There Some incorrect fields, error=%s\n", err.Error))

		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(UserResquest)
}
