package models

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	Name   string
	Types  []string
	Value  int32
	Img    string
	UserId uint
}
