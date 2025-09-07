package repositories

import (
	"ApiEscuela/models"
	"gorm.io/gorm"
)

type EstudianteUniversitarioRepository struct {
	db *gorm.DB
}

func NewEstudianteUniversitarioRepository(db *gorm.DB) *EstudianteUniversitarioRepository {
	return &EstudianteUniversitarioRepository{db: db}
}

// CreateEstudianteUniversitario crea un nuevo estudiante universitario
func (r *EstudianteUniversitarioRepository) CreateEstudianteUniversitario(estudiante *models.EstudianteUniversitario) error {
	return r.db.Create(estudiante).Error
}

// GetEstudianteUniversitarioByID obtiene un estudiante universitario por ID
func (r *EstudianteUniversitarioRepository) GetEstudianteUniversitarioByID(id uint) (*models.EstudianteUniversitario, error) {
	var estudiante models.EstudianteUniversitario
	err := r.db.Preload("Persona").Preload("VisitaDetalles").
		First(&estudiante, id).Error
	if err != nil {
		return nil, err
	}
	return &estudiante, nil
}

// GetAllEstudiantesUniversitarios obtiene todos los estudiantes universitarios
func (r *EstudianteUniversitarioRepository) GetAllEstudiantesUniversitarios() ([]models.EstudianteUniversitario, error) {
	var estudiantes []models.EstudianteUniversitario
	err := r.db.Preload("Persona").Preload("VisitaDetalles").
		Find(&estudiantes).Error
	return estudiantes, err
}

// UpdateEstudianteUniversitario actualiza un estudiante universitario
func (r *EstudianteUniversitarioRepository) UpdateEstudianteUniversitario(estudiante *models.EstudianteUniversitario) error {
	return r.db.Save(estudiante).Error
}

// DeleteEstudianteUniversitario elimina un estudiante universitario
func (r *EstudianteUniversitarioRepository) DeleteEstudianteUniversitario(id uint) error {
	return r.db.Delete(&models.EstudianteUniversitario{}, id).Error
}

// GetEstudiantesUniversitariosBySemestre obtiene estudiantes por semestre
func (r *EstudianteUniversitarioRepository) GetEstudiantesUniversitariosBySemestre(semestre int) ([]models.EstudianteUniversitario, error) {
	var estudiantes []models.EstudianteUniversitario
	err := r.db.Where("semestre = ?", semestre).
		Preload("Persona").Preload("VisitaDetalles").
		Find(&estudiantes).Error
	return estudiantes, err
}

// GetEstudianteUniversitarioByPersona obtiene estudiante universitario por persona
func (r *EstudianteUniversitarioRepository) GetEstudianteUniversitarioByPersona(personaID uint) (*models.EstudianteUniversitario, error) {
	var estudiante models.EstudianteUniversitario
	err := r.db.Where("persona_id = ?", personaID).
		Preload("Persona").Preload("VisitaDetalles").
		First(&estudiante).Error
	if err != nil {
		return nil, err
	}
	return &estudiante, nil
}