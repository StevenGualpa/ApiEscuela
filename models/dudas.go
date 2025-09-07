package models

import "gorm.io/gorm"

// Dudas representa las dudas que pueden tener los estudiantes
type Dudas struct {
	gorm.Model
	Pregunta        string `json:"pregunta" gorm:"not null"`
	Respuesta       string `json:"respuesta"`
	EstudianteID    uint   `json:"estudiante_id" gorm:"not null"`
	AutoridadUTEQID *uint  `json:"autoridad_uteq_id"` // Opcional, puede no estar asignada
	
	// Relaciones
	Estudiante    Estudiante   `json:"estudiante,omitempty" gorm:"foreignKey:EstudianteID"`
	AutoridadUTEQ *AutoridadUTEQ `json:"autoridad_uteq,omitempty" gorm:"foreignKey:AutoridadUTEQID"`
}