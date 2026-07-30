package entity

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	IsDone bool
	Title  string
}
