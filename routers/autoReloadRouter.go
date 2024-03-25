package routers

import (
	"assignment3/controllers"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
		AllowHeaders:     []string{"Content-Type", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.GET("/status", controllers.GetStatus)
	router.GET("/", controllers.RenderStatus)
	return router
}
