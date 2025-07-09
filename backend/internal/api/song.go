package api

import (
	"cadence/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterTempoRoutes attaches /api/songInfo to the router
func RegisterTempoRoutes(r *gin.Engine, gs *service.GetSongClient) {
	r.GET("/api/songInfo", func(c *gin.Context) {
		artist := c.Query("artist")
		title := c.Query("title")

		info, err := gs.Search(artist, title)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response := gin.H{
			"artist":       artist,
			"title":        title,
			"tempo":        info.Tempo,
			"danceability": info.Danceability,
			"acousticness": info.Acousticness,
			"key_of":       info.KeyOf,
			"open_key":     info.OpenKey,
		}
		c.JSON(http.StatusOK, response)
	})
}
