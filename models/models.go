package models

import "gorm.io/gorm"

type Professor struct {
	gorm.Model

	Name       string `json:"name"`
	University string `json:"university"`
}
