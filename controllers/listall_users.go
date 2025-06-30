package controllers

import (
	"crud-backend/models"
	"crud-backend/utils"
	"net/http"
	"github.com/gin-gonic/gin"
);

// ListAllUsers godoc
// @Summary      Get all users
// @Description  Get all users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.Response  "List of all users"
// @Router       /users/list-all [get]
func ListAllUsers(ctx *gin.Context) {
	users, err := models.FindAllUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to Show List All Users!",
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "success to show all list users",
		PageInfo: map[string]any{
			"totalData": len(users),
		},
		Result: users,
	})

}