package models

import "gorm.io/gorm"

// Noticia representa una noticia del sistema
type Noticia struct {
	gorm.Model
	Titulo          string `json:"titulo" gorm:"not null"`
	Descripcion     string `json:"descripcion"`
	URLNoticia      string `json:"url_noticia"`
	AutoridadUTEQID uint   `json:"autoridad_uteq_id" gorm:"not null"`

	// Relaciones
	AutoridadUTEQ AutoridadUTEQ `json:"autoridad_uteq,omitempty" gorm:"foreignKey:AutoridadUTEQID"`
}
