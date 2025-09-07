package models

import "gorm.io/gorm"

// VisitaDetalle representa el detalle de una visita con estudiantes y actividades
type VisitaDetalle struct {
	gorm.Model
	EstudianteUniversitarioID uint `json:"estudiante_universitario_id" gorm:"not null"`
	ActividadID               uint `json:"actividad_id" gorm:"not null"`
	ProgramaVisitaID          uint `json:"programa_visita_id" gorm:"not null"`
	Participantes             int  `json:"participantes"`
	
	// Relaciones
	EstudianteUniversitario EstudianteUniversitario `json:"estudiante_universitario,omitempty" gorm:"foreignKey:EstudianteUniversitarioID"`
	Actividad               Actividad               `json:"actividad,omitempty" gorm:"foreignKey:ActividadID"`
	ProgramaVisita          ProgramaVisita          `json:"programa_visita,omitempty" gorm:"foreignKey:ProgramaVisitaID"`
}