package controllers

import (
	"apiservice/authorization"
	"apiservice/connections"
	"apiservice/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db := connections.DBConn()
	var u models.UserLogin
	var user models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	result := db.Where("email = ?", u.Email).
		First(&user)

	if result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, result.Error)
	}
	//compare the user from the request, with the one we defined:
	// fmt.Println(user.Password, u.Password, CheckPasswordHash(user.Password, u.Password))
	// !CheckPasswordHash(user.Password, u.Password)
	if user.Email != u.Email || !authorization.CheckPasswordHash(u.Password, user.EncryptedPassword, user.Salt) {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := authorization.CreateToken(uint64(user.ID))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	saveErr := authorization.CreateAuth(uint64(user.ID), token)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
		return
	}

	tokens := map[string]string{
		"access_token": token.AccessToken,
	}

	c.JSON(http.StatusOK, tokens)
}

func Logout(c *gin.Context) {
	au, err := authorization.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	delErr := authorization.DeleteAuth(au.AccessUuid)
	if delErr != nil { //if any goes wrong
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}
