package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	Question string `json:"question" gorm:"text;not null; default:null"`
	Answer string `json:"answer" gorm:"text;not null; default:null"`
}

type Caps struct {
	gorm.Model
	Message string `json:"message" gorm:"text;not null; default:null"`
	Answer string `json:"answer" gorm:"text;not null; default:null"`
}
