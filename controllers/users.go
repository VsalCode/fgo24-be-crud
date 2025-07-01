package controllers

import (
	"context"
	"crud-backend/dto"
	"crud-backend/models"
	"crud-backend/utils"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ListAllUsers godoc
// @Summary      Get all users
// @Description  Get all users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.Response  "List of all users"
// @Router       /users [get]
func ListAllUsers(ctx *gin.Context) {
	err := utils.RedisClient.Ping(context.Background()).Err()
	noRedis := false
	if err != nil && strings.Contains(err.Error(), "refused") {
		noRedis = true
	}

	key := ctx.Query("search")
	redisKey := "all-users"
	if key != "" {
		redisKey = "all-users:" + key
	}

	var users []models.User

	if !noRedis {
		result := utils.RedisClient.Exists(context.Background(), redisKey)
		if result.Val() != 0 {
			data := utils.RedisClient.Get(context.Background(), redisKey)
			str := data.Val()
			json.Unmarshal([]byte(str), &users)
			ctx.JSON(http.StatusOK, utils.Response{
				Success: true,
				Message: "success to show all list users from redis",
				PageInfo: map[string]any{
					"totalData": len(users),
				},
				Result: users,
			})
			return
		}
	}

	users, err = models.FindUserByName(key)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to Show List All Users!",
		})
		return
	}

	if !noRedis {
		encoded, err := json.Marshal(users)
		if err == nil {
			utils.RedisClient.Set(context.Background(), redisKey, string(encoded), 0)
		}
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

// @Description Create a new user with username, email, and password
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserRequest true "User object"
// @Success 200 {object} utils.Response "User created successfully"
// @Router /users [post]
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

// @Description Get user details by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} utils.Response "User detail retrieved successfully"
// @Router /users/{id} [get]
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

// @Description Delete a user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} utils.Response "User deleted successfully"
// @Router /users/{id} [delete]
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

// @Description Update a user by ID (partial update)
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body dto.UpdateUserRequest true "User object to be updated"
// @Success 200 {object} utils.Response "User updated successfully"
// @Router /users/{id} [patch]
func UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var updateUser dto.UpdateUserRequest

	if err := ctx.ShouldBind(&updateUser); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	err := models.HandleUpdateUser(id, updateUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to update user",
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "User updated successfully",
	})
}
