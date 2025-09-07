package repositories

import (
	"ApiEscuela/models"
	"gorm.io/gorm"
)

type EstudianteRepository struct {
	db *gorm.DB
}

func NewEstudianteRepository(db *gorm.DB) *EstudianteRepository {
	return &EstudianteRepository{db: db}
}

// CreateEstudiante crea un nuevo estudiante
func (r *EstudianteRepository) CreateEstudiante(estudiante *models.Estudiante) error {
	return r.db.Create(estudiante).Error
}

// GetEstudianteByID obtiene un estudiante por ID
func (r *EstudianteRepository) GetEstudianteByID(id uint) (*models.Estudiante, error) {
	var estudiante models.Estudiante
	err := r.db.Preload("Persona").Preload("Institucion").
		Preload("Ciudad").Preload("Ciudad.Provincia").
		Preload("Dudas").First(&estudiante, id).Error
	if err != nil {
		return nil, err
	}
	return &estudiante, nil
}

// GetAllEstudiantes obtiene todos los estudiantes
func (r *EstudianteRepository) GetAllEstudiantes() ([]models.Estudiante, error) {
	var estudiantes []models.Estudiante
	err := r.db.Preload("Persona").Preload("Institucion").
		Preload("Ciudad").Preload("Ciudad.Provincia").
		Find(&estudiantes).Error
	return estudiantes, err
}

// UpdateEstudiante actualiza un estudiante
func (r *EstudianteRepository) UpdateEstudiante(estudiante *models.Estudiante) error {
	return r.db.Save(estudiante).Error
}

// DeleteEstudiante elimina un estudiante
func (r *EstudianteRepository) DeleteEstudiante(id uint) error {
	return r.db.Delete(&models.Estudiante{}, id).Error
}

// GetEstudiantesByCity obtiene estudiantes por ciudad
func (r *EstudianteRepository) GetEstudiantesByCity(ciudadID uint) ([]models.Estudiante, error) {
	var estudiantes []models.Estudiante
	err := r.db.Where("ciudad_id = ?", ciudadID).
		Preload("Persona").Preload("Institucion").
		Preload("Ciudad").Preload("Ciudad.Provincia").
		Find(&estudiantes).Error
	return estudiantes, err
}

// GetEstudiantesByInstitucion obtiene estudiantes por institución
func (r *EstudianteRepository) GetEstudiantesByInstitucion(institucionID uint) ([]models.Estudiante, error) {
	var estudiantes []models.Estudiante
	err := r.db.Where("institucion_id = ?", institucionID).
		Preload("Persona").Preload("Institucion").
		Preload("Ciudad").Preload("Ciudad.Provincia").
		Find(&estudiantes).Error
	return estudiantes, err
}

// GetEstudiantesByEspecialidad obtiene estudiantes por especialidad
func (r *EstudianteRepository) GetEstudiantesByEspecialidad(especialidad string) ([]models.Estudiante, error) {
	var estudiantes []models.Estudiante
	err := r.db.Where("especialidad ILIKE ?", "%"+especialidad+"%").
		Preload("Persona").Preload("Institucion").
		Preload("Ciudad").Preload("Ciudad.Provincia").
		Find(&estudiantes).Error
	return estudiantes, err
}