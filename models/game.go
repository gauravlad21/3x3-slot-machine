package models

type PlayRequest struct {
	PlayerID string `json:"player_id"`
}

type PlayResponse struct {
	Result  string `json:"result"`
	Credits int    `json:"credits"`
}
