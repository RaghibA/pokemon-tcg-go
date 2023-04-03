package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	Name   string
	Types  pq.StringArray `gorm:"type:text[]"`
	Value  float64
	Img    string
	CardId string
	UserId uint
}
