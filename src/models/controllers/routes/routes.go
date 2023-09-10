package routes

import (
	"github.com/Theusant/crud-go-api/src/models/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/GetUserById/:userId", controllers.FindUserById)
	r.GET("/GetUserbyEmail/:userEmail", controllers.FindUserByEmail)
	r.POST("CreateUser", controllers.CreateUser)
	r.PUT("/updateUser/:userId", controllers.UpdateUser)
	r.DELETE("/deleteUser/:userId", controllers.DeleteUser)

}
