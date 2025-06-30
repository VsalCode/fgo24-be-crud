package routers

import (
	"crud-backend/controllers"
	"github.com/gin-gonic/gin"
)

func usersRouter(r *gin.RouterGroup) {
	r.GET("/list-all", controllers.ListAllUsers)
	r.GET("/:id", controllers.DetailUser)
	r.POST("", controllers.CreateUser)
	r.DELETE("/:id", controllers.DeleteUser)
	r.PATCH("/:id", controllers.UpdateUser)
}