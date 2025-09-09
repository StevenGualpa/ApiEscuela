package models

import (
	"time"
	"gorm.io/gorm"
)

// ProgramaVisita representa un programa de visita programado
type ProgramaVisita struct {
	gorm.Model
	Fecha         time.Time   `json:"fecha" gorm:"not null"`
	InstitucionID uint        `json:"institucion_id" gorm:"not null"`
	
	// Relaciones
	Institucion                      Institucion                        `json:"institucion,omitempty" gorm:"foreignKey:InstitucionID"`
	VisitaDetalles                   []VisitaDetalle                    `json:"visita_detalles,omitempty" gorm:"foreignKey:ProgramaVisitaID"`
	DetalleAutoridadDetallesVisitas  []DetalleAutoridadDetallesVisita   `json:"detalle_autoridad_detalles_visitas,omitempty" gorm:"foreignKey:ProgramaVisitaID"`
}