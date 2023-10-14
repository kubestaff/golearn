package setting

import "gorm.io/gorm"

type Settings struct {
	gorm.Model
	AboutTitle string
	AboutText string
	VideosCountOnMainPage uint
}