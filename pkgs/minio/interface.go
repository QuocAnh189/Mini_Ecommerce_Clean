package minio

import (
	"context"
	"mime/multipart"
)

type IUploadService interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader, folder string) (string, error)
	DeleteFile(ctx context.Context, fileURL string) error
}
