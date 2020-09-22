package controllers

import (
	"apiservice/authorization"
	"apiservice/connections"
	"apiservice/models"
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	db := connections.DBConn()

	var json models.UserParam

	if err := c.ShouldBindJSON(&json); err == nil {
		s := sha1.New()
		s.Write([]byte(time.Now().Format("2006-01-02 15:04:05") + json.Email))
		salt := hex.EncodeToString(s.Sum(nil))

		h := sha1.New()
		h.Write([]byte(json.Password + salt))

		password := hex.EncodeToString(h.Sum(nil))

		user := models.User{UserName: json.UserName, Email: json.Email, EncryptedPassword: password, Salt: salt, CreatedDate: time.Now()}

		result := db.Create(&user)

		if result.Error != nil {
			c.JSON(500, gin.H{
				"messages": result.Error,
			})
			return
		}

		c.JSON(200, gin.H{
			"messages": "inserted",
		})
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}

func UpdateUser(c *gin.Context) {

	if authorization.CheckAuth(c.Request) != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	db := connections.DBConn()

	var json models.UpdateUserParam

	if err := c.ShouldBindJSON(&json); err == nil {
		user := models.User{UserName: json.UserName, Email: json.Email}

		result := db.Model(&models.User{}).Where("id = ?", c.Param("id")).Updates(user)

		if result.Error != nil {
			c.JSON(500, gin.H{
				"messages": result.Error,
			})
			return
		}

		c.JSON(200, gin.H{
			"messages": "updated",
		})
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}

func GetUser(c *gin.Context) {

	if authorization.CheckAuth(c.Request) != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	db := connections.DBConn()
	var user models.GetUser

	result := db.First(&user, c.Param("id"))
	if result.Error != nil {
		c.JSON(500, gin.H{
			"messages": "User not found",
		})
		return
	}

	data := map[string]interface{}{
		"User": user,
	}
	c.JSON(200, data)
}

func DeleteUser(c *gin.Context) {

	if authorization.CheckAuth(c.Request) != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	db := connections.DBConn()
	var user models.User

	result := db.Delete(&user, c.Param("id"))

	if result.Error != nil {
		c.JSON(500, gin.H{
			"messages": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"messages": "deleted",
	})
}
