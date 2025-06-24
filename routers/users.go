package routers

import (
	"crud-backend/controllers"
	"github.com/gin-gonic/gin"
)

func usersRouter(r *gin.RouterGroup) {
	r.GET("/list-all", controllers.ListAllUsers)
	// r.POST("/detail", controllers.DetailUsers )
}