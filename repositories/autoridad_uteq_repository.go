package repositories

import (
	"ApiEscuela/models"
	"gorm.io/gorm"
)

type AutoridadUTEQRepository struct {
	db *gorm.DB
}

func NewAutoridadUTEQRepository(db *gorm.DB) *AutoridadUTEQRepository {
	return &AutoridadUTEQRepository{db: db}
}

// CreateAutoridadUTEQ crea una nueva autoridad UTEQ
func (r *AutoridadUTEQRepository) CreateAutoridadUTEQ(autoridad *models.AutoridadUTEQ) error {
	return r.db.Create(autoridad).Error
}

// GetAutoridadUTEQByID obtiene una autoridad UTEQ por ID
func (r *AutoridadUTEQRepository) GetAutoridadUTEQByID(id uint) (*models.AutoridadUTEQ, error) {
	var autoridad models.AutoridadUTEQ
	err := r.db.Preload("Persona").Preload("ProgramasVisita").
		Preload("Dudas").First(&autoridad, id).Error
	if err != nil {
		return nil, err
	}
	return &autoridad, nil
}

// GetAllAutoridadesUTEQ obtiene todas las autoridades UTEQ
func (r *AutoridadUTEQRepository) GetAllAutoridadesUTEQ() ([]models.AutoridadUTEQ, error) {
	var autoridades []models.AutoridadUTEQ
	err := r.db.Preload("Persona").Preload("ProgramasVisita").
		Preload("Dudas").Find(&autoridades).Error
	return autoridades, err
}

// UpdateAutoridadUTEQ actualiza una autoridad UTEQ
func (r *AutoridadUTEQRepository) UpdateAutoridadUTEQ(autoridad *models.AutoridadUTEQ) error {
	return r.db.Save(autoridad).Error
}

// DeleteAutoridadUTEQ elimina una autoridad UTEQ
func (r *AutoridadUTEQRepository) DeleteAutoridadUTEQ(id uint) error {
	return r.db.Delete(&models.AutoridadUTEQ{}, id).Error
}

// GetAutoridadesUTEQByCargo obtiene autoridades por cargo
func (r *AutoridadUTEQRepository) GetAutoridadesUTEQByCargo(cargo string) ([]models.AutoridadUTEQ, error) {
	var autoridades []models.AutoridadUTEQ
	err := r.db.Where("cargo ILIKE ?", "%"+cargo+"%").
		Preload("Persona").Preload("ProgramasVisita").
		Preload("Dudas").Find(&autoridades).Error
	return autoridades, err
}

// GetAutoridadUTEQByPersona obtiene autoridad UTEQ por persona
func (r *AutoridadUTEQRepository) GetAutoridadUTEQByPersona(personaID uint) (*models.AutoridadUTEQ, error) {
	var autoridad models.AutoridadUTEQ
	err := r.db.Where("persona_id = ?", personaID).
		Preload("Persona").Preload("ProgramasVisita").
		Preload("Dudas").First(&autoridad).Error
	if err != nil {
		return nil, err
	}
	return &autoridad, nil
}