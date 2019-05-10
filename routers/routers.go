package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"mom-note-server/controllers"
)

/**
 * author: heinoc
 */

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 跨域中间件
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	api := r.Group("api")
	{
		user := api.Group("user")
		{
			user.POST("register", controllers.RegisterUser)
			user.POST("login", controllers.Login)
		}

		record := api.Group("record")
		{
			record.POST("mockPastRecord", controllers.MockPastRecord)
			record.GET("getRecords", controllers.GetRecords)
			record.POST("addRecord", controllers.AddRecord)
		}

	}

	return r
}
