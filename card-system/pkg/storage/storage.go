// pkg/storage/storage.go
package storage

import (
	"context"
	"io"
)

type Storage interface {
	Upload(ctx context.Context, filename string, reader io.Reader) error
	Delete(ctx context.Context, filename string) error
	GetURL(ctx context.Context, filename string) (string, error)
}

// 本地存储实现
type LocalStorage struct {
	BasePath string
}

func (s *LocalStorage) Upload(ctx context.Context, filename string, reader io.Reader) error {
	path := filepath.Join(s.BasePath, filename)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, reader)
	return err
}

// OSS存储实现（以MinIO为例）
type MinIOStorage struct {
	Client *minio.Client
	Bucket string
}

func (s *MinIOStorage) Upload(ctx context.Context, filename string, reader io.Reader) error {
	_, err := s.Client.PutObject(ctx, s.Bucket, filename, reader, -1, minio.PutObjectOptions{ContentType: "image/jpeg"})
	return err
}