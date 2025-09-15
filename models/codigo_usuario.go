package models

import (
	"time"
	"gorm.io/gorm"
)

// CodigoUsuario representa un código temporal asociado a un usuario
// Tabla exacta: codigosusuarios
// ExpiraEn: fecha/hora de expiración (10 minutos desde su creación)
type CodigoUsuario struct {
	gorm.Model
	UsuarioID uint      `json:"usuario_id" gorm:"not null;index"`
	Codigo    string    `json:"codigo" gorm:"not null;size:10;index"`
	ExpiraEn  time.Time `json:"expira_en" gorm:"not null;index"`
}

// TableName fuerza el nombre de la tabla a "codigosusuarios"
func (CodigoUsuario) TableName() string { return "codigosusuarios" }
