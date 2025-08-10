package initial

import (
	"ecommerce_clean/pkgs/logger"
	"ecommerce_clean/pkgs/minio"
)

func InitMinioClient(endpoint string, accessKey string, secretKey string, bucket string, baseURL string, useSSL bool) (*minio.MinioClient, error) {
	minioClient, err := minio.NewMinioClient(
		endpoint,
		accessKey,
		secretKey,
		bucket,
		baseURL,
		useSSL,
	)
	if err != nil {
		logger.Fatalf("Failed to connect to MinIO: %s", err)
	}

	return minioClient, nil
}
