package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	name string
}
