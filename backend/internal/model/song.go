package model

// SongInfo represents the subset of fields we care about
// from the getsong.co /song endpoint.
type SongInfo struct {
	ID           string `json:"id"`
	Tempo        string `json:"tempo"`
	Danceability int    `json:"danceability"`
	Acousticness int    `json:"acousticness"`
	KeyOf        string `json:"key_of"`
	OpenKey      string `json:"open_key"`
}
