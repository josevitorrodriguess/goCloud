package repository

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/josevitorrodriguess/goCloud/internal/domain"
	"github.com/josevitorrodriguess/goCloud/internal/logger"
)

type FileRepository struct {
	s3     *s3.Client
	bucket string
}

func NewFileRepository(s3 *s3.Client, bucket string) *FileRepository {
	return &FileRepository{
		s3:     s3,
		bucket: bucket,
	}
}

func (fr *FileRepository) UploadFile(file domain.File) error {
	logger.Info("Iniciando upload do arquivo: %s para o usuário: %s", file.Name, file.Email)
	key := file.Email + "/" + file.Name
	input := &s3.PutObjectInput{
		Bucket: &fr.bucket,
		Key:    &key,
		Body:   bytes.NewReader(file.File),
	}
	_, err := fr.s3.PutObject(context.TODO(), input)
	if err != nil {
		logger.Error("Erro ao fazer upload do arquivo: %v", err)
		return err
	}
	logger.Info("Upload realizado com sucesso para o arquivo: %s", file.Name)
	return nil
}

func (fr *FileRepository) GetFile(email, name string) ([]byte, error) {
	logger.Info("Buscando arquivo: %s do usuário: %s", name, email)
	key := email + "/" + name
	input := &s3.GetObjectInput{
		Bucket: &fr.bucket,
		Key:    &key,
	}
	result, err := fr.s3.GetObject(context.TODO(), input)
	if err != nil {
		logger.Error("Erro ao buscar arquivo: %v", err)
		return nil, err
	}
	defer result.Body.Close()
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(result.Body)
	if err != nil {
		logger.Error("Erro ao ler conteúdo do arquivo: %v", err)
		return nil, err
	}
	logger.Info("Arquivo %s baixado com sucesso.", name)
	return buf.Bytes(), nil
}

func (fr *FileRepository) GetAllFiles(email string) ([]string, error) {
	logger.Info("Listando arquivos do usuário: %s", email)
	prefix := email + "/"
	input := &s3.ListObjectsV2Input{
		Bucket: &fr.bucket,
		Prefix: &prefix,
	}
	result, err := fr.s3.ListObjectsV2(context.TODO(), input)
	if err != nil {
		logger.Error("Erro ao listar arquivos: %v", err)
		return nil, err
	}
	var files []string
	for _, obj := range result.Contents {
		if obj.Key != nil {
			files = append(files, (*obj.Key)[len(prefix):])
		}
	}
	logger.Info("Arquivos encontrados: %v", files)
	return files, nil
}

func (fr *FileRepository) DeleteFile(email, name string) error {
	logger.Info("Deletando arquivo: %s do usuário: %s", name, email)
	key := email + "/" + name
	input := &s3.DeleteObjectInput{
		Bucket: &fr.bucket,
		Key:    &key,
	}
	_, err := fr.s3.DeleteObject(context.TODO(), input)
	if err != nil {
		logger.Error("Erro ao deletar arquivo: %v", err)
		return err
	}
	logger.Info("Arquivo %s deletado com sucesso.", name)
	return nil
}
