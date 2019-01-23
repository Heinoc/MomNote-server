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
	userId := c.PostForm("userID")
	weight := c.PostForm("weight")
	waistline := c.PostForm("waistline")
	c.JSON(http.StatusOK, common.NewResponse(userId+weight+waistline))
}

func GetAllRecords(c *gin.Context)  {
	c.JSON(http.StatusOK, common.NewResponse("haha"))
}