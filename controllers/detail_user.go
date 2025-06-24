package controllers

import (
	"crud-backend/models"
	"crud-backend/utils"

	"github.com/gin-gonic/gin"
)

func DetailUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(400, utils.Response{
			Success: false,
			Message: "Search query cannot be empty",
		})
		return
	}

	users, err := models.GetUserDetail(id)
	if err != nil {
		ctx.JSON(500, utils.Response{
			Success: false,
			Message: "Failed to get user details",
		})
		return
	}
	ctx.JSON(200, utils.Response{
		Success: true,
		Message: "Get user detail successfully",
		Result:  users,
		PageInfo: map[string]any{
			"totalData": len(users),
		},
	})
}
