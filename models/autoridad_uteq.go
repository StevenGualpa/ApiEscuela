package models

import "gorm.io/gorm"

// AutoridadUTEQ representa una autoridad de la UTEQ
type AutoridadUTEQ struct {
	gorm.Model
	PersonaID uint   `json:"persona_id" gorm:"not null"`
	Cargo     string `json:"cargo"`
	
	// Relaciones
	Persona         Persona          `json:"persona,omitempty" gorm:"foreignKey:PersonaID"`
	ProgramasVisita []ProgramaVisita `json:"programas_visita,omitempty" gorm:"foreignKey:AutoridadUTEQID"`
	Dudas           []Dudas          `json:"dudas,omitempty" gorm:"foreignKey:AutoridadUTEQID"`
}