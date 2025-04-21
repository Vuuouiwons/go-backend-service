package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Vuuouiwons/go-backend-service/internal/controllers"
)

func addAblumsRoute(rg *gin.RouterGroup) {
	albums := rg.Group("/albums")

	albums.GET("/", controllers.GetAlbums)
	albums.GET("/:id", controllers.GetAlbumByID)
	albums.POST("/", controllers.PostAlbums)
}
