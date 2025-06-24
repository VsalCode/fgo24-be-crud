package controllers

import (
	"crud-backend/models"
	"crud-backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(ctx *gin.Context) {
	user := models.User{}
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