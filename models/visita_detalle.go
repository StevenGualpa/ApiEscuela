package models

import "gorm.io/gorm"

// VisitaDetalle representa el detalle de una visita con actividades
type VisitaDetalle struct {
	gorm.Model
	ActividadID      uint `json:"actividad_id" gorm:"not null"`
	ProgramaVisitaID uint `json:"programa_visita_id" gorm:"not null"`
	Participantes    int  `json:"participantes"`
	
	// Relaciones
	Actividad      Actividad      `json:"actividad,omitempty" gorm:"foreignKey:ActividadID"`
	ProgramaVisita ProgramaVisita `json:"programa_visita,omitempty" gorm:"foreignKey:ProgramaVisitaID"`
}