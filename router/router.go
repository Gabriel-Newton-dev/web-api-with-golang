package router

import (
	"github.com/Gabriel-Newton-dev/web-api-with-golang/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.ShowBook)
		}
		return router

	}

}
