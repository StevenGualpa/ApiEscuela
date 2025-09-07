package models

import "gorm.io/gorm"

// EstudianteUniversitario representa un estudiante universitario
type EstudianteUniversitario struct {
	gorm.Model
	PersonaID uint   `json:"persona_id" gorm:"not null"`
	Semestre  int    `json:"semestre"`
	
	// Relaciones
	Persona       Persona         `json:"persona,omitempty" gorm:"foreignKey:PersonaID"`
	VisitaDetalles []VisitaDetalle `json:"visita_detalles,omitempty" gorm:"foreignKey:EstudianteUniversitarioID"`
}