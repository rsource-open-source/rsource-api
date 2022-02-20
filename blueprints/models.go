package blueprints

import "gorm.io/gorm"

// type Professor struct {
// 	gorm.Model

// 	Name       string `json:"name"`
// 	University string `json:"university"`
// }

type Redirect struct {
	gorm.Model

	UserID    int    `json:"user_id"`
	PlaceID   int    `json:"place_id"`
	ServerID  string `json:"server_id"`  // UUID
	RequestID string `json:"request_id"` // UUID
}
