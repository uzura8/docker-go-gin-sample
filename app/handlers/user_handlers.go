package handlers

import (
	"net/http"

	"gin-sample/services"

	"github.com/gin-gonic/gin"
)

// GetUsers - Returns all users
func GetUsers(c *gin.Context) {
	// ユーザー一覧を取得するロジックを実装
	//c.JSON(http.StatusOK, gin.H{"status": "GetUsers called"})
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUser - Returns a user by ID
func GetUser(c *gin.Context) {
	id := c.Param("id")
	// idに基づいてユーザーを取得するロジックを実装
	c.JSON(http.StatusOK, gin.H{"status": "GetUser called", "id": id})
}

// CreateUser - Creates a new user
func CreateUser(c *gin.Context) {
	// 新規ユーザーを作成するロジックを実装
	c.JSON(http.StatusOK, gin.H{"status": "CreateUser called"})
}

// UpdateUser - Updates a user by ID
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	// idに基づいてユーザーを更新するロジックを実装
	c.JSON(http.StatusOK, gin.H{"status": "UpdateUser called", "id": id})
}

// DeleteUser - Deletes a user by ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	// idに基づいてユーザーを削除するロジックを実装
	c.JSON(http.StatusOK, gin.H{"status": "DeleteUser called", "id": id})
}
