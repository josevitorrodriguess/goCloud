package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Provider   string `gorm:"not null"`
	ProviderID string `gorm:"not null"`
	Name       string
	Email      string
	AvatarURL  string
}

// UserRepository define o contrato para persistência de usuários
// na arquitetura limpa.
type UserRepository interface {
	UpsertUser(user *User) error
	UpdateAvatar(userID uint, avatarURL string) error
	DeleteUser(userID uint) error
}
