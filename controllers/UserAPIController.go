package controllers

import (
	"github.com/gin-gonic/gin"
	"mom-note-server/common"
	"mom-note-server/models"
	"net/http"
)

/**
 * author: heinoc
 * created on: 2019-02-12
 */

func RegisterUser(c *gin.Context) {

	phone := c.PostForm("phone")
	password := c.PostForm("password")

	var user = &models.User{
		Phone:    phone,
		Password: password,
	}

	err := user.Inser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.NewResponse(user))

}

func Login(c *gin.Context) {
	phone := c.PostForm("phone")
	userName := c.PostForm("userName")
	password := c.PostForm("password")

	user, err := models.VerifyUserByPwd(phone, userName, password)

	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.NewResponse(user))

}
