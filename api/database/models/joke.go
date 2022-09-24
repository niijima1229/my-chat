package models

import "gorm.io/gorm"

type Joke struct {
	gorm.Model
	title  string
	joke   string
	UserId uint
}
