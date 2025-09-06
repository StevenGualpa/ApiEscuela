package models

import "gorm.io/gorm"

// Student representa un estudiante en el sistema
type Student struct {
	gorm.Model
	Cedula    string `json:"cedula" gorm:"unique;not null"`
	Nombres   string `json:"nombres" gorm:"not null"`
	Apellidos string `json:"apellidos" gorm:"not null"`
	Ciudad    string `json:"ciudad"`
}