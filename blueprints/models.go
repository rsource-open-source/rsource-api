package blueprints

import "gorm.io/gorm"

type Redirect struct {
	gorm.Model

	UserID    int    `json:"user_id"`
	PlaceID   int    `json:"place_id"`
	ServerID  string `json:"server_id"`  // UUID
	RequestID string `json:"request_id"` // SHA256
}

type Error struct {
	Message string `json:"message"`
}
