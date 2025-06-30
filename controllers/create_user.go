package controllers

import (
	"crud-backend/dto"
	"crud-backend/models"
	"crud-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Description Create a new user with username, email, and password
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserRequest true "User object"
// @Success 200 {object} utils.Response "User created successfully"
// @Router /users/create [post]
func CreateUser(ctx *gin.Context) {
	user := dto.CreateUserRequest{}
	ctx.ShouldBind(&user)

	if user.Username == "" || user.Email == "" || user.Password == "" {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Username, email, and password are required",
		})
		return
	}

	err := models.HandleCreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to create user",
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "User created successfully",
	})

}