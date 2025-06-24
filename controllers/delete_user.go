package controllers

import (
	"crud-backend/models"
	"crud-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "User ID is required",
		})
		return
	}

	err := models.HandleDeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to delete user",
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "User deleted successfully",
	})
}
