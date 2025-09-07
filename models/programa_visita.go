package models

import (
	"time"
	"gorm.io/gorm"
)

// ProgramaVisita representa un programa de visita programado
type ProgramaVisita struct {
	gorm.Model
	Fecha            time.Time   `json:"fecha" gorm:"not null"`
	AutoridadUTEQID  uint        `json:"autoridad_uteq_id" gorm:"not null"`
	InstitucionID    uint        `json:"institucion_id" gorm:"not null"`
	
	// Relaciones
	AutoridadUTEQ  AutoridadUTEQ   `json:"autoridad_uteq,omitempty" gorm:"foreignKey:AutoridadUTEQID"`
	Institucion    Institucion     `json:"institucion,omitempty" gorm:"foreignKey:InstitucionID"`
	VisitaDetalles []VisitaDetalle `json:"visita_detalles,omitempty" gorm:"foreignKey:ProgramaVisitaID"`
}