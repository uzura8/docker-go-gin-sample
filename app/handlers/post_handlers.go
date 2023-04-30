package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPosts - Returns all posts
func GetPosts(c *gin.Context) {
	// 商品一覧を取得するロジックを実装
	c.JSON(http.StatusOK, gin.H{"status": "GetPosts called"})
}

// GetPost - Returns a post by ID
func GetPost(c *gin.Context) {
	id := c.Param("id")
	// idに基づいて商品を取得するロジックを実装
	c.JSON(http.StatusOK, gin.H{"status": "GetPost called", "id": id})
}

// CreatePost - Creates a new post
func CreatePost(c *gin.Context) {
	// 新規商品を作成するロジックを実装
	c.JSON(http.StatusOK, gin.H{"status": "CreatePost called"})
}

// UpdatePost - Updates a post by ID
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	// idに基づいて商品を更新するロジックを実装
	c.JSON(http.StatusOK, gin.H{"status": "UpdatePost called", "id": id})
}

// DeletePost - Deletes a post by ID
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	// idに基づいて商品を削除するロジックを実装
	c.JSON(http.StatusOK, gin.H{"status": "DeletePost called", "id": id})
}
