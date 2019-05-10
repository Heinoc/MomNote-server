package controllers

import (
	"github.com/gin-gonic/gin"
	"mom-note-server/common"
	"mom-note-server/models"
	"net/http"
	"strconv"
	"time"
)

/**
 * author: heinoc
 */

type GetAllRecordsResponse struct {
	TotalPages int64           `json:"totalPages"`
	PageNum    int64           `json:"pageNum"`
	PageSize   int64           `json:"pageSize"`
	List       []models.Record `json:"list"`
}

func MockPastRecord(c *gin.Context) {
	userId := c.PostForm("userID")
	startDate := c.PostForm("startDate")

	startTime, err := time.Parse("2006-01-02", startDate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err.Error()))
		return
	}

	nowTime := time.Now()

	for startTime.Before(nowTime) {

		var record = new(models.Record)
		record.UserId = userId

		record.CreatedAt = startTime
		record.UpdatedAt = startTime

		err = record.Insert()

		if err != nil {
			c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err.Error()))
			return
		}

		startTime = startTime.AddDate(0, 0, 1)

	}

	c.JSON(http.StatusOK, common.NewResponse(""))

}

func AddRecord(c *gin.Context) {
	userId := c.PostForm("userID")
	weight := c.PostForm("weight")
	armline := c.PostForm("armline")
	waistline := c.PostForm("waistline")
	bust := c.PostForm("bust")
	hipline := c.PostForm("hipline")
	thighline := c.PostForm("thighline")

	var record = new(models.Record)
	record.UserId = userId
	record.Weight = weight
	record.Armline = armline
	record.Waistline = waistline
	record.Bust = bust
	record.Hipline = hipline
	record.Thighline = thighline

	err := record.Insert()
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.NewResponse(""))
}

func GetRecords(c *gin.Context) {

	userId := c.DefaultQuery("userID", "")
	pageNumStr := c.DefaultQuery("pageNum", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	pageNum, _ := strconv.ParseInt(pageNumStr, 10, 64)
	pageSize, _ := strconv.ParseInt(pageSizeStr, 10, 64)

	totalPages, list, err := models.FindAllRecords(userId, pageNum, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err.Error()))
		return
	}

	var response = &GetAllRecordsResponse{
		TotalPages: totalPages,
		PageNum:    pageNum,
		PageSize:   pageSize,
		List:       list,
	}

	c.JSON(http.StatusOK, common.NewResponse(response))

}
