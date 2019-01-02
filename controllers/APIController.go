package controllers

import (
	"github.com/gin-gonic/gin"
	"mom-note-server/common"
	"net/http"
)

/**
 * author: chenshuai09
 */

func AddRecord(c *gin.Context) {
	userId := c.PostForm("userId")
	wight := c.PostForm("wight")
	waistline := c.PostForm("waistline")
	c.JSON(http.StatusOK, common.NewResponse(userId+wight+waistline))
}
