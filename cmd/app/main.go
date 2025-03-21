package main

import (
	"ecommerce_clean/configs"
	"ecommerce_clean/db"
	"ecommerce_clean/pkgs/logger"
	"ecommerce_clean/pkgs/minio"
	"ecommerce_clean/pkgs/redis"
	"ecommerce_clean/pkgs/token"
	"ecommerce_clean/pkgs/validation"
	"sync"

	orderEntity "ecommerce_clean/internals/order/entity"
	productEntity "ecommerce_clean/internals/product/entity"
	httpServer "ecommerce_clean/internals/server/http"
	userEntity "ecommerce_clean/internals/user/entity"
)

var wg sync.WaitGroup

//	@title			Ecommerce Clean Architecture Swagger API
//	@version		1.0
//	@description	Swagger API for Go Clean Architecture.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Tran Phuoc Anh Quoc
//	@contact.email	anhquoc18092003@gmail.com

//	@license.name	MIT
//	@license.url	https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
func main() {
	cfg := configs.LoadConfig()
	logger.Initialize(cfg.Environment)

	database, err := db.NewDatabase(cfg.DatabaseURI)
	if err != nil {
		logger.Fatal("Cannot connect to database", err)
	}

	if err := database.AutoMigrate(&userEntity.User{}, &productEntity.Product{}, orderEntity.Order{}, orderEntity.OrderLine{}); err != nil {
		logger.Fatal("Database migration fail", err)
	}

	validator := validation.New()

	minioClient, err := minio.NewMinioClient(
		cfg.MinioEndpoint,
		cfg.MinioAccessKey,
		cfg.MinioSecretKey,
		cfg.MinioBucket,
		cfg.MinioBaseurl,
		cfg.MinioUseSSL,
	)
	if err != nil {
		logger.Fatalf("Failed to connect to MinIO: %s", err)
	}

	cache := redis.New(redis.Config{
		Address:  cfg.RedisURI,
		Password: cfg.RedisPassword,
		Database: cfg.RedisDB,
	})

	tokenMaker, err := token.NewJTWMarker()
	if err != nil {
		logger.Fatal(err)
	}

	httpSvr := httpServer.NewServer(validator, database, minioClient, cache, tokenMaker)

	wg.Add(1)

	// Run HTTP server
	go func() {
		defer wg.Done()
		if err := httpSvr.Run(); err != nil {
			logger.Fatal("Running HTTP server error:", err)
		}
	}()

	wg.Wait()
}
