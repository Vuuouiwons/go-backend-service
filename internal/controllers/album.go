package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Vuuouiwons/go-backend-service/internal/db"
	"github.com/Vuuouiwons/go-backend-service/internal/models"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, db.Albums)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range db.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	db.Albums = append(db.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
