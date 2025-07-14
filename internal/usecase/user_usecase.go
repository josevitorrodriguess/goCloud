package usecase

import (
	"errors"

	"github.com/josevitorrodriguess/goCloud/internal/domain"
	"github.com/josevitorrodriguess/goCloud/internal/logger"
	"github.com/josevitorrodriguess/goCloud/internal/repository"
)

type UserUsecase struct {
	Repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{Repo: repo}
}

func (uc *UserUsecase) SaveUser(user *domain.User) error {
	if user.Email == "" {
		return errors.New("e-mail obrigatório")
	}
	err := uc.Repo.UpsertUser(user)
	if err != nil {
		logger.Error("Erro ao salvar usuário: %v", err)
		return err
	}
	logger.Info("Usuário salvo com sucesso: %+v", user)
	return nil
}

func (uc *UserUsecase) UpdateAvatar(userID uint, avatarURL string) error {
	if avatarURL == "" {
		return errors.New("URL do avatar obrigatória")
	}
	return uc.Repo.UpdateAvatar(userID, avatarURL)
}

func (uc *UserUsecase) DeleteUser(userID uint) error {
	return uc.Repo.DeleteUser(userID)
}
