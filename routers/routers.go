package routers


import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"mom-note-server/controllers"
)


/**
 * author: chenshuai09
 */

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 跨域中间件
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	api := r.Group("api")
	{
		api.GET("getAllRecords", controllers.GetAllRecords)
		api.POST("addRecord", controllers.AddRecord)
	}

	return r
}
