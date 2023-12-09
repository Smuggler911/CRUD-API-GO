package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model     //встроеная модель gorm
	Name           string
	SecondName     string
	Age            int
	WorkExperience string
}
