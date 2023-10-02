package schemas

import (
	"gorm.io/gorm"
)

type Anime struct {
	gorm.Model
	AnimeId       string
	Name          string
	CommunityRate int8
	Year          int16
	Station       Season
}

type Season int

const (
	Spring Season = iota
	Summer
	Autumn
	Winter
)
