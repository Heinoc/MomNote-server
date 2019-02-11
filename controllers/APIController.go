package controllers

import (
	"github.com/gin-gonic/gin"
	"mom-note-server/common"
	"mom-note-server/models"
	"net/http"
)

/**
 * author: chenshuai09
 */

func AddRecord(c *gin.Context) {
	//userId := c.PostForm("userID")
	weight := c.PostForm("weight")
	waistline := c.PostForm("waistline")

	var record = new(models.Record)
	record.Weight = weight
	record.Waistline = waistline

	c.JSON(http.StatusOK, common.NewResponse(record))
}

func GetAllRecords(c *gin.Context)  {
	c.JSON(http.StatusOK, common.NewResponse("haha"))
}