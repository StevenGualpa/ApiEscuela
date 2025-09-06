package repositories

import (
	"ApiEscuela/models"
	"gorm.io/gorm"
)

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

// CreateStudent crea un nuevo estudiante
func (r *StudentRepository) CreateStudent(student *models.Student) error {
	return r.db.Create(student).Error
}

// GetStudentByID obtiene un estudiante por ID
func (r *StudentRepository) GetStudentByID(id uint) (*models.Student, error) {
	var student models.Student
	err := r.db.First(&student, id).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

// GetStudentByCedula obtiene un estudiante por cédula
func (r *StudentRepository) GetStudentByCedula(cedula string) (*models.Student, error) {
	var student models.Student
	err := r.db.Where("cedula = ?", cedula).First(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

// GetAllStudents obtiene todos los estudiantes
func (r *StudentRepository) GetAllStudents() ([]models.Student, error) {
	var students []models.Student
	err := r.db.Find(&students).Error
	return students, err
}

// UpdateStudent actualiza un estudiante
func (r *StudentRepository) UpdateStudent(student *models.Student) error {
	return r.db.Save(student).Error
}

// DeleteStudent elimina un estudiante
func (r *StudentRepository) DeleteStudent(id uint) error {
	return r.db.Delete(&models.Student{}, id).Error
}

// GetStudentsByCity obtiene estudiantes por ciudad
func (r *StudentRepository) GetStudentsByCity(ciudad string) ([]models.Student, error) {
	var students []models.Student
	err := r.db.Where("ciudad = ?", ciudad).Find(&students).Error
	return students, err
}