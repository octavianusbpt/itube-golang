package models

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Url     string
	VideoId int
}
