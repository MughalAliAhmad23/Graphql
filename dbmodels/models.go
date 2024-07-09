package dbmodels

import (
	"gorm.io/gorm"
)

type Joke struct {
	gorm.Model
	Text *string `gorm:"not null"`
}

type JokeInput struct {
	Text string `gorm:"not null"`
}
