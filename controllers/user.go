package controllers

import (
	"go-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetAllUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.User
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetUserById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.User

	db.First(&user, "id = ?", c.Param("id"))

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}

func CreateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := db.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User has been created",
		"data":    user,
	})
}

func UpdateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.User

	db.First(&user, "id = ?", c.Param("id"))

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User Not Found",
		})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db.Save(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "User has been updated",
	})
}

func DeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.User

	db.First(&user, "id = ?", c.Param("id"))

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "User has been deleted",
	})
}
