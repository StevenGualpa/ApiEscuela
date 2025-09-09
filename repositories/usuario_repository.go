package repositories

import (
	"ApiEscuela/models"
	"gorm.io/gorm"
)

type UsuarioRepository struct {
	db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{db: db}
}

// CreateUsuario crea un nuevo usuario
func (r *UsuarioRepository) CreateUsuario(usuario *models.Usuario) error {
	return r.db.Create(usuario).Error
}

// GetUsuarioByID obtiene un usuario por ID
func (r *UsuarioRepository) GetUsuarioByID(id uint) (*models.Usuario, error) {
	var usuario models.Usuario
	err := r.db.Preload("Persona").Preload("TipoUsuario").
		First(&usuario, id).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

// GetAllUsuarios obtiene todos los usuarios
func (r *UsuarioRepository) GetAllUsuarios() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	err := r.db.Preload("Persona").Preload("TipoUsuario").
		Find(&usuarios).Error
	return usuarios, err
}

// UpdateUsuario actualiza un usuario
func (r *UsuarioRepository) UpdateUsuario(usuario *models.Usuario) error {
	return r.db.Save(usuario).Error
}

// DeleteUsuario elimina un usuario
func (r *UsuarioRepository) DeleteUsuario(id uint) error {
	return r.db.Delete(&models.Usuario{}, id).Error
}

// GetUsuarioByUsername busca usuario por nombre de usuario
func (r *UsuarioRepository) GetUsuarioByUsername(username string) (*models.Usuario, error) {
	var usuario models.Usuario
	err := r.db.Where("usuario = ?", username).
		Preload("Persona").Preload("TipoUsuario").
		First(&usuario).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

// GetUsuariosByTipo obtiene usuarios por tipo
func (r *UsuarioRepository) GetUsuariosByTipo(tipoUsuarioID uint) ([]models.Usuario, error) {
	var usuarios []models.Usuario
	err := r.db.Where("tipo_usuario_id = ?", tipoUsuarioID).
		Preload("Persona").Preload("TipoUsuario").
		Find(&usuarios).Error
	return usuarios, err
}

// GetUsuariosByPersona obtiene usuarios por persona
func (r *UsuarioRepository) GetUsuariosByPersona(personaID uint) ([]models.Usuario, error) {
	var usuarios []models.Usuario
	err := r.db.Where("persona_id = ?", personaID).
		Preload("Persona").Preload("TipoUsuario").
		Find(&usuarios).Error
	return usuarios, err
}

// ValidateLogin valida credenciales de usuario
func (r *UsuarioRepository) ValidateLogin(username, password string) (*models.Usuario, error) {
	var usuario models.Usuario
	err := r.db.Where("usuario = ? AND contraseña = ?", username, password).
		Preload("Persona").Preload("TipoUsuario").
		First(&usuario).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

// GetAllUsuariosIncludingDeleted obtiene todos los usuarios incluyendo los eliminados
func (r *UsuarioRepository) GetAllUsuariosIncludingDeleted() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	err := r.db.Unscoped().Preload("Persona").Preload("TipoUsuario").
		Find(&usuarios).Error
	return usuarios, err
}

// GetDeletedUsuarios obtiene solo los usuarios eliminados
func (r *UsuarioRepository) GetDeletedUsuarios() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	err := r.db.Unscoped().Where("deleted_at IS NOT NULL").
		Preload("Persona").Preload("TipoUsuario").
		Find(&usuarios).Error
	return usuarios, err
}

// RestoreUsuario restaura un usuario eliminado (soft delete)
func (r *UsuarioRepository) RestoreUsuario(id uint) error {
	return r.db.Unscoped().Model(&models.Usuario{}).
		Where("id = ?", id).Update("deleted_at", nil).Error
}

// GetUsuarioByIDIncludingDeleted obtiene un usuario por ID incluyendo eliminados
func (r *UsuarioRepository) GetUsuarioByIDIncludingDeleted(id uint) (*models.Usuario, error) {
	var usuario models.Usuario
	err := r.db.Unscoped().Preload("Persona").Preload("TipoUsuario").
		First(&usuario, id).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}