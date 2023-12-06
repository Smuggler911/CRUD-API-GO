package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name           string
	SecondName     string
	Age            int
	WorkExperience string
}
