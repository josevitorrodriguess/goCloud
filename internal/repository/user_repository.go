package repository

import (
	"github.com/josevitorrodriguess/goCloud/internal/domain"
	"github.com/josevitorrodriguess/goCloud/internal/logger"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) UpsertUser(user *domain.User) error {
	var existing domain.User
	err := r.DB.Where("provider = ? AND provider_id = ?", user.Provider, user.ProviderID).First(&existing).Error
	if err == gorm.ErrRecordNotFound {
		return r.DB.Create(user).Error
	} else if err != nil {
		return err
	}
	
	logger.Info("Atualizando usuário")
	existing.Name = user.Name
	existing.Email = user.Email
	existing.AvatarURL = user.AvatarURL
	logger.Info("Usuário atualizado: %+v", existing)
	return r.DB.Save(&existing).Error
}

func (r *UserRepository) UpdateAvatar(userID uint, avatarURL string) error {
	return r.DB.Model(&domain.User{}).Where("id = ?", userID).Update("avatar_url", avatarURL).Error
}

func (r *UserRepository) DeleteUser(userID uint) error {
	return r.DB.Delete(&domain.User{}, userID).Error
}
