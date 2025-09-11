package repositories

import (
	"ApiEscuela/models"
	"gorm.io/gorm"
)

type VisitaDetalleRepository struct {
	db *gorm.DB
}

func NewVisitaDetalleRepository(db *gorm.DB) *VisitaDetalleRepository {
	return &VisitaDetalleRepository{db: db}
}

// CreateVisitaDetalle crea un nuevo detalle de visita
func (r *VisitaDetalleRepository) CreateVisitaDetalle(detalle *models.VisitaDetalle) error {
	return r.db.Create(detalle).Error
}

// GetVisitaDetalleByID obtiene un detalle de visita por ID
func (r *VisitaDetalleRepository) GetVisitaDetalleByID(id uint) (*models.VisitaDetalle, error) {
	var detalle models.VisitaDetalle
	err := r.db.Preload("Actividad").Preload("Actividad.Tematica").
		Preload("ProgramaVisita").Preload("ProgramaVisita.Institucion").
		Preload("ProgramaVisita.AutoridadUTEQ").
		First(&detalle, id).Error
	if err != nil {
		return nil, err
	}
	return &detalle, nil
}

// GetAllVisitaDetalles obtiene todos los detalles de visita
func (r *VisitaDetalleRepository) GetAllVisitaDetalles() ([]models.VisitaDetalle, error) {
	var detalles []models.VisitaDetalle
	err := r.db.Preload("Actividad").Preload("Actividad.Tematica").
		Preload("ProgramaVisita").Preload("ProgramaVisita.Institucion").
		Preload("ProgramaVisita.AutoridadUTEQ").
		Find(&detalles).Error
	return detalles, err
}

// UpdateVisitaDetalle actualiza un detalle de visita
func (r *VisitaDetalleRepository) UpdateVisitaDetalle(detalle *models.VisitaDetalle) error {
	return r.db.Save(detalle).Error
}

// DeleteVisitaDetalle elimina un detalle de visita
func (r *VisitaDetalleRepository) DeleteVisitaDetalle(id uint) error {
	return r.db.Delete(&models.VisitaDetalle{}, id).Error
}

// GetVisitaDetallesByActividad obtiene detalles por actividad
func (r *VisitaDetalleRepository) GetVisitaDetallesByActividad(actividadID uint) ([]models.VisitaDetalle, error) {
	var detalles []models.VisitaDetalle
	err := r.db.Where("actividad_id = ?", actividadID).
		Preload("Actividad").Preload("Actividad.Tematica").
		Preload("ProgramaVisita").Preload("ProgramaVisita.Institucion").
		Preload("ProgramaVisita.AutoridadUTEQ").
		Find(&detalles).Error
	return detalles, err
}

// GetVisitaDetallesByPrograma obtiene detalles por programa de visita
func (r *VisitaDetalleRepository) GetVisitaDetallesByPrograma(programaID uint) ([]models.VisitaDetalle, error) {
	var detalles []models.VisitaDetalle
	err := r.db.Where("programa_visita_id = ?", programaID).
		Preload("Actividad").Preload("Actividad.Tematica").
		Preload("ProgramaVisita").Preload("ProgramaVisita.Institucion").
		Preload("ProgramaVisita.AutoridadUTEQ").
		Find(&detalles).Error
	return detalles, err
}

// GetVisitaDetallesByParticipantes obtiene detalles por rango de participantes
func (r *VisitaDetalleRepository) GetVisitaDetallesByParticipantes(minParticipantes, maxParticipantes int) ([]models.VisitaDetalle, error) {
	var detalles []models.VisitaDetalle
	err := r.db.Where("participantes BETWEEN ? AND ?", minParticipantes, maxParticipantes).
		Preload("Actividad").Preload("Actividad.Tematica").
		Preload("ProgramaVisita").Preload("ProgramaVisita.Institucion").
		Preload("ProgramaVisita.AutoridadUTEQ").
		Find(&detalles).Error
	return detalles, err
}

// GetEstadisticasParticipacion obtiene estadísticas de participación
func (r *VisitaDetalleRepository) GetEstadisticasParticipacion() (map[string]interface{}, error) {
	var totalDetalles int64
	var totalParticipantes int64
	var promedioParticipantes float64

	// Contar total de detalles
	if err := r.db.Model(&models.VisitaDetalle{}).Count(&totalDetalles).Error; err != nil {
		return nil, err
	}

	// Sumar total de participantes
	if err := r.db.Model(&models.VisitaDetalle{}).Select("COALESCE(SUM(participantes), 0)").Scan(&totalParticipantes).Error; err != nil {
		return nil, err
	}

	// Calcular promedio
	if totalDetalles > 0 {
		promedioParticipantes = float64(totalParticipantes) / float64(totalDetalles)
	}

	return map[string]interface{}{
		"total_visitas":           totalDetalles,
		"total_participantes":     totalParticipantes,
		"promedio_participantes":  promedioParticipantes,
	}, nil
}