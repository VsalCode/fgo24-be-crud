package routers;

import (
	"github.com/gin-gonic/gin"
)

func CombineRouter(r *gin.Engine) {
	usersRouter(r.Group("/users"))
}
