package repositories

import (
	"ApiEscuela/models"
	"errors"
	"strings"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type EstudianteUniversitarioRepository struct {
	db *gorm.DB
}

var (
	ErrEstudianteUnivDuplicado = errors.New("estudiante universitario ya existe")
)

func classifyUniqueEstudianteUnivError(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" { // unique_violation
			if strings.Contains(pgErr.Detail, "(persona_id)") {
				return ErrEstudianteUnivDuplicado
			}
			return ErrEstudianteUnivDuplicado
		}
	}
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return ErrEstudianteUnivDuplicado
	}
	return err
}

func NewEstudianteUniversitarioRepository(db *gorm.DB) *EstudianteUniversitarioRepository {
	return &EstudianteUniversitarioRepository{db: db}
}

// CreateEstudianteUniversitario crea un nuevo estudiante universitario
func (r *EstudianteUniversitarioRepository) CreateEstudianteUniversitario(estudiante *models.EstudianteUniversitario) error {
	// Prevalidar por persona
	var count int64
	if err := r.db.Model(&models.EstudianteUniversitario{}).Where("persona_id = ?", estudiante.PersonaID).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return ErrEstudianteUnivDuplicado
	}
	if err := r.db.Create(estudiante).Error; err != nil {
		return classifyUniqueEstudianteUnivError(err)
	}
	return nil
}

// GetEstudianteUniversitarioByID obtiene un estudiante universitario por ID
func (r *EstudianteUniversitarioRepository) GetEstudianteUniversitarioByID(id uint) (*models.EstudianteUniversitario, error) {
	var estudiante models.EstudianteUniversitario
	err := r.db.Preload("Persona").Preload("VisitaDetalleEstudiantesUniversitarios").
		First(&estudiante, id).Error
	if err != nil {
		return nil, err
	}
	return &estudiante, nil
}

// GetAllEstudiantesUniversitarios obtiene todos los estudiantes universitarios
func (r *EstudianteUniversitarioRepository) GetAllEstudiantesUniversitarios() ([]models.EstudianteUniversitario, error) {
	var estudiantes []models.EstudianteUniversitario
	err := r.db.Preload("Persona").Preload("VisitaDetalleEstudiantesUniversitarios").
		Find(&estudiantes).Error
	return estudiantes, err
}

// UpdateEstudianteUniversitario actualiza un estudiante universitario
func (r *EstudianteUniversitarioRepository) UpdateEstudianteUniversitario(estudiante *models.EstudianteUniversitario) error {
	// Verificar duplicado por persona en otro registro
	var count int64
	if err := r.db.Model(&models.EstudianteUniversitario{}).Where("persona_id = ? AND id <> ?", estudiante.PersonaID, estudiante.ID).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return ErrEstudianteUnivDuplicado
	}
	if err := r.db.Save(estudiante).Error; err != nil {
		return classifyUniqueEstudianteUnivError(err)
	}
	return nil
}

// DeleteEstudianteUniversitario elimina un estudiante universitario
func (r *EstudianteUniversitarioRepository) DeleteEstudianteUniversitario(id uint) error {
	return r.db.Delete(&models.EstudianteUniversitario{}, id).Error
}

// GetEstudiantesUniversitariosBySemestre obtiene estudiantes por semestre
func (r *EstudianteUniversitarioRepository) GetEstudiantesUniversitariosBySemestre(semestre int) ([]models.EstudianteUniversitario, error) {
	var estudiantes []models.EstudianteUniversitario
	err := r.db.Where("semestre = ?", semestre).
		Preload("Persona").Preload("VisitaDetalleEstudiantesUniversitarios").
		Find(&estudiantes).Error
	return estudiantes, err
}

// GetEstudianteUniversitarioByPersona obtiene estudiante universitario por persona
func (r *EstudianteUniversitarioRepository) GetEstudianteUniversitarioByPersona(personaID uint) (*models.EstudianteUniversitario, error) {
	var estudiante models.EstudianteUniversitario
	err := r.db.Where("persona_id = ?", personaID).
		Preload("Persona").Preload("VisitaDetalleEstudiantesUniversitarios").
		First(&estudiante).Error
	if err != nil {
		return nil, err
	}
	return &estudiante, nil
}