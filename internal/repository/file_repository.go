package repository

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/josevitorrodriguess/goCloud/internal/domain"
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
	key := file.Email + "/" + file.Name
	input := &s3.PutObjectInput{
		Bucket: &fr.bucket,
		Key:    &key,
		Body:   bytes.NewReader(file.File),
	}
	_, err := fr.s3.PutObject(context.TODO(), input)
	return err
}

func (fr *FileRepository) GetFile(email, name string) ([]byte, error) {
	key := email + "/" + name
	input := &s3.GetObjectInput{
		Bucket: &fr.bucket,
		Key:    &key,
	}
	result, err := fr.s3.GetObject(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(result.Body)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (fr *FileRepository) GetAllFiles(email string) ([]string, error) {
	prefix := email + "/"
	input := &s3.ListObjectsV2Input{
		Bucket: &fr.bucket,
		Prefix: &prefix,
	}
	result, err := fr.s3.ListObjectsV2(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	var files []string
	for _, obj := range result.Contents {

		if obj.Key != nil {
			files = append(files, (*obj.Key)[len(prefix):])
		}
	}
	return files, nil
}

func (fr *FileRepository) DeleteFile(email, name string) error {
	key := email + "/" + name
	input := &s3.DeleteObjectInput{
		Bucket: &fr.bucket,
		Key:    &key,
	}
	_, err := fr.s3.DeleteObject(context.TODO(), input)
	return err
}
