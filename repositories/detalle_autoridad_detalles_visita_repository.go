package repositories

import (
	"ApiEscuela/models"
	"gorm.io/gorm"
)

type DetalleAutoridadDetallesVisitaRepository struct {
	db *gorm.DB
}

func NewDetalleAutoridadDetallesVisitaRepository(db *gorm.DB) *DetalleAutoridadDetallesVisitaRepository {
	return &DetalleAutoridadDetallesVisitaRepository{db: db}
}

// CreateDetalleAutoridadDetallesVisita crea un nuevo detalle de autoridad para visita
func (r *DetalleAutoridadDetallesVisitaRepository) CreateDetalleAutoridadDetallesVisita(detalle *models.DetalleAutoridadDetallesVisita) error {
	return r.db.Create(detalle).Error
}

// GetDetalleAutoridadDetallesVisitaByID obtiene un detalle por ID
func (r *DetalleAutoridadDetallesVisitaRepository) GetDetalleAutoridadDetallesVisitaByID(id uint) (*models.DetalleAutoridadDetallesVisita, error) {
	var detalle models.DetalleAutoridadDetallesVisita
	err := r.db.Preload("ProgramaVisita").Preload("AutoridadUTEQ").First(&detalle, id).Error
	if err != nil {
		return nil, err
	}
	return &detalle, nil
}

// GetAllDetalleAutoridadDetallesVisitas obtiene todos los detalles
func (r *DetalleAutoridadDetallesVisitaRepository) GetAllDetalleAutoridadDetallesVisitas() ([]models.DetalleAutoridadDetallesVisita, error) {
	var detalles []models.DetalleAutoridadDetallesVisita
	err := r.db.Preload("ProgramaVisita").Preload("AutoridadUTEQ").Find(&detalles).Error
	return detalles, err
}

// UpdateDetalleAutoridadDetallesVisita actualiza un detalle
func (r *DetalleAutoridadDetallesVisitaRepository) UpdateDetalleAutoridadDetallesVisita(detalle *models.DetalleAutoridadDetallesVisita) error {
	return r.db.Save(detalle).Error
}

// DeleteDetalleAutoridadDetallesVisita elimina un detalle
func (r *DetalleAutoridadDetallesVisitaRepository) DeleteDetalleAutoridadDetallesVisita(id uint) error {
	return r.db.Delete(&models.DetalleAutoridadDetallesVisita{}, id).Error
}

// GetDetallesByProgramaVisitaID obtiene todos los detalles de un programa de visita específico
func (r *DetalleAutoridadDetallesVisitaRepository) GetDetallesByProgramaVisitaID(programaVisitaID uint) ([]models.DetalleAutoridadDetallesVisita, error) {
	var detalles []models.DetalleAutoridadDetallesVisita
	err := r.db.Where("programa_visita_id = ?", programaVisitaID).
		Preload("ProgramaVisita").Preload("AutoridadUTEQ").Find(&detalles).Error
	return detalles, err
}

// GetDetallesByAutoridadID obtiene todos los detalles de una autoridad específica
func (r *DetalleAutoridadDetallesVisitaRepository) GetDetallesByAutoridadID(autoridadID uint) ([]models.DetalleAutoridadDetallesVisita, error) {
	var detalles []models.DetalleAutoridadDetallesVisita
	err := r.db.Where("autoridad_uteqid = ?", autoridadID).
		Preload("ProgramaVisita").Preload("AutoridadUTEQ").Find(&detalles).Error
	return detalles, err
}

// DeleteDetallesByProgramaVisitaID elimina todos los detalles de un programa de visita
func (r *DetalleAutoridadDetallesVisitaRepository) DeleteDetallesByProgramaVisitaID(programaVisitaID uint) error {
	return r.db.Where("programa_visita_id = ?", programaVisitaID).Delete(&models.DetalleAutoridadDetallesVisita{}).Error
}