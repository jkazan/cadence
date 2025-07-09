package service

import (
	"cadence/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type GetSongClient struct {
	APIKey string
}

func NewGetSongClient(apiKey string) *GetSongClient { return &GetSongClient{APIKey: apiKey} }

// Search calls /search and returns the first hit as SongInfo.
func (c *GetSongClient) Search(artist, title string) (model.SongInfo, error) {
	lookup := fmt.Sprintf("song:%s artist:%s", title, artist)
	q := url.Values{
		"type":    {"both"},
		"lookup":  {lookup},
		"api_key": {c.APIKey},
	}

	resp, err := http.Get("https://api.getsong.co/search/?" + q.Encode())
	if err != nil {
		return model.SongInfo{}, err
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	return parseData(raw)
}

// ----- helpers ----------------------------------------------------------

func parseData(raw []byte) (model.SongInfo, error) {
	var env struct {
		Search json.RawMessage `json:"search"`
	}
	if err := json.Unmarshal(raw, &env); err != nil {
		return model.SongInfo{}, err
	}

	// error response?
	var errObj struct {
		Error string `json:"error"`
	}
	if json.Unmarshal(env.Search, &errObj) == nil && errObj.Error != "" {
		return model.SongInfo{}, fmt.Errorf(errObj.Error)
	}

	// success array
	var hits []model.SongInfo
	if err := json.Unmarshal(env.Search, &hits); err != nil {
		return model.SongInfo{}, err
	}
	if len(hits) == 0 {
		return model.SongInfo{}, fmt.Errorf("no results")
	}
	return hits[0], nil
}
