package routers

import (
	"crud-backend/controllers"
	"github.com/gin-gonic/gin"
)

func usersRouter(r *gin.RouterGroup) {
	r.GET("/list-all", controllers.ListAllUsers)
	r.GET("/detail/:id", controllers.DetailUser)
	r.POST("/create", controllers.CreateUser)
	r.DELETE("/delete/:id", controllers.DeleteUser)
}