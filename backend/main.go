package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const apiKey = ""

func main() {
	router := gin.Default()

	// Configure CORS to allow requests from frontend
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	// Define API route
	router.GET("/api/tempo", func(c *gin.Context) {
		artist := c.Query("artist")
		title := c.Query("title")
		info, _ := getSongInfo(artist, title)

		c.JSON(http.StatusOK, gin.H{
			"artist":       artist,
			"title":        title,
			"tempo":        info.Tempo,
			"danceability": info.Danceability,
			"acousticness": info.Acousticness,
			"key_of":       info.KeyOf,
			"open_key":     info.OpenKey,
		})
	})

	router.Run(":8080")
}

type songInfo struct {
	ID           string `json:"id"`
	Tempo        string `json:"tempo"`
	Danceability int    `json:"danceability"`
	Acousticness int    `json:"acousticness"`
	KeyOf        string `json:"key_of"`
	OpenKey      string `json:"open_key"`
}

type getSongResponse struct {
	Search []songInfo `json:"search"`
}

func firstResult(raw []byte) (songInfo, error) {
	var envelope struct {
		Search json.RawMessage `json:"search"`
	}
	if err := json.Unmarshal(raw, &envelope); err != nil {
		return songInfo{}, err
	}

	// Check if error object was returned, i.e. {"error":"…"}
	var errObj struct {
		Error string `json:"error"`
	}
	if json.Unmarshal(envelope.Search, &errObj) == nil && errObj.Error != "" {
		return songInfo{}, fmt.Errorf(errObj.Error)
	}

	// Decode as array
	var hits []songInfo
	if err := json.Unmarshal(envelope.Search, &hits); err != nil {
		return songInfo{}, err
	}
	if len(hits) == 0 {
		return songInfo{}, fmt.Errorf("no results")
	}

	// Return first item
	fmt.Println(hits[0].Tempo, hits[0].Acousticness, hits[0].Danceability, hits[0].KeyOf, hits[0].OpenKey)
	return hits[0], nil
}

// -------------------------------------------------------------------
// getSongInfo: search title (and optional artist) → tempo
// -------------------------------------------------------------------
func getSongInfo(artist, title string) (songInfo, error) {
	// Construct lookup string
	lookup := fmt.Sprintf("song:%s artist:%s", title, artist)

	// Build full URL with query parameters
	query := url.Values{}
	query.Set("api_key", apiKey)
	query.Set("type", "both")
	query.Set("lookup", lookup)

	// Compose the final request URL
	fullURL := "https://api.getsong.co/search/?" + query.Encode()
	resp, err := http.Get(fullURL)
	if err != nil {
		return songInfo{}, err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	fmt.Println()
	fmt.Println(string(data))
	fmt.Println()
	first, err := firstResult(data)
	if err != nil {
		// if strings.ToLower(err.Error()) == "no result" {
		// 	return "No result", nil
		// }
		return songInfo{}, err
	}
	return first, nil
}
