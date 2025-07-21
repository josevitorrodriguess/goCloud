package usecase

import (
	"github.com/josevitorrodriguess/goCloud/internal/domain"
	"github.com/josevitorrodriguess/goCloud/internal/encryption"
	"github.com/josevitorrodriguess/goCloud/internal/logger"
	"github.com/josevitorrodriguess/goCloud/internal/repository"
)

type FileUsecase struct {
	Repo repository.FileRepository
}

func NewFileUsecase(repo repository.FileRepository) *FileUsecase {
	return &FileUsecase{
		Repo: repo,
	}
}

func (fu *FileUsecase) UploadFileUC(file []byte, filename, email string) error {
	encryptedFile, err := encryption.EncryptFile(file, email)
	if err != nil {
		logger.Error("failed to encrypt file: %v", err)
		return err
	}

	data := domain.File{
		File:  encryptedFile,
		Name:  filename,
		Email: email,
	}

	err = fu.Repo.UploadFile(data)
	if err != nil {
		logger.Error("error to upload to bucket: %v", err)
		return err
	}

	logger.Info("Sucessfull to upload file")
	return nil
}

func (fu *FileUsecase) GetFileUC(email, filename string) ([]byte, error) {
	logger.Info("Buscando arquivo %s do usuário %s", filename, email)
	fileData, err := fu.Repo.GetFile(email, filename)
	if err != nil {
		logger.Error("erro ao buscar arquivo: %v", err)
		return nil, err
	}

	decryptedFile, err := encryption.DecryptFile(fileData, email)
	if err != nil {
		logger.Error("erro ao descriptografar arquivo: %v", err)
		return nil, err
	}

	logger.Info("Arquivo %s recuperado e descriptografado com sucesso", filename)
	return decryptedFile, nil
}

func (fu *FileUsecase) GetAllFilesUC(email string) ([]string, error) {
	logger.Info("Listando arquivos do usuário %s", email)
	files, err := fu.Repo.GetAllFiles(email)
	if err != nil {
		logger.Error("erro ao listar arquivos: %v", err)
		return nil, err
	}
	logger.Info("Arquivos listados com sucesso para o usuário %s", email)
	return files, nil
}

func (fu *FileUsecase) DeleteFileUC(email, filename string) error {
	logger.Info("Deletando arquivo %s do usuário %s", filename, email)
	err := fu.Repo.DeleteFile(email, filename)
	if err != nil {
		logger.Error("erro ao deletar arquivo: %v", err)
		return err
	}
	logger.Info("Arquivo %s deletado com sucesso para o usuário %s", filename, email)
	return nil
}
