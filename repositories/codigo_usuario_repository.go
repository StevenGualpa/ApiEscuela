package repositories

import (
	"ApiEscuela/models"
	"time"
	"gorm.io/gorm"
)

type CodigoUsuarioRepository struct {
	db *gorm.DB
}

func NewCodigoUsuarioRepository(db *gorm.DB) *CodigoUsuarioRepository {
	return &CodigoUsuarioRepository{db: db}
}

// Crear inserta un nuevo código para un usuario con expiración de 10 minutos
func (r *CodigoUsuarioRepository) Crear(usuarioID uint, codigo string) error {
	record := &models.CodigoUsuario{
		UsuarioID: usuarioID,
		Codigo:    codigo,
		ExpiraEn:  time.Now().Add(10 * time.Minute),
	}
	return r.db.Create(record).Error
}

// ExisteVigentePorUsuario verifica si el usuario tiene un código no expirado
func (r *CodigoUsuarioRepository) ExisteVigentePorUsuario(usuarioID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.CodigoUsuario{}).
		Where("usuario_id = ? AND expira_en > ?", usuarioID, time.Now()).
		Count(&count).Error
	return count > 0, err
}

// FindLatestByCodigo obtiene el último registro creado para un código dado
func (r *CodigoUsuarioRepository) FindLatestByCodigo(codigo string) (*models.CodigoUsuario, error) {
	var rec models.CodigoUsuario
	err := r.db.Where("codigo = ?", codigo).Order("created_at DESC").First(&rec).Error
	if err != nil {
		return nil, err
	}
	return &rec, nil
}
